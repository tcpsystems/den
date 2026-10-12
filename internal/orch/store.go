package orch

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func statePath(repo, waveID string) string {
	return filepath.Join(repo, ".den", "orch", waveID+".json")
}

// Save writes the wave ledger under repo/.den/orch. That directory is gitignored.
func Save(w Wave) error {
	if w.Created.IsZero() {
		w.Created = time.Now().UTC()
	}
	dir := filepath.Dir(statePath(w.Repo, w.ID))
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(statePath(w.Repo, w.ID), append(b, '\n'), 0o600)
}

// Load reads a wave ledger.
func Load(repo, waveID string) (Wave, error) {
	b, err := os.ReadFile(statePath(repo, waveID))
	if err != nil {
		return Wave{}, err
	}
	var w Wave
	if err := json.Unmarshal(b, &w); err != nil {
		return Wave{}, err
	}
	return w, nil
}

func (w *Wave) seat(id string) (*Seat, error) {
	for i := range w.Seats {
		if w.Seats[i].ID == id {
			return &w.Seats[i], nil
		}
	}
	return nil, fmt.Errorf("seat %s is not on wave %s", id, w.ID)
}
