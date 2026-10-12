// Package inject returns one skill path.
//
// It does not rank, embed, or call a model. The caller already chose the
// description. A grant that names skill:<id> returns that catalog path.
// No match returns the default path. The catalog is not copied into a prompt.
package inject

import (
	"github.com/tcpsystems/den/internal/catalog"
	"github.com/tcpsystems/den/internal/grant"
)

const DefaultPath = "skills/use-harness/SKILL.md"

// One returns the path for the first skill scope the grant allows and the
// catalog lists. An unsigned or expired grant returns DefaultPath and false.
func One(idx catalog.Index, signed grant.Signed, seat, description string, parentOK bool) (string, bool) {
	if !parentOK || description == "" {
		return DefaultPath, false
	}
	for _, scope := range signed.Contract.Scope {
		id, ok := skillID(scope)
		if !ok {
			continue
		}
		if !contains(signed.Contract.Scope, "skill:"+id) {
			continue
		}
		if signed.Contract.Seat != seat {
			continue
		}
		for _, e := range idx.Entries {
			if e.Kind == catalog.KindSkill && e.ID == id && e.Path != "" {
				return e.Path, true
			}
		}
	}
	return DefaultPath, false
}

func skillID(scope string) (string, bool) {
	const prefix = "skill:"
	if len(scope) <= len(prefix) || scope[:len(prefix)] != prefix {
		return "", false
	}
	return scope[len(prefix):], true
}

func contains(ss []string, want string) bool {
	for _, s := range ss {
		if s == want {
			return true
		}
	}
	return false
}
