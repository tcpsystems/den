package read

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestJailRejectsParent(t *testing.T) {
	root := t.TempDir()
	if _, err := Jail(root, "../outside"); err == nil {
		t.Fatal("parent path accepted")
	}
}

func TestNumbered(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "a.txt")
	if err := os.WriteFile(path, []byte("one\ntwo\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := Numbered(&buf, root, "a.txt"); err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	if !strings.Contains(got, "1→one\n") || !strings.Contains(got, "2→two\n") {
		t.Fatalf("numbered = %q", got)
	}
}

func TestBatchStopsOnEscape(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "a.txt"), []byte("ok\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	err := Batch(&buf, root, []string{"a.txt", "../../etc/passwd"})
	if err == nil {
		t.Fatal("batch accepted a path outside the root")
	}
}

func TestCondenseRelatedAndFlow(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "plan"), 0o755); err != nil {
		t.Fatal(err)
	}
	body := "# Lead\n\nSee [base](plan/BASE.md) and plan/nest.md.\n\n## Next\n\nfunc Run() {}\n"
	if err := os.WriteFile(filepath.Join(root, "note.md"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "plan", "BASE.md"), []byte("# base\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "plan", "nest.md"), []byte("# nest\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ov, err := Condense(root, "note.md")
	if err != nil {
		t.Fatal(err)
	}
	if ov.Lines != 7 || ov.Kind != "note" {
		t.Fatalf("overview = %+v", ov)
	}
	if len(ov.Heads) != 2 || ov.Heads[0] != "Lead" {
		t.Fatalf("heads = %#v", ov.Heads)
	}
	joined := strings.Join(ov.Related, ",")
	if !strings.Contains(joined, "plan/BASE.md") || !strings.Contains(joined, "plan/nest.md") {
		t.Fatalf("related = %#v", ov.Related)
	}
	var buf bytes.Buffer
	WriteOverview(&buf, ov)
	out := buf.String()
	if strings.Contains(out, "See [base]") {
		t.Fatal("overview copied the body")
	}
	if !strings.Contains(out, "related:") || !strings.Contains(out, "flow:") {
		t.Fatalf("overview missing links or flow:\n%s", out)
	}
}
