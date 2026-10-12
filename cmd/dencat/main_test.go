package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/tcpsystems/den/internal/catalog"
)

func TestCatalogPrintsLead(t *testing.T) {
	idx := catalog.Builtin()
	if err := catalog.LoadPlugins(&idx, filepath.Join(t.TempDir(), "plugins")); err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(idx)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(b, []byte(`"den-1"`)) {
		t.Fatalf("%s", b)
	}
	if catalog.Terms == "" {
		t.Fatal("terms empty")
	}
	_ = os.Stdout
}
