package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestStateJSONNamesRepo(t *testing.T) {
	repo := t.TempDir()
	b, err := stateJSON(context.Background(), repo)
	if err != nil {
		t.Fatal(err)
	}
	var body map[string]any
	if err := json.Unmarshal(b, &body); err != nil {
		t.Fatal(err)
	}
	if body["repo"] == "" {
		t.Fatalf("%s", b)
	}
	req := httptest.NewRequest(http.MethodGet, "/api/state", nil)
	rec := httptest.NewRecorder()
	writeState(rec, req, repo)
	if rec.Code != 200 {
		t.Fatal(rec.Code)
	}
}

func TestEventsStreamsState(t *testing.T) {
	repo := t.TempDir()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/events", eventsHandler(repo))
	srv := httptest.NewServer(lanOnly(mux))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/api/events", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := srv.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatal(res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); !strings.Contains(ct, "text/event-stream") {
		t.Fatal(ct)
	}

	type line struct {
		text string
		err  error
	}
	lines := make(chan line, 8)
	go func() {
		sc := bufio.NewScanner(res.Body)
		for sc.Scan() {
			lines <- line{text: sc.Text()}
		}
		lines <- line{err: sc.Err()}
	}()
	var event, data string
	sawRetry := false
	deadline := time.After(2 * time.Second)
	for event == "" || data == "" {
		select {
		case <-deadline:
			t.Fatal("no state event")
		case ln := <-lines:
			if ln.err != nil {
				t.Fatal(ln.err)
			}
			switch {
			case strings.HasPrefix(ln.text, "retry:"):
				sawRetry = true
			case strings.HasPrefix(ln.text, "event:"):
				event = strings.TrimSpace(strings.TrimPrefix(ln.text, "event:"))
			case strings.HasPrefix(ln.text, "data:"):
				data = strings.TrimSpace(strings.TrimPrefix(ln.text, "data:"))
			}
		}
	}
	if !sawRetry {
		t.Fatal("missing retry")
	}
	if event != "state" {
		t.Fatal(event)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(data), &body); err != nil {
		t.Fatal(err)
	}
	if body["repo"] == "" || body["repo"] == nil {
		t.Fatalf("%s", data)
	}
}

func newSteerMux() *http.ServeMux {
	store := newNotesStore()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/steer", steerHandler(store))
	mux.HandleFunc("/api/notes", notesHandler(store))
	return mux
}

func getNotes(t *testing.T, mux *http.ServeMux) []note {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/api/notes", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("notes status %d", rec.Code)
	}
	var notes []note
	if err := json.Unmarshal(rec.Body.Bytes(), &notes); err != nil {
		t.Fatal(err)
	}
	return notes
}

func postSteer(t *testing.T, mux *http.ServeMux, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/api/steer", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	return rec
}

func TestNotesEmptyAtStart(t *testing.T) {
	mux := newSteerMux()
	if notes := getNotes(t, mux); len(notes) != 0 {
		t.Fatalf("want empty, got %v", notes)
	}
}

func TestSteerNoteThenList(t *testing.T) {
	mux := newSteerMux()
	if rec := postSteer(t, mux, `{"action":"note","text":"first"}`); rec.Code != http.StatusOK {
		t.Fatalf("first post status %d", rec.Code)
	}
	if rec := postSteer(t, mux, `{"action":"note","text":"second"}`); rec.Code != http.StatusOK {
		t.Fatalf("second post status %d", rec.Code)
	}
	notes := getNotes(t, mux)
	if len(notes) != 2 {
		t.Fatalf("want 2 notes, got %d", len(notes))
	}
	if notes[0].Text != "second" || notes[1].Text != "first" {
		t.Fatalf("want newest first, got %v", notes)
	}
}

func TestSteerBadActionIs400(t *testing.T) {
	mux := newSteerMux()
	for _, body := range []string{
		`{"action":"evict","text":"x"}`,
		`{"action":"","text":"x"}`,
		`not json`,
	} {
		if rec := postSteer(t, mux, body); rec.Code != http.StatusBadRequest {
			t.Fatalf("body %q status %d", body, rec.Code)
		}
	}
	if notes := getNotes(t, mux); len(notes) != 0 {
		t.Fatalf("bad actions must not add notes, got %v", notes)
	}
}
