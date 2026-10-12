package console

import (
	"context"
	"os"
	"os/exec"
	"sync"
	"path/filepath"
	"strings"
	"time"
)

// Slot is one later expansion. This build implements only the controller itself.
type Slot struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Note   string `json:"note"`
}

// State is the only payload the page renders.
type State struct {
	Repo     string            `json:"repo"`
	Services map[string]string `json:"services"`
	Counts   map[string]int    `json:"counts"`
	Branches []Branch          `json:"branches"`
	Slots    []Slot            `json:"slots"`
	Projects []Project         `json:"projects"`
}

// Project is one repo sitting in the nest. The harness does not own its GitHub history.
type Project struct {
	Org    string `json:"org"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

// Branch is one git ref the controller is allowed to show.
type Branch struct {
	Name    string `json:"name"`
	SHA     string `json:"sha"`
	Subject string `json:"subject"`
	Ahead   int    `json:"ahead"`
	Kind    string `json:"kind"`
}

// Slots is the expansion list. One table, used by the page and by tests.
func Slots() []Slot {
	return []Slot{
		{ID: "go", Status: "this", Note: "API binary"},
		{ID: "tui", Status: "this", Note: "cmd/dentui, Bubble Tea, same event stream"},
		{ID: "svelte", Status: "this", Note: "live page over SSE, no Node at runtime"},
		{ID: "sse", Status: "this", Note: "event: state, id, retry 3000, comment keepalive. MDN EventSource"},
		{ID: "rust", Status: "later", Note: "not this binary"},
		{ID: "typescript", Status: "later", Note: "fox panel language, not the controller"},
		{ID: "javascript", Status: "later", Note: "only where a page already uses it"},
		{ID: "react", Status: "later", Note: "not the controller"},
		{ID: "shadcn", Status: "later", Note: "not the controller"},
		{ID: "webcomponents", Status: "later", Note: "not required for this page"},
		{ID: "websocket", Status: "later", Note: "bidirectional control, not this read stream"},
		{ID: "quic", Status: "later", Note: "not the loopback listen"},
		{ID: "anthropic", Status: "later", Note: "a named model, not the page"},
		{ID: "reasonix", Status: "later", Note: "a named model, not the page"},
		{ID: "cloudflare", Status: "later", Note: "fox edge, not this listen address"},
		{ID: "workers", Status: "later", Note: "fox, not denctl"},
		{ID: "hetzner", Status: "later", Note: "lab and staging only"},
		{ID: "scaleblade", Status: "later", Note: "not customer metal"},
		{ID: "hosting", Status: "later", Note: "capacity product stays in fox"},
		{ID: "linux", Status: "later", Note: "Debian bootstrap is separate"},
		{ID: "design", Status: "later", Note: "one accent, readable type"},
		{ID: "niche", Status: "later", Note: "plan/niches"},
		{ID: "usage", Status: "later", Note: "the page flow in plan/controller.md"},
		{ID: "optimise", Status: "later", Note: "cut the page, do not add a kit"},
		{ID: "maintain", Status: "later", Note: "one probe file per service"},
		{ID: "sentry", Status: "later", Note: "no DSN in this payload, not shipped in fox"},
		{ID: "newrelic", Status: "later", Note: "not wired"},
		{ID: "opentelemetry", Status: "later", Note: "no collector on this listen address"},
		{ID: "telemetry", Status: "later", Note: "counts on this page are the only telemetry"},
		{ID: "webperf", Status: "later", Note: "not a lab run inside the controller"},
		{ID: "goperf", Status: "later", Note: "pprof is not served"},
		{ID: "cgo", Status: "later", Note: "this binary does not use cgo"},
		{ID: "sveltekit", Status: "later", Note: "fox panel, not this page"},
		{ID: "sql", Status: "later", Note: "den state is git, not a database"},
		{ID: "d1", Status: "later", Note: "fox database, not denctl"},
		{ID: "data", Status: "later", Note: "radar and OpenViking stay the stores"},
		{ID: "privacy", Status: "this", Note: "no secret in /api/state or the event stream"},
		{ID: "dataflow", Status: "later", Note: "probe to page, one direction"},
		{ID: "userflow", Status: "later", Note: "open, read, reload"},
		{ID: "agentflow", Status: "later", Note: "leader, orchestrator, seat"},
		{ID: "devevalflow", Status: "later", Note: "note, test, reeval"},
		{ID: "dry", Status: "this", Note: "slots and branches come from this package"},
	}
}

// Collect reads git and probes. Failures become down or an empty list.
func Collect(ctx context.Context, repo string) State {
	rows := readBranches(ctx, repo)
	ahead := 0
	sections := 0
	drafts := 0
	for _, row := range rows {
		if row.Ahead > 0 {
			ahead++
		}
		switch row.Kind {
		case "section":
			sections++
		case "draft":
			drafts++
		}
	}
	return State{
		Repo: "tcpsystems/den",
		Services: map[string]string{
			"mesh":      probe(ctx, "http://127.0.0.1:8787/api/health"),
			"embed":     probe(ctx, "http://127.0.0.1:8091/v1/models"),
			"memory":    probe(ctx, "http://127.0.0.1:1933/health"),
			"omniroute": probe(ctx, "https://ai.tcp.systems/api/health"),
		},
		Counts: map[string]int{
			"branches": len(rows),
			"sections": sections,
			"drafts":   drafts,
			"ahead":    ahead,
		},
		Branches: rows,
		Slots:    Slots(),
		Projects: Projects(""),
	}
}

// Projects lists the nest. home defaults to the current user. Status is present or missing. Paths are not secret files.
func Projects(home string) []Project {
	if home == "" {
		home, _ = os.UserHomeDir()
	}
	rows := []Project{
		{"tcpsystems", "den", "Projects/harness-plan", "harness", ""},
		{"tcpsystems", "radar", "Projects/tcp-systems", "trackers, sources, prices, models", ""},
		{"foxcore-cloud", "fox", "Projects/org/foxcore-cloud/fox", "company panel", ""},
		{"foxcore-cloud", "foxd", "Projects/org/foxcore-cloud/foxd", "node agent", ""},
		{"foxcore-cloud", "floki", "Projects/org/foxcore-cloud/floki", "hypervisor agent", ""},
		{"foxcore-cloud", "ideology", "Projects/org/foxcore-cloud/ideology", "doctrine", ""},
		{"foxcore-cloud", "agent-mesh", "Projects/agent-mesh", "archived harness", ""},
	}
	for i := range rows {
		full := filepath.Join(home, rows[i].Path)
		rows[i].Path = full
		if st, err := os.Stat(full); err == nil && st.IsDir() {
			rows[i].Status = "present"
		} else {
			rows[i].Status = "missing"
		}
	}
	return rows
}

func readBranches(ctx context.Context, repo string) []Branch {
	out, err := run(ctx, repo, "for-each-ref", "--format=%(refname:short)|%(objectname:short)|%(subject)", "refs/heads")
	if err != nil {
		return nil
	}
	var rows []Branch
	for _, line := range strings.Split(out, "\n") {
		parts := strings.SplitN(line, "|", 3)
		if len(parts) < 3 || parts[0] == "" {
			continue
		}
		name := parts[0]
		kind := "line"
		switch {
		case name == "main":
			kind = "archive"
		case strings.HasPrefix(name, "section/"):
			kind = "section"
		case strings.HasPrefix(name, "r"):
			kind = "draft"
		}
		n := 0
		if count, err := run(ctx, repo, "rev-list", "--count", "develop.."+name); err == nil {
			n = atoi(strings.TrimSpace(count))
		}
		if kind == "draft" && n == 0 {
			continue
		}
		rows = append(rows, Branch{Name: name, SHA: parts[1], Subject: parts[2], Ahead: n, Kind: kind})
	}
	return rows
}

var probeCache = struct {
	sync.Mutex
	at     time.Time
	values map[string]string
}{values: map[string]string{}}

func probe(ctx context.Context, rawURL string) string {
	probeCache.Lock()
	if time.Since(probeCache.at) < 5*time.Second {
		if value, ok := probeCache.values[rawURL]; ok {
			probeCache.Unlock()
			return value
		}
	}
	probeCache.Unlock()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "curl", "-fsS", "-o", "/dev/null", "-w", "%{http_code}", rawURL)
	out, err := cmd.Output()
	if err != nil {
		return "down"
	}
	status := "down"
	if err == nil && strings.TrimSpace(string(out)) == "200" {
		status = "up"
	}
	probeCache.Lock()
	probeCache.values[rawURL] = status
	probeCache.at = time.Now()
	probeCache.Unlock()
	return status
}

func run(ctx context.Context, repo string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", append([]string{"-C", repo}, args...)...)
	out, err := cmd.Output()
	return string(out), err
}

func atoi(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return n
}
