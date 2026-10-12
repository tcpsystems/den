package cover

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEveryPackageHasATest(t *testing.T) {
	root := repoRoot(t)
	var missing []string
	for _, dir := range []string{"internal", "cmd"} {
		filepath.WalkDir(filepath.Join(root, dir), func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			matches, _ := filepath.Glob(filepath.Join(filepath.Dir(path), "*_test.go"))
			if len(matches) == 0 {
				missing = append(missing, strings.TrimPrefix(filepath.Dir(path), root+"/"))
			}
			return nil
		})
	}
	if len(missing) > 0 {
		t.Fatalf("packages with no test: %s", strings.Join(unique(missing), ", "))
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("go.mod not found")
		}
		dir = parent
	}
}

func unique(in []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}
