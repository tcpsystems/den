// semantic_commit.go
package orch

import (
    "fmt"
    "strings"
)

// Conventional commit types
var commitTypes = map[string]string{
    "feat": "new feature",
    "fix": "bug fix",
    "docs": "documentation",
    "style": "formatting",
    "refactor": "code refactoring",
    "perf": "performance improvement",
    "test": "adding tests",
    "chore": "maintenance",
    "build": "build system",
    "ci": "ci configuration",
    "revert": "revert changes",
}

// SemanticCommit validates and returns normalized commit message
func SemanticCommit(msg string) (string, error) {
    // Strip leading/trailing whitespace
    msg = strings.TrimSpace(msg)
    if msg == "" {
        return "", fmt.Errorf("empty commit message")
    }

    // Split into scope and description
    parts := strings.SplitN(msg, ":", 2)
    var scope, description string
    if len(parts) == 2 {
        scope = strings.TrimSpace(parts[0])
        description = strings.TrimSpace(parts[1])
    } else {
        description = strings.TrimSpace(parts[0])
    }

    // Extract type and optional scope
    words := strings.Fields(description)
    if len(words) == 0 {
        return "", fmt.Errorf("no description after type")
    }

    // Find commit type
    var commitType string
    for t := range commitTypes {
        if strings.HasPrefix(words[0], t+"(") || strings.HasPrefix(words[0], t+":") || t == words[0] {
            commitType = t
            break
        }
    }

    if commitType == "" {
        return "", fmt.Errorf("invalid commit type: %s", words[0])
    }

    // Reconstruct normalized message
    normalized := commitType
    if scope != "" {
        normalized += "(" + scope + ")"
    }
    normalized += ": " + strings.TrimSpace(strings.TrimPrefix(description, words[0]))

    return normalized, nil
}

// IsValidCommit checks if message follows conventional commit
func IsValidCommit(msg string) bool {
    _, err := SemanticCommit(msg)
    return err == nil
}

// CommitType extracts the type from a valid commit message
func CommitType(msg string) string {
    if !IsValidCommit(msg) {
        return ""
    }
    words := strings.Fields(strings.SplitN(msg, ":", 2)[1])
    for t := range commitTypes {
        if strings.HasPrefix(words[0], t+"(") || strings.HasPrefix(words[0], t+":") || t == words[0] {
            return t
        }
    }
    return ""
}

// VersionBump determines if commit warrants version bump
func VersionBump(commitType string, breaking bool) string {
    if breaking {
        return "major"
    }
    switch commitType {
    case "feat":
        return "minor"
    case "fix", "docs", "style", "refactor", "perf", "test", "chore", "build", "ci":
        return "patch"
    default:
        return ""
    }
}
