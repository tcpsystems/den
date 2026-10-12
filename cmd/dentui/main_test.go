package main

import (
	"path/filepath"
	"testing"

	"github.com/tcpsystems/den/internal/initd"
)

func TestSeatKeyStaysOffWhenPartMissing(t *testing.T) {
	m := model{repo: t.TempDir(), board: initd.Default(t.TempDir(), filepath.Join(t.TempDir(), "den.sock"))}
	if m.enableOK("pages") {
		t.Fatal("pages is not a process in this build")
	}
	if !m.enableOK("seats") {
		t.Fatal("seats is a process in this build")
	}
}
