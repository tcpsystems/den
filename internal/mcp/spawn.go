// Package mcp starts many provider containers.
//
// The catalog may list any number of providers. None is required.
// Spawn runs them in parallel. A container starts only when the signed
// grant includes mcp:<id> for the seat. Secret values are not arguments.
package mcp

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"sync"
	"time"

	"github.com/tcpsystems/den/internal/grant"
)

// Server is one MCP provider. Image empty means the operator has not chosen a container.
type Server struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Image    string `json:"image,omitempty"`
	Note     string `json:"note,omitempty"`
}

// Runner starts one container. Tests substitute a fake.
type Runner interface {
	Run(ctx context.Context, name string, argv []string) error
}

// Result is one provider's start attempt.
type Result struct {
	ID    string `json:"id"`
	Error string `json:"error,omitempty"`
}

// Spawn starts every allowed server at once, bounded by parallel.
func Spawn(ctx context.Context, run Runner, parent ed25519.PublicKey, signed grant.Signed, seat string, now time.Time, servers []Server, parallel int) []Result {
	if parallel < 1 {
		parallel = 1
	}
	out := make([]Result, len(servers))
	sem := make(chan struct{}, parallel)
	var wg sync.WaitGroup
	for i, srv := range servers {
		wg.Add(1)
		go func(i int, srv Server) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			out[i] = startOne(ctx, run, parent, signed, seat, now, srv)
		}(i, srv)
	}
	wg.Wait()
	return out
}

func startOne(ctx context.Context, run Runner, parent ed25519.PublicKey, signed grant.Signed, seat string, now time.Time, srv Server) Result {
	res := Result{ID: srv.ID}
	if err := grant.Verify(parent, signed, seat, "mcp:"+srv.ID, now); err != nil {
		res.Error = err.Error()
		return res
	}
	if srv.Image == "" {
		res.Error = "no image set for " + srv.ID
		return res
	}
	argv := []string{
		"docker", "run", "-d", "--name", "den-mcp-" + srv.ID,
		"--read-only", "--cap-drop", "ALL", "--network", "den-mcp",
		srv.Image,
	}
	if err := run.Run(ctx, srv.ID, argv); err != nil {
		res.Error = err.Error()
		return res
	}
	return res
}

// ArgsFor exposes the docker argv for tests.
func ArgsFor(id, image string) []string {
	return []string{"docker", "run", "-d", "--name", "den-mcp-" + id, "--read-only", "--cap-drop", "ALL", "--network", "den-mcp", image}
}

// Deadline is a convenience for callers that pass time.Now.
func Deadline(seconds int) time.Time {
	return time.Now().Add(time.Duration(seconds) * time.Second)
}

// FormatError joins result errors. An empty string means every start succeeded.
func FormatError(results []Result) error {
	var parts []string
	for _, r := range results {
		if r.Error != "" {
			parts = append(parts, r.ID+": "+r.Error)
		}
	}
	if len(parts) == 0 {
		return nil
	}
	return fmt.Errorf("%s", join(parts))
}

func join(parts []string) string {
	out := parts[0]
	for _, p := range parts[1:] {
		out += "; " + p
	}
	return out
}
