// semantic_commit_test.go
package orch

import (
    "fmt"
    "testing"
)

func TestSemanticCommit(t *testing.T) {
    tests := []struct {
        msg string
        want string
        err  bool
    }{{
        {"feat: add new authentication flow", "feat: add new authentication flow", false},
        {"fix: resolve memory leak in embedder", "fix: resolve memory leak in embedder", false},
        {"docs: update API documentation", "docs: update API documentation", false},
        {"refactor: clean up permission module", "refactor: clean up permission module", false},
        {"feat(authentication): add OAuth support", "feat(authentication): add OAuth support", false},
        {"invalid type", "", true},
        {"", "", true},
        {"feat", "feat: ", false},
    }}

    for _, tt := range tests {
        got, err := SemanticCommit(tt.msg)
        if tt.err {
            if err == nil {
                t.Errorf("SemanticCommit(%q) expected error, got none", tt.msg)
            }
        } else {
            if err != nil {
                t.Errorf("SemanticCommit(%q) unexpected error: %v", tt.msg, err)
            } else if got != tt.want {
                t.Errorf("SemanticCommit(%q) = %q, want %q", tt.msg, got, tt.want)
            }
        }
    }
}

func TestVersionBump(t *testing.T) {
    tests := []struct {
        current string
        commit  string
        breaking bool
        want    string
    }{{
        {"1.2.3", "feat", false, "1.3.0"},
        {"1.2.3", "fix", false, "1.2.4"},
        {"1.2.3", "feat", true, "2.0.0"},
        {"1.2.3-alpha", "fix", false, "1.2.4"},
        {"1.2.3", "revert", false, "1.2.3"},
    }}

    for _, tt := range tests {
        got, err := BumpVersion(tt.current, tt.commit, tt.breaking)
        if err != nil {
            t.Errorf("BumpVersion(%s, %s, %v) error: %v", tt.current, tt.commit, tt.breaking, err)
        } else if got != tt.want {
        }
    }
}
