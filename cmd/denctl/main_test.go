package main

import (
	"bufio"
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
