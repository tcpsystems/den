package inject

import (
	"testing"

	"github.com/tcpsystems/den/internal/catalog"
	"github.com/tcpsystems/den/internal/grant"
)

func TestOneReturnsGrantedSkill(t *testing.T) {
	idx := catalog.Builtin()
	signed := grant.Signed{Contract: grant.Contract{
		Seat: "den-1", Scope: []string{"skill:use-git"},
	}}
	path, ok := One(idx, signed, "den-1", "commit this file", true)
	if !ok || path != "skills/use-git/SKILL.md" {
		t.Fatalf("%s %v", path, ok)
	}
}

func TestOneDefaultsWhenUngranted(t *testing.T) {
	idx := catalog.Builtin()
	signed := grant.Signed{Contract: grant.Contract{
		Seat: "den-1", Scope: []string{"mcp:exa"},
	}}
	path, ok := One(idx, signed, "den-1", "search", true)
	if ok || path != DefaultPath {
		t.Fatalf("%s %v", path, ok)
	}
}

func TestOneRefusesWrongSeat(t *testing.T) {
	idx := catalog.Builtin()
	signed := grant.Signed{Contract: grant.Contract{
		Seat: "other", Scope: []string{"skill:use-git"},
	}}
	path, ok := One(idx, signed, "den-1", "commit", true)
	if ok || path != DefaultPath {
		t.Fatalf("%s %v", path, ok)
	}
}

func TestOneRefusesEmptyDescription(t *testing.T) {
	idx := catalog.Builtin()
	signed := grant.Signed{Contract: grant.Contract{
		Seat: "den-1", Scope: []string{"skill:use-git"},
	}}
	if _, ok := One(idx, signed, "den-1", "", true); ok {
		t.Fatal("empty description")
	}
}
