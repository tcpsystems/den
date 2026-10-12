package grant

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"
	"time"
)

func TestSignVerifyScopeAndExpiry(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Unix(1_700_000_000, 0)
	c := Contract{
		ID: "g1", Seat: "r2", Parent: "orchestrator", Reason: "read the search key for this step",
		Scope: []string{"secret:SEARXNG_NONE", "mcp:exa"}, NotBefore: now.Unix(), Expires: now.Add(time.Hour).Unix(),
	}
	s, err := Sign(priv, c)
	if err != nil {
		t.Fatal(err)
	}
	if err := Verify(pub, s, "r2", "mcp:exa", now.Add(time.Minute)); err != nil {
		t.Fatal(err)
	}
	if err := Verify(pub, s, "r9", "mcp:exa", now); err == nil {
		t.Fatal("wrong seat")
	}
	if err := Verify(pub, s, "r2", "mcp:other", now); err == nil {
		t.Fatal("wrong scope")
	}
	if err := Verify(pub, s, "r2", "mcp:exa", now.Add(2*time.Hour)); err == nil {
		t.Fatal("expired")
	}
	_, other, _ := ed25519.GenerateKey(rand.Reader)
	forged, err := Sign(other, c)
	if err != nil {
		t.Fatal(err)
	}
	if err := Verify(pub, forged, "r2", "mcp:exa", now); err == nil {
		t.Fatal("other key accepted")
	}
}

func TestContractRejectsEmptyReason(t *testing.T) {
	_, priv, _ := ed25519.GenerateKey(rand.Reader)
	_, err := Sign(priv, Contract{ID: "g", Seat: "s", Parent: "p", Scope: []string{"mcp:x"}, NotBefore: 1, Expires: 2})
	if err == nil {
		t.Fatal("want reason")
	}
}
