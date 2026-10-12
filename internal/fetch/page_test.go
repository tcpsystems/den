package fetch

import (
	"context"
	"testing"
	"time"
)

func TestPublic(t *testing.T) {
	ok := []string{
		"https://www.anthropic.com/engineering/effective-harnesses-for-long-running-agents",
		"https://github.com/ryanalberts/best-of-agent-harnesses",
		"http://example.com/docs",
	}
	for _, u := range ok {
		if err := Public(u); err != nil {
			t.Errorf("Public(%q) = %v", u, err)
		}
	}
	bad := []string{
		"https://127.0.0.1/admin",
		"https://10.0.0.5/",
		"https://192.168.0.166/",
		"https://172.16.0.1/",
		"https://[::1]/",
		"https://169.254.1.1/",
		"http://localhost:2026/",
		"https://openviking.local/",
		"https://vault.internal/note",
		"file:///etc/passwd",
		"ftp://example.com/",
	}
	for _, u := range bad {
		if err := Public(u); err == nil {
			t.Errorf("Public(%q) accepted a non-public URL", u)
		}
	}
}

func TestScrapeRejectsPrivateBeforeRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := Scrape(ctx, nil, "test-key", "https://127.0.0.1/")
	if err == nil {
		t.Fatal("private URL was sent")
	}
}
