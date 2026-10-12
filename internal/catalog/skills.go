package catalog

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

// LoadSkills indexes every skills/<id>/SKILL.md under root.
// The body stays on disk. The index stores name, description, and path.
// A missing skills directory leaves the index unchanged.
func LoadSkills(idx *Index, root string) error {
	dir := filepath.Join(root, "skills")
	ents, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	names := make([]string, 0, len(ents))
	for _, ent := range ents {
		if ent.IsDir() {
			names = append(names, ent.Name())
		}
	}
	sort.Strings(names)
	for _, name := range names {
		path := filepath.Join(dir, name, "SKILL.md")
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		e, err := skillEntry(path, name)
		if err != nil {
			return err
		}
		if err := idx.Add(e); err != nil {
			return err
		}
	}
	return nil
}

func skillEntry(path, dirName string) (Entry, error) {
	name, desc, err := skillFront(path)
	if err != nil {
		return Entry{}, err
	}
	if name != dirName {
		return Entry{}, fmt.Errorf("skill %s: name %s does not match directory", dirName, name)
	}
	rel := filepath.ToSlash(filepath.Join("skills", dirName, "SKILL.md"))
	return Entry{
		Kind:    KindSkill,
		ID:      name,
		Title:   skillTitle(name),
		Summary: desc,
		Path:    rel,
		Needs:   []string{"skill:" + name},
	}, nil
}

func skillFront(path string) (name, description string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	if !sc.Scan() || strings.TrimSpace(sc.Text()) != "---" {
		return "", "", fmt.Errorf("%s: skill front matter must start with ---", path)
	}
	var key string
	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "---" {
			break
		}
		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
			if key == "" {
				return "", "", fmt.Errorf("%s: continuation before a field", path)
			}
			if key == "description" {
				description = strings.TrimSpace(description + " " + strings.TrimSpace(line))
			}
			continue
		}
		k, v, ok := strings.Cut(line, ":")
		if !ok {
			return "", "", fmt.Errorf("%s: front matter line %q has no key", path, line)
		}
		key = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		v = strings.Trim(v, `"'`)
		switch key {
		case "name":
			name = v
		case "description":
			description = v
		}
	}
	if err := sc.Err(); err != nil {
		return "", "", err
	}
	if name == "" || description == "" {
		return "", "", fmt.Errorf("%s: skill needs name and description", path)
	}
	return name, description, nil
}

func skillTitle(name string) string {
	parts := strings.Split(name, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		r := []rune(p)
		r[0] = unicode.ToUpper(r[0])
		parts[i] = string(r)
	}
	return strings.Join(parts, " ")
}

// SkillGrant is the scope a grant must name before this skill loads.
func SkillGrant(id string) string {
	return "skill:" + id
}
