// Package vault holds secret values inside the runtime.
//
// A prompt, a log, and a catalog see names only. Reveal returns one value
// when a signed grant names that secret, that seat, and has not expired.
package vault

import (
	"crypto/ed25519"
	"fmt"
	"strings"
	"time"

	"github.com/tcpsystems/den/internal/grant"
)

// Store is the runtime's secret map. Implementations must not log values.
type Store interface {
	Get(name string) (string, bool)
}

// Map is an in-memory store for tests and for a process that already loaded values.
type Map map[string]string

func (m Map) Get(name string) (string, bool) {
	v, ok := m[name]
	return v, ok
}

// Redact replaces values of lines that look like secrets. Names stay.
func Redact(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		eq := strings.IndexByte(line, '=')
		if eq <= 0 {
			continue
		}
		name := strings.TrimSpace(line[:eq])
		if secretName(name) {
			lines[i] = name + "=[redacted]"
		}
	}
	return strings.Join(lines, "\n")
}

func secretName(name string) bool {
	n := strings.ToLower(name)
	for _, p := range []string{"key", "token", "secret", "password", "passwd", "credential"} {
		if strings.Contains(n, p) {
			return true
		}
	}
	return false
}

// Reveal returns one secret when the grant allows secret:<name> for this seat.
func Reveal(parent ed25519.PublicKey, store Store, signed grant.Signed, seat, name string, now time.Time) (string, error) {
	if err := grant.Verify(parent, signed, seat, "secret:"+name, now); err != nil {
		return "", err
	}
	v, ok := store.Get(name)
	if !ok || v == "" {
		return "", fmt.Errorf("secret %s is not in the runtime store", name)
	}
	return v, nil
}
