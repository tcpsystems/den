package orch

import (
	"fmt"
	"regexp"
	"strings"
)

// conventional is type(scope): subject, with an optional breaking bang.
var conventional = regexp.MustCompile(`^(feat|fix|chore|refactor|docs|test|perf|style)(\([^)]+\))?\!?: .+`)

// roleType is the commit type a role should use. Empty means any allowed type.
func roleType(role string) string {
	switch role {
	case RoleResearch, RolePlan:
		return "docs"
	case RoleBuild:
		return "feat"
	case RoleReview:
		return "test"
	default:
		return ""
	}
}

// CheckMessage reports whether subject follows conventional commits.
// When role is set, the type must match that role, except fix which any role may use.
func CheckMessage(role, subject string) error {
	subject = strings.TrimSpace(subject)
	if !conventional.MatchString(subject) {
		return fmt.Errorf("commit subject %q is not conventional (feat|fix|chore|refactor|docs|test|perf|style)", subject)
	}
	want := roleType(role)
	if want == "" {
		return nil
	}
	got := strings.SplitN(subject, "(", 2)[0]
	got = strings.TrimSuffix(got, "!")
	if got == "fix" || got == want || (role == RoleBuild && (got == "fix" || got == "refactor" || got == "feat")) {
		return nil
	}
	if role == RoleResearch && (got == "docs" || got == "chore") {
		return nil
	}
	return fmt.Errorf("role %s uses %s(, got %s", role, want, got)
}
