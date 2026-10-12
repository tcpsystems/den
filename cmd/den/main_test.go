package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tcpsystems/den/internal/initd"
)

func TestSessionWritesBoardBeforeStartingProcesses(t *testing.T) {
	repo := t.TempDir()
	home := t.TempDir()
	t.Setenv("HOME", home)
	if err := os.Chdir(repo); err != nil {
		t.Fatal(err)
	}
	sock := filepath.Join(home, ".local", "share", "harness", "den.sock")
	board, err := initd.Load(repo, sock)
	if err != nil {
		t.Fatal(err)
	}
	board.Repo = repo
	board.Sock = sock
	if err := initd.Save(board); err != nil {
		t.Fatal(err)
	}
	loaded, err := initd.Load(repo, sock)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Sock != sock || !initd.Present("controller") {
		t.Fatalf("%+v", loaded)
	}
}
