package mcp

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"sync"
	"testing"
	"time"

	"github.com/tcpsystems/den/internal/grant"
)

type fake struct {
	mu   sync.Mutex
	argv [][]string
}

func (f *fake) Run(ctx context.Context, name string, argv []string) error {
	f.mu.Lock()
	f.argv = append(f.argv, argv)
	f.mu.Unlock()
	return nil
}

func TestSpawnManyProvidersInParallel(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Unix(1_700_000_000, 0)
	c := grant.Contract{
		ID: "g", Seat: "r1", Parent: "orchestrator", Reason: "start two search providers",
		Scope:     []string{"mcp:exa", "mcp:firecrawl"},
		NotBefore: now.Unix(), Expires: now.Add(time.Hour).Unix(),
	}
	signed, err := grant.Sign(priv, c)
	if err != nil {
		t.Fatal(err)
	}
	f := &fake{}
	servers := []Server{
		{ID: "exa", Provider: "exa", Image: "example/exa:1"},
		{ID: "firecrawl", Provider: "firecrawl", Image: "example/firecrawl:1"},
		{ID: "browserbase", Provider: "browserbase", Image: "example/browserbase:1"},
	}
	got := Spawn(context.Background(), f, pub, signed, "r1", now, servers, 3)
	if got[0].Error != "" || got[1].Error != "" {
		t.Fatal(got)
	}
	if got[2].Error == "" {
		t.Fatal("browserbase was not granted")
	}
	if len(f.argv) != 2 {
		t.Fatalf("started %d", len(f.argv))
	}
}

func TestSpawnSkipsMissingImage(t *testing.T) {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	now := time.Unix(1_700_000_000, 0)
	signed, err := grant.Sign(priv, grant.Contract{
		ID: "g", Seat: "r1", Parent: "p", Reason: "index only",
		Scope: []string{"mcp:searxng"}, NotBefore: now.Unix(), Expires: now.Add(time.Minute).Unix(),
	})
	if err != nil {
		t.Fatal(err)
	}
	got := Spawn(context.Background(), &fake{}, pub, signed, "r1", now, []Server{{ID: "searxng", Provider: "local"}}, 1)
	if got[0].Error == "" {
		t.Fatal("want missing image")
	}
}
