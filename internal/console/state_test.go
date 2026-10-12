package console

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSlotsAreOneList(t *testing.T) {
	slots := Slots()
	if len(slots) < 10 {
		t.Fatalf("slots %d", len(slots))
	}
	seen := map[string]bool{}
	for _, slot := range slots {
		if seen[slot.ID] {
			t.Fatalf("duplicate %s", slot.ID)
		}
		seen[slot.ID] = true
		if slot.Status != "this" && slot.Status != "later" {
			t.Fatalf("%s status %s", slot.ID, slot.Status)
		}
	}
	if !seen["go"] || !seen["svelte"] || !seen["sse"] || !seen["dry"] {
		t.Fatal("missing this-build slots")
	}
	for _, id := range []string{"react", "shadcn", "cloudflare", "hetzner", "scaleblade", "sentry", "d1", "privacy", "devevalflow"} {
		if !seen[id] {
			t.Fatalf("missing later slot %s", id)
		}
	}
}

func TestProjectsStayOutOfSecrets(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(dir+"/Projects/harness-plan", 0o755); err != nil {
		t.Fatal(err)
	}
	rows := Projects(dir)
	if len(rows) < 5 {
		t.Fatalf("projects %d", len(rows))
	}
	var den Project
	for _, row := range rows {
		if strings.Contains(strings.ToLower(row.Role), "token") || strings.Contains(strings.ToLower(row.Name), "secret") {
			t.Fatalf("secret in project row %+v", row)
		}
		if row.Name == "den" {
			den = row
		}
	}
	if den.Status != "present" || den.Org != "tcpsystems" {
		t.Fatalf("den %+v", den)
	}
}

func TestAtoi(t *testing.T) {
	if atoi("12") != 12 || atoi("x") != 0 {
		t.Fatal("atoi")
	}
}

// TestProbeDownDoesNotStall covers a service that accepts the connection
// and then hangs: probe must report "down" and return at roughly
// probeTimeout, never near the handler's full sleep.
func TestProbeDownDoesNotStall(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	start := time.Now()
	status := probe(context.Background(), srv.URL)
	elapsed := time.Since(start)

	if status != "down" {
		t.Fatalf("status = %q, want down", status)
	}
	if elapsed > probeTimeout+500*time.Millisecond {
		t.Fatalf("probe took %v, want roughly %v", elapsed, probeTimeout)
	}
}
