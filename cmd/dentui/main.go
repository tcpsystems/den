package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tcpsystems/den/internal/console"
)

func main() {
	home, _ := os.UserHomeDir()
	sock := flag.String("sock", filepath.Join(home, ".local", "share", "harness", "den.sock"), "daemon unix socket")
	flag.Parse()
	if _, err := tea.NewProgram(newModel(*sock)).Run(); err != nil {
		fmt.Println(err)
	}
}

func client(sock string) *http.Client {
	return &http.Client{Transport: &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "unix", sock)
		},
	}}
}

type model struct {
	sock   string
	sub    chan tea.Msg
	state  console.State
	err    string
	status string
}

func newModel(sock string) model {
	return model{sock: sock, sub: make(chan tea.Msg), status: "connecting"}
}

type stateMsg struct{ state console.State }
type errMsg struct{ err string }

func (m model) Init() tea.Cmd {
	go readEvents(m.sock, m.sub)
	return wait(m.sub)
}

func wait(sub chan tea.Msg) tea.Cmd {
	return func() tea.Msg { return <-sub }
}

func readEvents(sock string, sub chan tea.Msg) {
	c := client(sock)
	for {
		req, err := http.NewRequest(http.MethodGet, "http://den/api/events", nil)
		if err != nil {
			sub <- errMsg{err.Error()}
			time.Sleep(3 * time.Second)
			continue
		}
		resp, err := c.Do(req)
		if err != nil {
			sub <- errMsg{err.Error()}
			time.Sleep(3 * time.Second)
			continue
		}
		sc := bufio.NewScanner(resp.Body)
		sc.Buffer(make([]byte, 0, 64*1024), 2*1024*1024)
		var event string
		var data strings.Builder
		for sc.Scan() {
			line := sc.Text()
			if line == "" {
				if event == "state" && data.Len() > 0 {
					var st console.State
					if json.Unmarshal([]byte(data.String()), &st) == nil {
						sub <- stateMsg{st}
					}
				}
				event = ""
				data.Reset()
				continue
			}
			if strings.HasPrefix(line, ":") {
				continue
			}
			if strings.HasPrefix(line, "event:") {
				event = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
			}
			if strings.HasPrefix(line, "data:") {
				data.WriteString(strings.TrimSpace(strings.TrimPrefix(line, "data:")))
			}
		}
		resp.Body.Close()
		if err := sc.Err(); err != nil && err != io.EOF {
			sub <- errMsg{err.Error()}
		} else {
			sub <- errMsg{"event stream closed"}
		}
		time.Sleep(3 * time.Second)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case stateMsg:
		m.state = msg.state
		m.status = "live"
		m.err = ""
		return m, wait(m.sub)
	case errMsg:
		m.err = msg.err
		m.status = "reconnecting"
		return m, wait(m.sub)
	}
	return m, nil
}

func (m model) View() string {
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("208")).Render("den")
	head := title + " task manager  " + m.status
	if m.err != "" {
		head += "  " + m.err
	}
	var b strings.Builder
	b.WriteString(head + "\n\n")
	b.WriteString("services\n")
	for name, st := range m.state.Services {
		b.WriteString("  " + name + "  " + st + "\n")
	}
	b.WriteString("\nnest\n")
	for _, p := range m.state.Projects {
		b.WriteString("  " + p.Org + "/" + p.Name + "  " + p.Status + "  " + p.Role + "\n")
	}
	b.WriteString("\nbranches ahead of develop\n")
	shown := 0
	for _, br := range m.state.Branches {
		if br.Ahead == 0 && br.Kind == "draft" {
			continue
		}
		if br.Kind == "archive" {
			continue
		}
		b.WriteString(fmt.Sprintf("  %-22s +%d  %s\n", br.Name, br.Ahead, br.Subject))
		shown++
		if shown == 20 {
			break
		}
	}
	if shown == 0 {
		b.WriteString("  none\n")
	}
	b.WriteString("\nq quit\n")
	return b.String()
}
