// denmcp indexes MCP providers and starts their containers in parallel.
//
// No provider is mandatory. SearXNG stays a local search engine and is not
// replaced here. Images are whatever the catalog names. Secret values are
// not accepted on the command line.
package main

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/tcpsystems/den/internal/grant"
	"github.com/tcpsystems/den/internal/mcp"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "list":
		if err := cmdList(os.Args[2:]); err != nil {
			fail(err)
		}
	case "spawn":
		if err := cmdSpawn(os.Args[2:]); err != nil {
			fail(err)
		}
	default:
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "denmcp list|spawn")
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func cmdList(args []string) error {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	file := fs.String("catalog", "mcp/catalog.json", "provider catalog")
	if err := fs.Parse(args); err != nil {
		return err
	}
	servers, err := load(*file)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(servers, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func cmdSpawn(args []string) error {
	fs := flag.NewFlagSet("spawn", flag.ContinueOnError)
	file := fs.String("catalog", "mcp/catalog.json", "provider catalog")
	ids := fs.String("ids", "", "comma-separated provider ids")
	grantPath := fs.String("grant", "", "signed grant json")
	keyPath := fs.String("parent-key", "", "parent ed25519 public key, raw 32 bytes")
	seat := fs.String("seat", "", "seat id")
	parallel := fs.Int("parallel", 4, "containers started at once")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *ids == "" || *grantPath == "" || *keyPath == "" || *seat == "" {
		return fmt.Errorf("spawn needs -ids -grant -parent-key -seat")
	}
	servers, err := load(*file)
	if err != nil {
		return err
	}
	want := map[string]bool{}
	for _, id := range strings.Split(*ids, ",") {
		want[strings.TrimSpace(id)] = true
	}
	var chosen []mcp.Server
	for _, s := range servers {
		if want[s.ID] {
			chosen = append(chosen, s)
		}
	}
	pub, err := os.ReadFile(*keyPath)
	if err != nil {
		return err
	}
	if len(pub) != ed25519.PublicKeySize {
		return fmt.Errorf("parent key must be %d bytes", ed25519.PublicKeySize)
	}
	gb, err := os.ReadFile(*grantPath)
	if err != nil {
		return err
	}
	var signed grant.Signed
	if err := json.Unmarshal(gb, &signed); err != nil {
		return err
	}
	results := mcp.Spawn(context.Background(), dockerRunner{}, ed25519.PublicKey(pub), signed, *seat, time.Now(), chosen, *parallel)
	b, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return mcp.FormatError(results)
}

func load(path string) ([]mcp.Server, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var servers []mcp.Server
	if err := json.Unmarshal(b, &servers); err != nil {
		return nil, err
	}
	return servers, nil
}

type dockerRunner struct{}

func (dockerRunner) Run(ctx context.Context, name string, argv []string) error {
	cmd := exec.CommandContext(ctx, argv[0], argv[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			msg = err.Error()
		}
		return fmt.Errorf("%s", msg)
	}
	return nil
}
