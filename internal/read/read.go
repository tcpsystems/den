// Package read prints local files for a seat.
//
// A path must stay inside the root. The three shapes are numbered lines,
// a batch of files, and a content-light overview that names related files.
package read

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	maxLines    = 4000
	maxOverview = 240
	maxRelated  = 12
)

// Refused is a path that left the root.
type Refused struct{ Path string }

func (e Refused) Error() string { return "read: path left the root: " + e.Path }

// Jail returns the absolute path of rel inside root.
func Jail(root, rel string) (string, error) {
	base, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	base = filepath.Clean(base)
	candidate := rel
	if !filepath.IsAbs(rel) {
		candidate = filepath.Join(base, rel)
	}
	candidate, err = filepath.Abs(candidate)
	if err != nil {
		return "", err
	}
	candidate = filepath.Clean(candidate)
	if candidate != base && !strings.HasPrefix(candidate, base+string(os.PathSeparator)) {
		return "", Refused{Path: rel}
	}
	return candidate, nil
}

// Numbered writes every line of path with a 1-based line number.
func Numbered(w io.Writer, root, rel string) error {
	path, err := Jail(root, rel)
	if err != nil {
		return err
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	n := 0
	for sc.Scan() {
		n++
		if n > maxLines {
			fmt.Fprintf(w, "... truncated after %d lines\n", maxLines)
			return nil
		}
		fmt.Fprintf(w, "%d→%s\n", n, sc.Text())
	}
	return sc.Err()
}

// Batch writes numbered lines for each rel, in order.
func Batch(w io.Writer, root string, rels []string) error {
	for i, rel := range rels {
		if i > 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "# %s\n", rel)
		if err := Numbered(w, root, rel); err != nil {
			return fmt.Errorf("%s: %w", rel, err)
		}
	}
	return nil
}

// Overview is a content-light picture of one file.
type Overview struct {
	Path    string
	Kind    string
	Lines   int
	Heads   []string
	Symbols []string
	Related []string
}

var (
	heading = regexp.MustCompile(`^(#{1,3}) (.+)$`)
	gofunc  = regexp.MustCompile(`^func (\([^)]+\) )?[A-Z][A-Za-z0-9]*`)
	pydef   = regexp.MustCompile(`^(def|class) [A-Za-z_]`)
	linkRe  = regexp.MustCompile(`\[[^\]]+\]\(([^)]+)\)`)
	pathRe  = regexp.MustCompile(`(?:^|[\s"'(])((?:[A-Za-z0-9_.-]+/)+[A-Za-z0-9_.-]+)`)
)

// Condense reads path and returns headings, symbols, and related files
// that exist inside root. It does not copy the body.
func Condense(root, rel string) (Overview, error) {
	path, err := Jail(root, rel)
	if err != nil {
		return Overview{}, err
	}
	f, err := os.Open(path)
	if err != nil {
		return Overview{}, err
	}
	defer f.Close()
	ov := Overview{Path: rel, Kind: kindOf(rel)}
	seen := map[string]bool{}
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		ov.Lines++
		if ov.Lines > maxOverview && len(ov.Heads) > 0 && len(ov.Symbols) > 0 {
			continue
		}
		line := sc.Text()
		if m := heading.FindStringSubmatch(line); m != nil && len(ov.Heads) < 24 {
			ov.Heads = append(ov.Heads, strings.Repeat("  ", len(m[1])-1)+m[2])
		}
		if len(ov.Symbols) < 24 && (gofunc.MatchString(line) || pydef.MatchString(strings.TrimSpace(line))) {
			ov.Symbols = append(ov.Symbols, strings.TrimSpace(line))
		}
		if len(ov.Related) >= maxRelated {
			continue
		}
		for _, raw := range mentions(line) {
			if seen[raw] || raw == rel {
				continue
			}
			if _, err := Jail(root, raw); err != nil {
				continue
			}
			candidate := raw
			if !filepath.IsAbs(raw) {
				candidate = filepath.Join(filepath.Dir(path), raw)
				if _, err := os.Stat(candidate); err != nil {
					candidate = filepath.Join(root, raw)
				}
			}
			if st, err := os.Stat(candidate); err != nil || st.IsDir() {
				continue
			}
			seen[raw] = true
			ov.Related = append(ov.Related, raw)
			if len(ov.Related) >= maxRelated {
				break
			}
		}
	}
	if err := sc.Err(); err != nil {
		return Overview{}, err
	}
	sort.Strings(ov.Related)
	return ov, nil
}

func kindOf(rel string) string {
	switch strings.ToLower(filepath.Ext(rel)) {
	case ".go":
		return "go"
	case ".py":
		return "python"
	case ".md":
		return "note"
	case ".ts", ".tsx", ".js":
		return "script"
	default:
		return "text"
	}
}

func mentions(line string) []string {
	var out []string
	add := func(p string) {
		p = strings.TrimRight(p, ".,;:)")
		if p == "" || strings.Contains(p, "://") {
			return
		}
		out = append(out, p)
	}
	for _, m := range linkRe.FindAllStringSubmatch(line, -1) {
		p := m[1]
		if i := strings.IndexAny(p, "#? "); i >= 0 {
			p = p[:i]
		}
		add(p)
	}
	for _, m := range pathRe.FindAllStringSubmatch(line, -1) {
		add(m[1])
	}
	return out
}

// WriteOverview prints a short picture. It does not print the file body.
func WriteOverview(w io.Writer, ov Overview) {
	fmt.Fprintf(w, "# %s\nkind: %s\nlines: %d\n", ov.Path, ov.Kind, ov.Lines)
	if len(ov.Heads) > 0 {
		fmt.Fprintln(w, "flow:")
		for _, h := range ov.Heads {
			fmt.Fprintf(w, "- %s\n", h)
		}
	}
	if len(ov.Symbols) > 0 {
		fmt.Fprintln(w, "symbols:")
		for _, s := range ov.Symbols {
			fmt.Fprintf(w, "- %s\n", s)
		}
	}
	if len(ov.Related) > 0 {
		fmt.Fprintln(w, "related:")
		for _, r := range ov.Related {
			fmt.Fprintf(w, "- %s\n", r)
		}
	}
}
