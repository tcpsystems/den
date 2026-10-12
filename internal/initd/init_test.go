package initd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndEnable(t *testing.T) {
	repo := t.TempDir()
	board := Default(repo, "/tmp/den.sock")
	if err := Save(board); err != nil {
		t.Fatal(err)
	}
	got, err := Load(repo, "/tmp/den.sock")
	if err != nil {
		t.Fatal(err)
	}
	if got.Parts[0].Status != "on" || got.Parts[2].Status != "off" {
		t.Fatalf("board = %+v", got.Parts)
	}
	if Enable(&got, "web") {
		t.Fatal("web has no process in this build")
	}
	if got.Parts[2].Status != "off" {
		t.Fatal("web was marked on")
	}
	if _, err := os.Stat(filepath.Join(repo, ".den", "init.json")); err != nil {
		t.Fatal(err)
	}
}

func TestLoadMissing(t *testing.T) {
	got, err := Load(t.TempDir(), "/tmp/den.sock")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Parts) != 6 {
		t.Fatalf("parts = %d", len(got.Parts))
	}
}
