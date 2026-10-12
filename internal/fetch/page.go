// Package fetch reads one public page through Firecrawl.
//
// Firecrawl is the page reader already configured for this machine.
// The client sends the URL as itself. It does not change its identity,
// and it does not open a private host.
package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const Endpoint = "https://api.firecrawl.dev/v2/scrape"

// MaxMarkdown is the most markdown one call keeps.
const MaxMarkdown = 256 << 10

// ErrPrivate is returned when the URL names a host this tool must not open.
var ErrPrivate = errors.New("fetch: host is not public")

// Page is the markdown Firecrawl returned for one URL.
type Page struct {
	URL      string
	Markdown string
}

type scrapeRequest struct {
	URL     string   `json:"url"`
	Formats []string `json:"formats"`
}

type scrapeResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Markdown string `json:"markdown"`
	} `json:"data"`
	Error string `json:"error"`
}

// Public reports whether rawURL is an http or https URL on a public host.
// It refuses loopback, link-local, private, and unspecified addresses,
// and the names localhost and *.local.
func Public(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("fetch: scheme %q", u.Scheme)
	}
	host := u.Hostname()
	if host == "" {
		return errors.New("fetch: missing host")
	}
	low := strings.ToLower(host)
	if low == "localhost" || strings.HasSuffix(low, ".local") || strings.HasSuffix(low, ".internal") {
		return ErrPrivate
	}
	if ip := net.ParseIP(host); ip != nil && !publicIP(ip) {
		return ErrPrivate
	}
	return nil
}

func publicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsUnspecified() || ip.IsMulticast() {
		return false
	}
	return true
}

// Scrape posts rawURL to Firecrawl and returns markdown.
// key is the Firecrawl bearer token. It is not written into Page or into an error.
func Scrape(ctx context.Context, client *http.Client, key, rawURL string) (Page, error) {
	if err := Public(rawURL); err != nil {
		return Page{}, err
	}
	if strings.TrimSpace(key) == "" {
		return Page{}, errors.New("fetch: missing key")
	}
	if client == nil {
		client = &http.Client{Timeout: 60 * time.Second}
	}
	body, err := json.Marshal(scrapeRequest{URL: rawURL, Formats: []string{"markdown"}})
	if err != nil {
		return Page{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, Endpoint, bytes.NewReader(body))
	if err != nil {
		return Page{}, err
	}
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return Page{}, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return Page{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Page{}, fmt.Errorf("fetch: firecrawl status %d", resp.StatusCode)
	}
	var parsed scrapeResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return Page{}, err
	}
	if !parsed.Success {
		msg := parsed.Error
		if msg == "" {
			msg = "unsuccessful"
		}
		return Page{}, fmt.Errorf("fetch: firecrawl %s", msg)
	}
	text := parsed.Data.Markdown
	if len(text) > MaxMarkdown {
		text = text[:MaxMarkdown]
	}
	return Page{URL: rawURL, Markdown: text}, nil
}
