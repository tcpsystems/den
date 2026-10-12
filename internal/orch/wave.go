package orch

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Open creates the wave branch from Base when it does not exist, writes the
// ledger, and adds one worktree per seat. Existing worktrees are reused.
func Open(w Wave) (Wave, error) {
	if w.Repo == "" || w.Branch == "" || w.ID == "" || w.Base == "" {
		return Wave{}, fmt.Errorf("wave needs id, branch, base, and repo")
	}
	if w.Worktrees == "" {
		w.Worktrees = filepath.Join(filepath.Dir(w.Repo), "harness-plan-wt")
	}
	if _, err := git(w.Repo, "rev-parse", "--verify", w.Base); err != nil {
		return Wave{}, err
	}
	if _, err := git(w.Repo, "rev-parse", "--verify", "refs/heads/"+w.Branch); err != nil {
		if _, err := git(w.Repo, "branch", w.Branch, w.Base); err != nil {
			return Wave{}, err
		}
	}
	if err := os.MkdirAll(w.Worktrees, 0o755); err != nil {
		return Wave{}, err
	}
	for i := range w.Seats {
		s := &w.Seats[i]
		if s.ID == "" || s.Role == "" || len(s.Owns) == 0 {
			return Wave{}, fmt.Errorf("seat %d needs id, role, and owns", i)
		}
		dir := filepath.Join(w.Worktrees, w.ID+"-"+s.ID)
		seatBranch := seatBranchName(w.Branch, s.ID)
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			s.Status = "ready"
			h, herr := head(dir)
			if herr != nil {
				return Wave{}, herr
			}
			s.Head = h
			continue
		}
		// One branch per seat. Git forbids two worktrees on the same branch.
		if _, err := git(w.Repo, "worktree", "add", "-B", seatBranch, dir, w.Branch); err != nil {
			return Wave{}, fmt.Errorf("seat %s: %w", s.ID, err)
		}
		s.Status = "ready"
		h, err := head(dir)
		if err != nil {
			return Wave{}, err
		}
		s.Head = h
	}
	if err := Save(w); err != nil {
		return Wave{}, err
	}
	return w, nil
}

// Commit records the seat's owned paths with a conventional message.
// Paths outside Owns are left unstaged and block the commit.
func Commit(repo, waveID, seatID, message string) (Wave, error) {
	if err := CheckMessage("", message); err != nil {
		return Wave{}, err
	}
	w, err := Load(repo, waveID)
	if err != nil {
		return Wave{}, err
	}
	s, err := w.seat(seatID)
	if err != nil {
		return Wave{}, err
	}
	if err := CheckMessage(s.Role, message); err != nil {
		return Wave{}, err
	}
	dir := filepath.Join(w.Worktrees, w.ID+"-"+s.ID)
	owned, foreign, err := ownedDirty(dir, s.Owns)
	if err != nil {
		return Wave{}, err
	}
	if len(foreign) > 0 {
		return Wave{}, fmt.Errorf("seat %s does not own %s", s.ID, strings.Join(foreign, ", "))
	}
	if len(owned) == 0 {
		return Wave{}, fmt.Errorf("seat %s has no changes in %s", s.ID, strings.Join(s.Owns, ", "))
	}
	for _, p := range owned {
		if _, err := git(dir, "add", "--", p); err != nil {
			return Wave{}, err
		}
	}
	if _, err := git(dir, "commit", "-m", message); err != nil {
		return Wave{}, err
	}
	h, err := head(dir)
	if err != nil {
		return Wave{}, err
	}
	s.Head = h
	s.Status = "committed"
	s.Note = message
	if err := Save(w); err != nil {
		return Wave{}, err
	}
	return w, nil
}

// Take merges one committed seat onto the wave branch in the main checkout
// after checks pass and the seat tree is clean. It does not push.
func Take(repo, waveID, seatID string, checks []Check) (Wave, error) {
	w, err := Load(repo, waveID)
	if err != nil {
		return Wave{}, err
	}
	s, err := w.seat(seatID)
	if err != nil {
		return Wave{}, err
	}
	dir := filepath.Join(w.Worktrees, w.ID+"-"+s.ID)
	owned, foreign, err := ownedDirty(dir, s.Owns)
	if err != nil {
		return Wave{}, err
	}
	if len(owned) > 0 || len(foreign) > 0 {
		return Wave{}, fmt.Errorf("seat %s still has uncommitted changes", s.ID)
	}
	seatBranch := seatBranchName(w.Branch, s.ID)
	br, err := currentBranch(dir)
	if err != nil {
		return Wave{}, err
	}
	if br != seatBranch && br != w.Branch {
		return Wave{}, fmt.Errorf("seat %s is on %s, want %s", s.ID, br, seatBranch)
	}
	for _, c := range checks {
		cmd := exec.Command(c.Argv[0], c.Argv[1:]...)
		cmd.Dir = dir
		out, err := cmd.CombinedOutput()
		if err != nil {
			s.Status = "feedback"
			s.Note = c.Name + " failed: " + strings.TrimSpace(string(out))
			_ = Save(w)
			return w, fmt.Errorf("%s", s.Note)
		}
	}
	mainBranch, err := currentBranch(repo)
	if err != nil {
		return Wave{}, err
	}
	if mainBranch != w.Branch {
		if _, err := git(repo, "checkout", w.Branch); err != nil {
			return Wave{}, err
		}
	}
	seatHead, err := git(dir, "rev-parse", "HEAD")
	if err != nil {
		return Wave{}, err
	}
	// Merge the seat branch onto the shared wave branch. Prefer ff; fall back to merge commit.
	if _, err := git(repo, "merge", "--ff-only", seatHead); err != nil {
		msg := fmt.Sprintf("orch: take seat %s", s.ID)
		if _, merr := git(repo, "merge", "--no-ff", "-m", msg, seatHead); merr != nil {
			s.Status = "feedback"
			s.Note = merr.Error()
			_ = Save(w)
			return w, merr
		}
	}
	h, err := head(repo)
	if err != nil {
		return Wave{}, err
	}
	s.Head = h
	s.Status = "merged"
	s.Note = "ff " + h
	if err := Save(w); err != nil {
		return Wave{}, err
	}
	return w, nil
}
