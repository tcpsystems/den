// dencat prints the runtime catalog. It does not start agents.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/tcpsystems/den/internal/catalog"
)

func main() {
	repo := flag.String("repo", ".", "repository root")
	kind := flag.String("kind", "", "filter to one kind")
	terms := flag.Bool("terms", false, "print the vocabulary")
	flag.Parse()
	if *terms {
		fmt.Print(catalog.Terms)
		return
	}
	idx := catalog.Builtin()
	if err := catalog.LoadPlugins(&idx, *repo+"/plugins"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *kind != "" {
		idx.Entries = idx.ByKind(catalog.Kind(*kind))
	}
	b, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
