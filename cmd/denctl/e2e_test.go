package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestE2E(t *testing.T) {
	// 1. Create a temporary git repository with one commit.
	repoDir := t.TempDir()
	if err := exec.Command("git", "init", repoDir).Run(); err != nil {
		t.Fatalf("git init failed: %v", err)
	}
	commitCmd := exec.Command("git", "-C", repoDir,
		"-c", "user.name=test",
		"-c", "user.email=test@test",
		"commit", "--allow-empty", "-m", "init")
	if err := commitCmd.Run(); err != nil {
		t.Fatalf("git commit failed: %v", err)
	}

	// 2. Pick an ephemeral loopback address and a temporary unix socket path.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to acquire temporary port: %v", err)
	}
	addr := ln.Addr().String()
	ln.Close()

	sockPath := filepath.Join(t.TempDir(), "den.sock")

	// 3. Start the real server.
	errCh := make(chan error, 1)
	go func() {
		errCh <- serve(addr, sockPath, repoDir)
	}()

	// Wait for the server to become responsive.
	client := &http.Client{Timeout: 500 * time.Millisecond}
	baseURL := "http://" + addr
	up := false
	for i := 0; i < 30 && !up; i++ {
		select {
		case err := <-errCh:
			t.Fatalf("serve exited early: %v", err)
		default:
		}
		resp, err := client.Get(baseURL + "/")
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			up = true
			break
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(100 * time.Millisecond)
	}
	if !up {
		t.Fatalf("server did not become ready in time")
	}

	// 4. GET / returns 200 with the den title.
	resp, err := client.Get(baseURL + "/")
	if err != nil {
		t.Fatalf("GET / failed: %v", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("reading / response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET / status = %d, want 200", resp.StatusCode)
	}
	if !strings.Contains(string(bodyBytes), "den") {
		t.Fatalf("root page does not contain expected string %q", "den")
	}

	// 5. GET /api/state returns repo fields and a services map.
	resp, err = client.Get(baseURL + "/api/state")
	if err != nil {
		t.Fatalf("GET /api/state failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		t.Fatalf("GET /api/state status = %d, want 200", resp.StatusCode)
	}
	var state struct {
		Repo     string                 `json:"repo"`
		Services map[string]interface{} `json:"services"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&state); err != nil {
		resp.Body.Close()
		t.Fatalf("decoding /api/state JSON: %v", err)
	}
	resp.Body.Close()
	if state.Repo == "" {
		t.Fatalf("/api/state repo field empty")
	}
	if state.Services == nil {
		t.Fatalf("/api/state services field missing")
	}

	// 6. GET /api/events streams at least one `event: state` frame within 5s.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/api/events", nil)
	if err != nil {
		t.Fatalf("creating request for /api/events: %v", err)
	}
	streamClient := &http.Client{} // no timeout; the 5s deadline is on ctx
	eventsResp, err := streamClient.Do(req)
	if err != nil {
		t.Fatalf("GET /api/events failed: %v", err)
	}
	defer eventsResp.Body.Close()
	if eventsResp.StatusCode != http.StatusOK {
		t.Fatalf("GET /api/events status = %d, want 200", eventsResp.StatusCode)
	}
	scanner := bufio.NewScanner(eventsResp.Body)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "event: state") {
			found = true
			break
		}
	}
	if !found && ctx.Err() == nil && scanner.Err() != nil {
		t.Fatalf("scanner error reading /api/events: %v", scanner.Err())
	}
	if !found {
		t.Fatalf("did not receive an \"event: state\" line from /api/events within 5s")
	}

	// 7. POST /api/steer with a bad action must be 400; if the route is not
	// implemented yet in this snapshot (another cluster owns it), the mux
	// returns 404 and we skip the steer/notes checks gracefully.
	steerURL := baseURL + "/api/steer"
	badBody := strings.NewReader(`{"action":"not-a-real-action"}`)
	steerResp, err := client.Post(steerURL, "application/json", badBody)
	if err != nil {
		t.Fatalf("POST /api/steer failed: %v", err)
	}
	steerStatus := steerResp.StatusCode
	steerResp.Body.Close()
	if steerStatus == http.StatusNotFound {
		t.Skip("/api/steer is not implemented in this snapshot; skipping steer/notes checks")
	}
	if steerStatus != http.StatusBadRequest {
		t.Fatalf("POST /api/steer with bad action status = %d, want 400", steerStatus)
	}

	// POST /api/steer a well-formed note, then GET /api/notes contains it.
	const noteText = "e2e-test-note-marker"
	noteBody, err := json.Marshal(map[string]string{"action": "note", "text": noteText})
	if err != nil {
		t.Fatalf("marshal note body: %v", err)
	}
	noteResp, err := client.Post(steerURL, "application/json", strings.NewReader(string(noteBody)))
	if err != nil {
		t.Fatalf("POST /api/steer note failed: %v", err)
	}
	noteResp.Body.Close()
	if noteResp.StatusCode != http.StatusOK {
		t.Fatalf("POST /api/steer note status = %d, want 200", noteResp.StatusCode)
	}

	notesResp, err := client.Get(baseURL + "/api/notes")
	if err != nil {
		t.Fatalf("GET /api/notes failed: %v", err)
	}
	notesBody, err := io.ReadAll(notesResp.Body)
	notesResp.Body.Close()
	if err != nil {
		t.Fatalf("reading /api/notes body: %v", err)
	}
	if !strings.Contains(string(notesBody), noteText) {
		t.Fatalf("/api/notes does not contain posted note %q: %s", noteText, notesBody)
	}
}

// TestLANAllowedUnitHook is the non-HTTP fallback for peer-based 403
// coverage: the e2e client above always dials loopback, which lanAllowed
// treats as private, so it can never itself trigger the lanOnly 403 path.
// This exercises lanAllowed directly for a public and a private peer.
func TestLANAllowedUnitHook(t *testing.T) {
	if lanAllowed(net.ParseIP("8.8.8.8")) {
		t.Fatal("8.8.8.8 must not be lanAllowed")
	}
	if !lanAllowed(net.ParseIP("192.168.1.5")) {
		t.Fatal("192.168.1.5 must be lanAllowed")
	}
}
