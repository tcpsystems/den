package orch

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func git(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		"GIT_AUTHOR_NAME=kivox",
		"GIT_AUTHOR_EMAIL=petar@petarmarkov.co.uk",
		"GIT_COMMITTER_NAME=kivox",
		"GIT_COMMITTER_EMAIL=petar@petarmarkov.co.uk",
	)
	var out, errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(errb.String())
		if msg == "" {
			msg = strings.TrimSpace(out.String())
		}
		return "", fmt.Errorf("git %s: %s", strings.Join(args, " "), msg)
	}
	return strings.TrimSpace(out.String()), nil
}

func head(dir string) (string, error) {
	return git(dir, "rev-parse", "--short", "HEAD")
}

func currentBranch(dir string) (string, error) {
	return git(dir, "rev-parse", "--abbrev-ref", "HEAD")
}

func porcelain(dir string) (string, error) {
	return git(dir, "status", "--porcelain")
}

// ownedDirty lists porcelain lines whose path is outside the seat's Owns prefixes.
// A clean tree returns empty. Unowned dirty paths are an error for the caller.
func ownedDirty(dir string, owns []string) (owned, foreign []string, err error) {
	raw, err := porcelain(dir)
	if err != nil {
		return nil, nil, err
	}
	if raw == "" {
		return nil, nil, nil
	}
	for _, line := range strings.Split(raw, "\n") {
		if len(line) < 4 {
			continue
		}
		path := strings.TrimSpace(line[3:])
		if i := strings.Index(path, " -> "); i >= 0 {
			path = path[i+4:]
		}
		if strings.Contains(path, "..") {
			foreign = append(foreign, path)
			continue
		}
		if ownsPath(owns, path) {
			owned = append(owned, path)
			continue
		}
		foreign = append(foreign, path)
	}
	return owned, foreign, nil
}

func ownsPath(owns []string, path string) bool {
	path = filepath.ToSlash(path)
	for _, o := range owns {
		o = filepath.ToSlash(strings.TrimSuffix(o, "/"))
		if path == o || strings.HasPrefix(path, o+"/") {
			return true
		}
	}
	return false
}
