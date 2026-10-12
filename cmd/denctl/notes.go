package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// note is one human steering note. Never persisted to disk.
type note struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

// notesStore holds steering notes in memory, newest first, capped at 64.
type notesStore struct {
	mu    sync.Mutex
	notes []note
}

func newNotesStore() *notesStore {
	return &notesStore{notes: make([]note, 0)}
}

func (s *notesStore) add(n note) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.notes = append([]note{n}, s.notes...)
	if len(s.notes) > 64 {
		s.notes = s.notes[:64]
	}
}

func (s *notesStore) getAll() []note {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp := make([]note, len(s.notes))
	copy(cp, s.notes)
	return cp
}

// steerHandler accepts {"action":"note","text":"..."}; any other action is 400.
func steerHandler(s *notesStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var body struct {
			Action string `json:"action"`
			Text   string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if body.Action != "note" || body.Text == "" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		n := note{Text: body.Text, Timestamp: time.Now()}
		s.add(n)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(n)
	}
}

func notesHandler(s *notesStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		notes := s.getAll()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(notes)
	}
}
