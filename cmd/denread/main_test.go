package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tcpsystems/den/internal/read"
)

func TestReadStaysInRoot(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "a.txt"), []byte("one"), 0o644); err != nil {
		t.Fatal(err)
	}
	var b strings.Builder
	if err := read.Numbered(&b, root, "a.txt"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(b.String(), "1→one") {
		t.Fatal(b.String())
	}
	if err := read.Numbered(&b, root, "../outside"); err == nil {
		t.Fatal("parent path accepted")
	}
}
