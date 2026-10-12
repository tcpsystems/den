package catalog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Manifest is plugins/<id>/plugin.json. It only adds catalog rows.
type Manifest struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Summary string  `json:"summary"`
	Entries []Entry `json:"entries"`
}

// LoadPlugin reads one plugin manifest and adds its rows plus the plugin itself.
func LoadPlugin(idx *Index, dir string) error {
	b, err := os.ReadFile(filepath.Join(dir, "plugin.json"))
	if err != nil {
		return err
	}
	var m Manifest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	if m.ID == "" || m.Title == "" {
		return fmt.Errorf("plugin %s needs id and title", dir)
	}
	if err := idx.Add(Entry{Kind: KindPlugin, ID: m.ID, Title: m.Title, Summary: m.Summary, Path: dir}); err != nil {
		return err
	}
	for _, e := range m.Entries {
		if e.Path != "" && !filepath.IsAbs(e.Path) {
			e.Path = filepath.Join(dir, e.Path)
		}
		if err := idx.Add(e); err != nil {
			return fmt.Errorf("plugin %s: %w", m.ID, err)
		}
	}
	return nil
}

// LoadPlugins reads every directory under root that contains plugin.json.
func LoadPlugins(idx *Index, root string) error {
	ents, err := os.ReadDir(root)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, ent := range ents {
		if !ent.IsDir() {
			continue
		}
		dir := filepath.Join(root, ent.Name())
		if _, err := os.Stat(filepath.Join(dir, "plugin.json")); err != nil {
			continue
		}
		if err := LoadPlugin(idx, dir); err != nil {
			return err
		}
	}
	return nil
}
