// semantic_version.go
package orch

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

// Version represents semantic version components
type Version struct {
    Major int
    Minor int
    Patch int
    PreRelease string
    BuildMetadata string
}

// ParseVersion extracts version components from tag
func ParseVersion(tag string) (*Version, error) {
    // Try parsing manual format
    parts := strings.Split(strings.TrimPrefix(tag, "v"), ".")
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid version format: %s", tag)
    }
    v := &Version{}
    if _, err := fmt.Sscanf(parts[0], "%d", &v.Major); err != nil {
        return nil, fmt.Errorf("invalid major version: %s", parts[0])
    }
    if _, err := fmt.Sscanf(parts[1], "%d", &v.Minor); err != nil {
        return nil, fmt.Errorf("invalid minor version: %s", parts[1])
    }
    if _, err := fmt.Sscanf(parts[2], "%d", &v.Patch); err != nil {
        return nil, fmt.Errorf("invalid patch version: %s", parts[2])
    }
    if len(parts) > 3 {
        v.PreRelease = strings.Join(parts[3:], ".")
    }
    return v, nil
}

// BumpVersion increments version based on commit type and breaking change
func BumpVersion(current string, commitType string, breaking bool) (string, error) {
    v, err := ParseVersion(current)
    if err != nil {
        return "", err
    }
    switch commitType {
    case "feat":
        v.Minor++
        v.Patch = 0
    case "fix", "docs", "style", "refactor", "perf", "test", "chore", "build", "ci":
        v.Patch++
    case "revert":
        return current, nil
    default:
        return "", fmt.Errorf("unknown commit type: %s", commitType)
    }
    if breaking {
        v.Major++
        v.Minor = 0
        v.Patch = 0
    }
    return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch), nil
}

// LatestVersionTag finds latest version tag in repository
func LatestVersionTag(repoPath string) (string, error) {
    cmd := exec.Command("git", "tag", "--list", "--sort=-v:refname")
    cmd.Dir = repoPath
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    tags := strings.Split(strings.TrimSpace(string(out)), "\n")
    for _, tag := range tags {
        if tag != "" {
            return tag, nil
        }
    }
    return "0.0.0", nil
}

// UpdateVersionFile writes version to VERSION file
func UpdateVersionFile(repoPath string, version string) error {
    versionFile := filepath.Join(repoPath, "VERSION")
    return os.WriteFile(versionFile, []byte(version+"\n"), 0644)
}
