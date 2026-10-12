package catalog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadSkillsIndexesEveryFile(t *testing.T) {
	root := repoRoot(t)
	idx := Index{}
	if err := LoadSkills(&idx, root); err != nil {
		t.Fatal(err)
	}
	onDisk := skillFiles(t, root)
	got := idx.ByKind(KindSkill)
	if len(got) != len(onDisk) {
		t.Fatalf("indexed %d skills, disk has %d", len(got), len(onDisk))
	}
	byID := map[string]Entry{}
	for _, e := range got {
		byID[e.ID] = e
		if e.Path != "skills/"+e.ID+"/SKILL.md" {
			t.Fatalf("%s path %s", e.ID, e.Path)
		}
		if e.Summary == "" || e.Title == "" {
			t.Fatalf("%s missing title or summary", e.ID)
		}
		if len(e.Needs) != 1 || e.Needs[0] != SkillGrant(e.ID) {
			t.Fatalf("%s grant %v", e.ID, e.Needs)
		}
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(e.Path))); err != nil {
			t.Fatal(err)
		}
	}
	for _, id := range onDisk {
		if _, ok := byID[id]; !ok {
			t.Fatalf("missing %s", id)
		}
	}
	git, ok := byID["use-git"]
	if !ok || !strings.Contains(git.Summary, "Does not push") {
		t.Fatalf("use-git summary: %+v", git)
	}
}

func TestLoadSkillsReplacesBuiltinRows(t *testing.T) {
	root := repoRoot(t)
	idx := Builtin()
	before := len(idx.ByKind(KindSkill))
	if err := LoadSkills(&idx, root); err != nil {
		t.Fatal(err)
	}
	got := idx.ByKind(KindSkill)
	if len(got) <= before {
		t.Fatalf("skills %d, builtin had %d", len(got), before)
	}
	for _, e := range got {
		if !strings.HasPrefix(e.Path, "skills/") || !strings.HasSuffix(e.Path, "/SKILL.md") {
			t.Fatalf("stale path %s", e.Path)
		}
	}
	if err := Bind(idx, Lead()); err != nil {
		t.Fatal(err)
	}
}

func TestLoadSkillsRejectsNameMismatch(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "skills", "one")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := "---\nname: other\ndescription: Not this directory.\n---\n\n# Other\n"
	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	idx := Index{}
	if err := LoadSkills(&idx, root); err == nil {
		t.Fatal("want name mismatch")
	}
}

func TestLoadSkillsMissingDir(t *testing.T) {
	idx := Builtin()
	n := len(idx.Entries)
	if err := LoadSkills(&idx, t.TempDir()); err != nil {
		t.Fatal(err)
	}
	if len(idx.Entries) != n {
		t.Fatalf("entries %d want %d", len(idx.Entries), n)
	}
}

func skillFiles(t *testing.T, root string) []string {
	t.Helper()
	ents, err := os.ReadDir(filepath.Join(root, "skills"))
	if err != nil {
		t.Fatal(err)
	}
	var ids []string
	for _, ent := range ents {
		if !ent.IsDir() {
			continue
		}
		if _, err := os.Stat(filepath.Join(root, "skills", ent.Name(), "SKILL.md")); err != nil {
			continue
		}
		ids = append(ids, ent.Name())
	}
	if len(ids) < 2 {
		t.Fatalf("expected skills on disk, got %d", len(ids))
	}
	return ids
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("go.mod not found")
		}
		dir = parent
	}
}
