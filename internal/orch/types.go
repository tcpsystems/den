// Package orch is the local git orchestrator for a swarm on one wave branch.
//
// Workers write only inside the paths they own, in their own worktrees.
// The orchestrator checks conventional commits, runs the named checks, and
// merges a clean seat onto the wave branch. It does not push.
package orch

import "time"

// Roles the swarm seats use. The prefix is the conventional-commit type
// the seat is expected to use.
const (
	RoleResearch = "researcher"
	RolePlan     = "architect"
	RoleBuild    = "builder"
	RoleReview   = "reviewer"
)

// Seat is one worker assignment on a wave branch.
type Seat struct {
	ID     string   `json:"id"`
	Role   string   `json:"role"`
	Brief  string   `json:"brief"`
	Owns   []string `json:"owns"`
	Status string   `json:"status"`
	Head   string   `json:"head,omitempty"`
	Note   string   `json:"note,omitempty"`
}

// Wave is the shared branch and the seats writing to it.
type Wave struct {
	ID        string    `json:"id"`
	Branch    string    `json:"branch"`
	Base      string    `json:"base"`
	Repo      string    `json:"repo"`
	Worktrees string    `json:"worktrees"`
	Seats     []Seat    `json:"seats"`
	Created   time.Time `json:"created"`
}

// Check is one command the orchestrator runs before a merge.
type Check struct {
	Name string   `json:"name"`
	Argv []string `json:"argv"`
}

// DefaultChecks are the tests this repo can run without extra services.
func DefaultChecks() []Check {
	return []Check{{Name: "go-test", Argv: []string{"go", "test", "./..."}}}
}
