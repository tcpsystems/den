package fetch

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestSearchDropsPrivateHits(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.String() != SearchEndpoint {
			t.Errorf("url %s", r.URL)
		}
		if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ") {
			t.Error("missing bearer")
		}
		body := `{"results":[
			{"title":"local","url":"https://127.0.0.1/secret","snippet":"no"},
			{"title":"lan","url":"http://192.168.0.166/","snippet":"no"},
			{"title":"public","url":"https://example.com/harness","snippet":"yes"}]}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	hits, err := Search(ctx, client, "k", "agent harness")
	if err != nil {
		t.Fatal(err)
	}
	if len(hits) != 1 || hits[0].URL != "https://example.com/harness" {
		t.Fatalf("hits %+v", hits)
	}
}

func TestSearchEmptyQuery(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if _, err := Search(ctx, nil, "k", "  "); err == nil {
		t.Fatal("empty query accepted")
	}
}
