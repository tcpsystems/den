package store

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func sample(id, key string) SectionRecord {
	return SectionRecord{
		Section:  SectionID(id),
		Function: "append one section",
		Read:     []string{"sections/store/LEAD.md"},
		Do:       []string{"put", "query"},
		Write:    key,
		Next:     NextRef{Kind: "id", ID: "store-next"},
		Rail:     "",
	}
}

func mustOpen(t *testing.T) *Store {
	t.Helper()
	s, err := open(t.TempDir(), t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func TestPutThenQuery(t *testing.T) {
	s := mustOpen(t)
	rec := sample("store-1", "kb/store-1")
	if err := s.Put(rec); err != nil {
		t.Fatal(err)
	}
	got, err := s.Query("store-1")
	if err != nil {
		t.Fatal(err)
	}
	if got.Section != rec.Section || got.Write != rec.Write || got.Function != rec.Function {
		t.Fatalf("query mismatch: %+v", got)
	}
	cur, err := s.Query("")
	if err != nil {
		t.Fatal(err)
	}
	if cur.Section != rec.Section {
		t.Fatalf("current = %s", cur.Section)
	}
}

func TestLaterSectionSupersedesWithoutDelete(t *testing.T) {
	s := mustOpen(t)
	a := sample("store-a", "kb/a")
	b := sample("store-b", "kb/b")
	if err := s.Put(a); err != nil {
		t.Fatal(err)
	}
	if err := s.Put(b); err != nil {
		t.Fatal(err)
	}
	cur, err := s.Query("")
	if err != nil {
		t.Fatal(err)
	}
	if cur.Section != "store-b" {
		t.Fatalf("current = %s", cur.Section)
	}
	old, err := s.Query("store-a")
	if err != nil {
		t.Fatal(err)
	}
	if old.Write != "kb/a" {
		t.Fatalf("earlier row missing: %+v", old)
	}
}

func TestPutRefusesEmptyWriteAndMissingFields(t *testing.T) {
	s := mustOpen(t)
	rec := sample("store-1", "")
	if err := s.Put(rec); !errors.Is(err, ErrEmptyWrite) {
		t.Fatalf("empty write: %v", err)
	}
	rec = sample("", "kb/x")
	rec.Section = ""
	if err := s.Put(rec); !errors.Is(err, ErrMissingFields) {
		t.Fatalf("missing section: %v", err)
	}
}

func TestPutRefusesSecondStoreAndRewrite(t *testing.T) {
	s := mustOpen(t)
	rec := sample("store-1", "kb/store-1")
	if err := s.Put(rec); err != nil {
		t.Fatal(err)
	}
	if err := s.Put(rec); !errors.Is(err, ErrSecondStore) {
		t.Fatalf("second store: %v", err)
	}
	other := sample("store-2", "kb/store-1")
	if err := s.Put(other); !errors.Is(err, ErrRewriteCorpus) {
		t.Fatalf("rewrite: %v", err)
	}
}

func TestAppendRefusalKeepsCurrent(t *testing.T) {
	s := mustOpen(t)
	rec := sample("store-1", "kb/store-1")
	if err := s.Put(rec); err != nil {
		t.Fatal(err)
	}
	err := s.AppendRefusal(RefusalRow{
		Task:    "t1",
		State:   "cut",
		Refused: NextRef{Kind: "id", ID: "store-next"},
		Reason:  "lead cut",
	})
	if err != nil {
		t.Fatal(err)
	}
	cur, err := s.Query("")
	if err != nil {
		t.Fatal(err)
	}
	if cur.Section != "store-1" {
		t.Fatalf("cut moved current to %s", cur.Section)
	}
	if err := s.AppendRefusal(RefusalRow{State: "stop", Reason: "no"}); !errors.Is(err, ErrCutNotStop) {
		t.Fatalf("stop cut: %v", err)
	}
}

func TestAppendProgress(t *testing.T) {
	s := mustOpen(t)
	err := s.AppendProgress(ProgressRow{
		At:       time.Now().UTC(),
		Task:     "t1",
		State:    "stored",
		Artifact: "kb/store-1",
		Next:     "store-next",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFastLocal(t *testing.T) {
	s := mustOpen(t)
	dir := s.fastDir()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "note")
	if err := os.WriteFile(path, []byte("local fast"), 0o644); err != nil {
		t.Fatal(err)
	}
	notes, err := s.ReadFast()
	if err != nil {
		t.Fatal(err)
	}
	if len(notes) != 1 || notes[0].Value != "local fast" || notes[0].Tier != "fast" {
		t.Fatalf("fast notes: %+v", notes)
	}
	if notes[0].Source != path {
		t.Fatalf("source %s", notes[0].Source)
	}
}

func TestBackupIsNotTheReadPath(t *testing.T) {
	if _, err := Open(BackupRoot); !errors.Is(err, ErrBackupRead) {
		t.Fatalf("open backup: %v", err)
	}
	if _, err := Open("/opt/harness-memory/fast"); !errors.Is(err, ErrBackupRead) {
		t.Fatalf("open harness-memory: %v", err)
	}
	if _, err := ReadBackup(); !errors.Is(err, ErrBackupRead) {
		t.Fatalf("ReadBackup: %v", err)
	}
	s := mustOpen(t)
	s.home = "/opt/harness-memory"
	if _, err := s.ReadFast(); !errors.Is(err, ErrBackupRead) {
		t.Fatalf("ReadFast backup home: %v", err)
	}
	if _, err := s.ReadRecent(); !errors.Is(err, ErrBackupRead) {
		t.Fatalf("ReadRecent backup home: %v", err)
	}
}

func TestEncodeBodySevenFieldsOnly(t *testing.T) {
	body, err := EncodeBody(sample("store-1", "kb/store-1"))
	if err != nil {
		t.Fatal(err)
	}
	rec, err := DecodeBody(body)
	if err != nil {
		t.Fatal(err)
	}
	if rec.Next.Kind != "id" || rec.Next.ID != "store-next" {
		t.Fatalf("next: %+v", rec.Next)
	}
	if _, err := DecodeBody([]byte(`{"Section":"x","Function":"f","Read":[],"Do":[],"Write":"k","Next":"","Rail":"","Extra":1}`)); err == nil {
		t.Fatal("eighth field allowed")
	}
}

func TestChildWrite(t *testing.T) {
	parent := sample("p", "kb/p")
	ok := parent
	if err := ChildWrite(parent, ok); err != nil {
		t.Fatal(err)
	}
	bad := parent
	bad.Write = "kb/other"
	if err := ChildWrite(parent, bad); !errors.Is(err, ErrChildWrite) {
		t.Fatalf("child write: %v", err)
	}
}

func TestQueryMissing(t *testing.T) {
	s := mustOpen(t)
	if _, err := s.Query("nope"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing: %v", err)
	}
}
