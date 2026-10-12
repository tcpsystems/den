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
	"github.com/tcpsystems/den/internal/initd"
	"github.com/tcpsystems/den/internal/orch"
)

func main() {
	home, _ := os.UserHomeDir()
	sock := flag.String("sock", filepath.Join(home, ".local", "share", "harness", "den.sock"), "daemon unix socket")
	repo := flag.String("repo", ".", "repository whose .den board this screen edits")
	flag.Parse()
	if _, err := tea.NewProgram(newModel(*sock, *repo)).Run(); err != nil {
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
	repo   string
	sub    chan tea.Msg
	state  console.State
	board  initd.Board
	err    string
	status string
	notice string
	seats  []orch.Seat
}

func newModel(sock, repo string) model {
	board, err := initd.Load(repo, sock)
	m := model{sock: sock, repo: repo, sub: make(chan tea.Msg), status: "connecting", board: board}
	if err != nil {
		m.err = err.Error()
	}
	return m
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
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "w":
			m.notice = m.enable("web")
		case "p":
			m.notice = m.enable("pages")
		case "f":
			m.notice = m.enable("files")
		case "s":
			m.notice = m.openSeats()
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

func (m *model) openSeats() string {
	if !m.enableOK("seats") {
		return "seats stay off"
	}
	repo, err := filepath.Abs(m.repo)
	if err != nil {
		return err.Error()
	}
	w := orch.Wave{
		ID: "tui", Branch: "wave/tui-seats", Base: "HEAD", Repo: repo,
		Worktrees: filepath.Join(filepath.Dir(repo), "harness-plan-wt"),
		Seats: []orch.Seat{
			{ID: "s1", Role: "researcher", Owns: []string{"research/drafts/tui-s1"}},
			{ID: "s2", Role: "architect", Owns: []string{"plan/analysis/tui-s2"}},
			{ID: "s3", Role: "builder", Owns: []string{"internal/tui-s3"}},
			{ID: "s4", Role: "reviewer", Owns: []string{"plan/analysis/tui-s4"}},
		},
	}
	opened, err := orch.Open(w)
	if err != nil {
		return err.Error()
	}
	m.seats = opened.Seats
	if err := initd.Save(m.board); err != nil {
		return err.Error()
	}
	return fmt.Sprintf("seats on  %s  %d worktrees", opened.Branch, len(opened.Seats))
}

func (m *model) enableOK(id string) bool {
	return initd.Enable(&m.board, id)
}

func (m *model) enable(id string) string {
	if !initd.Enable(&m.board, id) {
		return id + " stays off. this build has no process for it"
	}
	if err := initd.Save(m.board); err != nil {
		return err.Error()
	}
	return id + " on"
}

func (m model) View() string {
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("208")).Render("den")
	head := title + " task manager  " + m.status
	if m.err != "" {
		head += "  " + m.err
	}
	var b strings.Builder
	b.WriteString(head + "\n")
	b.WriteString("w web   p pages   f files   s seats   q quit\n\n")
	b.WriteString("parts\n")
	for _, part := range m.board.Parts {
		b.WriteString(fmt.Sprintf("  %-12s %-5s %s\n", part.Title, part.Status, part.Note))
	}
	if m.notice != "" {
		b.WriteString("\n" + m.notice + "\n")
	}
	if len(m.seats) > 0 {
		b.WriteString("\nseats\n")
		for _, s := range m.seats {
			b.WriteString(fmt.Sprintf("  %-4s %-12s %-8s %s\n", s.ID, s.Role, s.Status, s.Head))
		}
	}
	b.WriteString("\nservices\n")
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
