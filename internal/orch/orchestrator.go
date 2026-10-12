// orchestrator.go
package orch

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
)

// Orchestrator manages the swarm workflow and commit lifecycle
type Orchestrator struct {
    WaveID string
    Branch string
    RepoPath string
    WorktreeDir string
}

// CommitEntry represents a worker's contribution
type CommitEntry struct {
    SeatID string
    Role string
    Path string
    Message string
    Verified bool
}

// ProcessCommit validates and adds commit to queue
func (o *Orchestrator) ProcessCommit(seatID, message string) error {
    if !IsValidCommit(message) {
        return fmt.Errorf("invalid commit message format")
    }
    entry := CommitEntry{SeatID: seatID, Message: message, Verified: false}
    return o.VerifyCommit(entry)
}

// VerifyCommit runs tests and linting on worktree
func (o *Orchestrator) VerifyCommit(entry CommitEntry) error {
    cmd := exec.Command("go", "test", "./...")
    cmd.Dir = filepath.Join(o.WorktreeDir, entry.Path)
    if out, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("tests failed: %s", string(out))
    }
    return nil
}

// SquashAndPush merges verified commits into feature branch
func (o *Orchestrator) SquashAndPush(commitMsg string) error {
    cmd := exec.Command("git", "checkout", o.Branch)
    cmd.Dir = o.RepoPath
    if out, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("checkout branch: %s", string(out))
    }
    cmd = exec.Command("git", "pull", "--rebase")
    cmd.Dir = o.RepoPath
    if out, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("pull rebase: %s", string(out))
    }
    return nil
}

// GenerateCommitMessage creates consolidated commit message
func (o *Orchestrator) GenerateCommitMessage(messages []string) string {
    var types []string
    for _, m := range messages {
        types = append(types, CommitType(m))
    }
    return fmt.Sprintf("chore: wave %s merge - automated verification", o.WaveID)
}

// UpdateVersion updates VERSION file based on commit types
func (o *Orchestrator) UpdateVersion(commitTypes []string) error {
    breaking := false
    for _, t := range commitTypes {
        if strings.Contains(strings.ToLower(t), "breaking") {
            breaking = true
        }
    }
    current, err := LatestVersionTag(o.RepoPath)
    if err != nil {
        return err
    }
    newVer, err := BumpVersion(current, "feat", breaking)
    if err != nil {
        return err
    }
    return UpdateVersionFile(o.RepoPath, newVer)
}

// GenerateLedger creates documentation entry for research
func (o *Orchestrator) GenerateLedger(commitTypes []string, messages []string) string {
    now := time.Now().Format("2006-01-02")
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("## %s - %s\n", o.WaveID, now))
    for _, msg := range messages {
        sb.WriteString(fmt.Sprintf("- %s\n", msg))
    }
    return sb.String()
}

// WriteLedger saves commit ledger to research drafts
func (o *Orchestrator) WriteLedger(content string) error {
    ledgerPath := filepath.Join(o.RepoPath, "research", "drafts", "wave-ledger.md")
    if err := os.MkdirAll(filepath.Dir(ledgerPath), 0755); err != nil {
        return err
    }
    return os.WriteFile(ledgerPath, []byte(content+"\n"), 0644)
}

// InitWorktree creates isolated worktrees for each seat
func (o *Orchestrator) InitWorktree() error {
    if err := os.MkdirAll(o.WorktreeDir, 0755); err != nil {
        return err
    }
    cmd := exec.Command("git", "worktree", "add", o.WorktreeDir, "HEAD")
    cmd.Dir = o.RepoPath
    if out, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("worktree add: %s", string(out))
    }
    return nil
}

// PrintStatus shows current orchestrator state
func (o *Orchestrator) PrintStatus() {
    fmt.Printf("Orchestrator: %s on branch %s\n", o.WaveID, o.Branch)
    fmt.Printf("Worktree: %s\n", o.WorktreeDir)
    fmt.Printf("Repo: %s\n", o.RepoPath)
}
