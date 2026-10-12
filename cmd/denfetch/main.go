// denfetch searches and reads public pages.
//
// Search goes to OmniRoute at https://ai.tcp.systems/v1/search.
// A page read goes to Firecrawl. Both clients send their own names.
// OMNIROUTE_API_KEY and FIRECRAWL_API_KEY are read from the environment
// and never printed. A private host is dropped before a page is opened.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tcpsystems/den/internal/fetch"
)

func refusePrivate(raw string) error {
	return fetch.Public(raw)
}

func main() {
	readN := flag.Int("read", 0, "after a search, read this many public pages (max 4)")
	flag.Parse()
	arg := strings.TrimSpace(flag.Arg(0))
	if arg == "" {
		fmt.Fprintln(os.Stderr, "usage: denfetch https://example.com/page")
		fmt.Fprintln(os.Stderr, "       denfetch -read 3 'agent harness'")
		os.Exit(2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	if strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://") {
		page, err := fetch.Scrape(ctx, nil, os.Getenv("FIRECRAWL_API_KEY"), arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("# %s\n\n%s\n", page.URL, page.Markdown)
		return
	}
	n := *readN
	if n == 0 {
		n = 3
	}
	brief, err := fetch.Research(ctx, nil, os.Getenv("OMNIROUTE_API_KEY"), os.Getenv("FIRECRAWL_API_KEY"), arg, n)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("# %s\n\n", brief.Query)
	for i, hit := range brief.Hits {
		fmt.Printf("%d. %s\n   %s\n   %s\n\n", i+1, hit.Title, hit.URL, oneLine(hit.Snippet))
	}
	for _, page := range brief.Pages {
		fmt.Printf("\n## %s\n\n%s\n", page.URL, page.Markdown)
	}
}

func oneLine(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	if len(s) > 240 {
		return s[:240]
	}
	return s
}
