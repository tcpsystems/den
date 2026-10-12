// Package grant is the signed permission contract.
//
// A contract names one seat, one parent, a reason, a scope, and a deadline.
// The parent signs the canonical JSON with ed25519. The runtime keeps the
// parent public key. A signature made by any other key is rejected.
// The contract never contains a secret value.
package grant

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// Contract is the bytes that get signed. Field order is the canonical order.
type Contract struct {
	ID        string   `json:"id"`
	Seat      string   `json:"seat"`
	Parent    string   `json:"parent"`
	Reason    string   `json:"reason"`
	Scope     []string `json:"scope"`
	NotBefore int64    `json:"nbf"`
	Expires   int64    `json:"exp"`
}

// Signed is a contract plus the parent's signature. Pub is informational.
// Verify uses the pinned parent key, not a key carried beside the signature.
type Signed struct {
	Contract Contract `json:"contract"`
	Sig      string   `json:"sig"`
}

// Canonical returns the exact bytes the parent signs.
func Canonical(c Contract) ([]byte, error) {
	if c.ID == "" || c.Seat == "" || c.Parent == "" || c.Reason == "" {
		return nil, fmt.Errorf("contract needs id, seat, parent, and reason")
	}
	if len(c.Scope) == 0 {
		return nil, fmt.Errorf("contract needs a scope")
	}
	if c.Expires <= c.NotBefore {
		return nil, fmt.Errorf("contract expiry must be after not-before")
	}
	c.Scope = append([]string(nil), c.Scope...)
	sort.Strings(c.Scope)
	return json.Marshal(c)
}

// Sign returns a signed contract. The private key stays with the parent.
func Sign(parent ed25519.PrivateKey, c Contract) (Signed, error) {
	b, err := Canonical(c)
	if err != nil {
		return Signed{}, err
	}
	return Signed{Contract: c, Sig: base64.StdEncoding.EncodeToString(ed25519.Sign(parent, b))}, nil
}

// Verify checks the pinned parent key, the seat, one scope item, and the clock.
func Verify(parent ed25519.PublicKey, s Signed, seat, scope string, now time.Time) error {
	b, err := Canonical(s.Contract)
	if err != nil {
		return err
	}
	sig, err := base64.StdEncoding.DecodeString(s.Sig)
	if err != nil {
		return fmt.Errorf("signature is not base64")
	}
	if !ed25519.Verify(parent, b, sig) {
		return fmt.Errorf("signature does not match the parent key")
	}
	if s.Contract.Seat != seat {
		return fmt.Errorf("contract seat is %s", s.Contract.Seat)
	}
	if !hasScope(s.Contract.Scope, scope) {
		return fmt.Errorf("scope %s is not granted", scope)
	}
	unix := now.Unix()
	if unix < s.Contract.NotBefore {
		return fmt.Errorf("contract is not active yet")
	}
	if unix >= s.Contract.Expires {
		return fmt.Errorf("contract has expired")
	}
	return nil
}

func hasScope(scope []string, want string) bool {
	for _, s := range scope {
		if s == want {
			return true
		}
	}
	return false
}
