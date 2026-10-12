// Package initd starts the smallest harness and records what the operator turns on.
package initd

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Part is one piece the operator can turn on after init.
type Part struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Note   string `json:"note"`
}

// Board is the init record.
type Board struct {
	Repo  string `json:"repo"`
	Sock  string `json:"sock"`
	Parts []Part `json:"parts"`
}

// Default is the board init writes. The controller and the TUI are on.
// Everything else stays off until the operator enables it.
func Default(repo, sock string) Board {
	return Board{
		Repo: repo,
		Sock: sock,
		Parts: []Part{
			{ID: "controller", Title: "Controller", Status: "on", Note: "unix socket"},
			{ID: "tui", Title: "TUI", Status: "on", Note: "this screen"},
			{ID: "web", Title: "Web", Status: "off", Note: "loopback page"},
			{ID: "pages", Title: "Page reader", Status: "off", Note: "denfetch"},
			{ID: "files", Title: "File reader", Status: "off", Note: "denread"},
			{ID: "seats", Title: "Seats", Status: "off", Note: "s opens four denorch worktrees"},
		},
	}
}

func path(repo string) string {
	return filepath.Join(repo, ".den", "init.json")
}

// Load reads the board. A missing file returns Default.
func Load(repo, sock string) (Board, error) {
	b, err := os.ReadFile(path(repo))
	if os.IsNotExist(err) {
		return Default(repo, sock), nil
	}
	if err != nil {
		return Board{}, err
	}
	var board Board
	if err := json.Unmarshal(b, &board); err != nil {
		return Board{}, err
	}
	if board.Sock == "" {
		board.Sock = sock
	}
	if board.Repo == "" {
		board.Repo = repo
	}
	return board, nil
}

// Save writes the board under repo/.den. The directory is not a git path.
func Save(board Board) error {
	dir := filepath.Join(board.Repo, ".den")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(board, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "init.json"), append(b, '\n'), 0o600)
}

// Present reports whether this build has a process for id.
func Present(id string) bool {
	switch id {
	case "controller", "tui", "seats":
		return true
	default:
		return false
	}
}

// Enable marks id on when Present is true. Otherwise the part stays off.
func Enable(board *Board, id string) bool {
	if !Present(id) {
		return false
	}
	for i := range board.Parts {
		if board.Parts[i].ID == id {
			board.Parts[i].Status = "on"
			return true
		}
	}
	board.Parts = append(board.Parts, Part{ID: id, Title: id, Status: "on"})
	return true
}
