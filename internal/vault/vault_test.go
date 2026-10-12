package vault

import (
	"crypto/ed25519"
	"crypto/rand"
	"strings"
	"testing"
	"time"

	"github.com/tcpsystems/den/internal/grant"
)

func TestRedactKeepsNames(t *testing.T) {
	in := "OMNIROUTE_API_KEY=super-secret\nPLAIN=visible\n"
	out := Redact(in)
	if strings.Contains(out, "super-secret") {
		t.Fatal(out)
	}
	if !strings.Contains(out, "OMNIROUTE_API_KEY=[redacted]") || !strings.Contains(out, "PLAIN=visible") {
		t.Fatal(out)
	}
}

func TestRevealRequiresGrant(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Unix(1_700_000_000, 0)
	store := Map{"SEARCH_TOKEN": "value-not-for-git"}
	c := grant.Contract{
		ID: "g", Seat: "r1", Parent: "orchestrator", Reason: "one search call",
		Scope: []string{"secret:SEARCH_TOKEN"}, NotBefore: now.Unix(), Expires: now.Add(time.Minute).Unix(),
	}
	signed, err := grant.Sign(priv, c)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Reveal(pub, store, signed, "r1", "SEARCH_TOKEN", now)
	if err != nil || got != "value-not-for-git" {
		t.Fatalf("got %q %v", got, err)
	}
	if _, err := Reveal(pub, store, signed, "r1", "OTHER", now); err == nil {
		t.Fatal("other secret")
	}
}
