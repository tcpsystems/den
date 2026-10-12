package orch

import "testing"

func TestOwnsPath(t *testing.T) {
	owns := []string{"research/drafts/eve"}
	if !ownsPath(owns, "research/drafts/eve/page.md") {
		t.Fatal("nested path")
	}
	if ownsPath(owns, "research/drafts/other.md") {
		t.Fatal("sibling")
	}
	if ownsPath([]string{"research/drafts/eve"}, "../secrets") {
		t.Fatal("escape")
	}
}

func TestCheckMessage(t *testing.T) {
	ok := []struct{ role, msg string }{
		{RoleResearch, "docs(research): note eve product page"},
		{RoleBuild, "feat(denorch): add wave merge"},
		{RoleBuild, "fix(orch): reject dirty trees"},
		{RoleReview, "test(orch): cover conventional subjects"},
		{"", "chore(git): ignore orch state"},
	}
	for _, c := range ok {
		if err := CheckMessage(c.role, c.msg); err != nil {
			t.Fatalf("%s %q: %v", c.role, c.msg, err)
		}
	}
	bad := []struct{ role, msg string }{
		{RoleResearch, "updated the readme"},
		{RoleResearch, "feat(x): builders only"},
		{RoleReview, "docs(x): reviewers test"},
	}
	for _, c := range bad {
		if err := CheckMessage(c.role, c.msg); err == nil {
			t.Fatalf("%s %q: want error", c.role, c.msg)
		}
	}
}
