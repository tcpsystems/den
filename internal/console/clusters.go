package console

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Cluster is one worker fan-out directory under the harness wave scratch
// space. The page shows it so a watching human sees work move.
type Cluster struct {
	Task      string `json:"task"`
	Workers   int    `json:"workers"`
	Consensus bool   `json:"consensus"`
	Note      string `json:"note"`
}

// Clusters scans the wave scratch directory the orchestrator uses for
// worker outputs. A missing directory is an empty list, not an error.
func Clusters(dir string) []Cluster {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var out []Cluster
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		c := Cluster{Task: e.Name()}
		files, err := os.ReadDir(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		for _, f := range files {
			name := f.Name()
			if strings.HasPrefix(name, "worker") && strings.HasSuffix(name, ".txt") {
				c.Workers++
			}
			if name == "SUMMARY.md" {
				c.Consensus = true
				if b, err := os.ReadFile(filepath.Join(dir, e.Name(), name)); err == nil {
					c.Note = strings.TrimSpace(strings.SplitN(string(b), "\n", 2)[0])
				}
			}
		}
		if c.Note == "" {
			c.Note = "running"
		}
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Task < out[j].Task })
	return out
}
