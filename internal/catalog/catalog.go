package catalog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Entry is one indexed thing. Body stays on disk. The index stores the path.
type Entry struct {
	Kind    Kind     `json:"kind"`
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Path    string   `json:"path,omitempty"`
	Needs   []string `json:"needs,omitempty"`
}

// Index is the runtime catalog. Plugins append entries. They do not replace kinds.
type Index struct {
	Entries []Entry `json:"entries"`
}

// Builtin is the index this repo ships before plugins load.
func Builtin() Index {
	return Index{Entries: []Entry{
		{Kind: KindRole, ID: "researcher", Title: "Researcher", Summary: "Writes sourced notes. Commit type docs."},
		{Kind: KindRole, ID: "architect", Title: "Architect", Summary: "Writes the plan for one slice. Commit type docs."},
		{Kind: KindRole, ID: "builder", Title: "Builder", Summary: "Writes code under owned paths. Commit type feat or fix."},
		{Kind: KindRole, ID: "reviewer", Title: "Reviewer", Summary: "Runs the checks and writes the verdict. Commit type test."},
		{Kind: KindTool, ID: "denread", Title: "denread", Summary: "Numbered local file read inside the path jail.", Path: "cmd/denread"},
		{Kind: KindTool, ID: "denfetch", Title: "denfetch", Summary: "Public page read. No credentials.", Path: "cmd/denfetch"},
		{Kind: KindTool, ID: "denorch", Title: "denorch", Summary: "Local git swarm. One branch per seat. Does not push.", Path: "cmd/denorch"},
		{Kind: KindMCP, ID: "omniroute", Title: "OmniRoute", Summary: "Gateway MCP. Models and usage. A grant must name it."},
		{Kind: KindMCP, ID: "tinyfish", Title: "TinyFish", Summary: "Search and public page fetch. A grant must name it."},
		{Kind: KindModel, ID: "auto/coding", Title: "Coding route", Summary: "Lead route through OmniRoute."},
		{Kind: KindModel, ID: "auto/coding:cheap", Title: "Cheap coding route", Summary: "Titles, labels, and checks."},
		{Kind: KindMemory, ID: "session-ledger", Title: "Session ledger", Summary: "Flat markdown and json under the repo. Survives an embedder rebuild."},
		{Kind: KindNiche, ID: "harness-runtime", Title: "Harness runtime", Summary: "denctl, catalog, seats, grants."},
		{Kind: KindIdentity, ID: "den-1", Title: "Den", Summary: "The seat that uses this harness. Id den-1, name Den, avatar den."},
		{Kind: KindSkill, ID: "use-harness", Title: "Use harness", Summary: "Route one reference, then one skill body.", Path: "skills/use-harness/SKILL.md"},
		{Kind: KindProfile, ID: "lead", Title: "Lead", Summary: "den-1, builder, harness-runtime, auto/coding, use-harness, session-ledger."},
	}}
}

// Lead is the profile a session binds so this harness can staff itself.
// MCP stays empty. A grant names a server later. None is required.
func Lead() Profile {
	return Profile{
		ID: "lead", Identity: "den-1", Role: "builder", Niche: "harness-runtime",
		Model: "auto/coding", Skills: []string{"use-harness"}, Memory: "session-ledger",
		Avatar: "den",
	}
}

// Load reads index.json. A missing file returns Builtin.
func Load(dir string) (Index, error) {
	b, err := os.ReadFile(filepath.Join(dir, "index.json"))
	if os.IsNotExist(err) {
		return Builtin(), nil
	}
	if err != nil {
		return Index{}, err
	}
	var idx Index
	if err := json.Unmarshal(b, &idx); err != nil {
		return Index{}, err
	}
	return idx, nil
}

// Save writes index.json.
func Save(dir string, idx Index) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "index.json"), append(b, '\n'), 0o644)
}

// Add inserts an entry. The same kind and id replaces the old row.
func (idx *Index) Add(e Entry) error {
	if err := e.validate(); err != nil {
		return err
	}
	for i := range idx.Entries {
		if idx.Entries[i].Kind == e.Kind && idx.Entries[i].ID == e.ID {
			idx.Entries[i] = e
			return nil
		}
	}
	idx.Entries = append(idx.Entries, e)
	return nil
}

// ByKind returns entries of one kind, sorted by id.
func (idx Index) ByKind(k Kind) []Entry {
	var out []Entry
	for _, e := range idx.Entries {
		if e.Kind == k {
			out = append(out, e)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

func (e Entry) validate() error {
	if e.Kind == "" || e.ID == "" || e.Title == "" || e.Summary == "" {
		return fmt.Errorf("entry needs kind, id, title, and summary")
	}
	if strings.Contains(e.ID, " ") || strings.Contains(string(e.Kind), " ") {
		return fmt.Errorf("kind and id are single tokens")
	}
	switch e.Kind {
	case KindIdentity, KindRole, KindNiche, KindProfile, KindSkill, KindTool, KindMCP, KindPlugin, KindMemory, KindModel:
		return nil
	default:
		return fmt.Errorf("unknown kind %s", e.Kind)
	}
}

// Profile is the bundle the session orchestrator binds before a seat starts.
type Profile struct {
	ID       string   `json:"id"`
	Identity string   `json:"identity"`
	Role     string   `json:"role"`
	Niche    string   `json:"niche"`
	Model    string   `json:"model"`
	Skills   []string `json:"skills"`
	MCP      []string `json:"mcp"`
	Memory   string   `json:"memory"`
	Avatar   string   `json:"avatar,omitempty"`
}

// Bind checks that every named id exists in the index.
func Bind(idx Index, p Profile) error {
	if p.ID == "" || p.Identity == "" || p.Role == "" || p.Niche == "" || p.Model == "" {
		return fmt.Errorf("profile %s needs id, identity, role, niche, and model", p.ID)
	}
	need := []struct {
		k  Kind
		id string
	}{
		{KindIdentity, p.Identity},
		{KindRole, p.Role},
		{KindNiche, p.Niche},
		{KindModel, p.Model},
	}
	if p.Memory != "" {
		need = append(need, struct {
			k  Kind
			id string
		}{KindMemory, p.Memory})
	}
	for _, id := range p.Skills {
		need = append(need, struct {
			k  Kind
			id string
		}{KindSkill, id})
	}
	for _, id := range p.MCP {
		need = append(need, struct {
			k  Kind
			id string
		}{KindMCP, id})
	}
	for _, n := range need {
		if !idx.has(n.k, n.id) {
			return fmt.Errorf("profile %s: %s %s is not in the catalog", p.ID, n.k, n.id)
		}
	}
	return nil
}

func (idx Index) has(k Kind, id string) bool {
	for _, e := range idx.Entries {
		if e.Kind == k && e.ID == id {
			return true
		}
	}
	return false
}
