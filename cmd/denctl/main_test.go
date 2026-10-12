package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
