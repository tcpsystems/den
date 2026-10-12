package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// SearchEndpoint is OmniRoute's search gateway.
// The live router is https://ai.tcp.systems/v1/search.
const SearchEndpoint = "https://ai.tcp.systems/v1/search"

// MaxHits is the most results one search keeps.
const MaxHits = 8

// Hit is one public result from OmniRoute search.
type Hit struct {
	Title   string
	URL     string
	Snippet string
}

type searchRequest struct {
	Query      string `json:"query"`
	MaxResults int    `json:"max_results"`
}

type searchResponse struct {
	Results []struct {
		Title   string `json:"title"`
		URL     string `json:"url"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"results"`
	Data []struct {
		Title   string `json:"title"`
		URL     string `json:"url"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"data"`
	Organic []struct {
		Title   string `json:"title"`
		URL     string `json:"url"`
		Link    string `json:"link"`
		Snippet string `json:"snippet"`
	} `json:"organic"`
}

// Search asks OmniRoute for public pages about query.
// key is the OmniRoute bearer token. It is not written into the result or an error.
func Search(ctx context.Context, client *http.Client, key, query string) ([]Hit, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, errors.New("fetch: empty query")
	}
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("fetch: missing key")
	}
	if client == nil {
		client = &http.Client{Timeout: 60 * time.Second}
	}
	body, err := json.Marshal(searchRequest{Query: query, MaxResults: MaxHits})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, SearchEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch: search status %d", resp.StatusCode)
	}
	var parsed searchResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, err
	}
	rows := parsed.Results
	if len(rows) == 0 {
		rows = parsed.Data
	}
	if len(rows) == 0 {
		rows = parsed.Organic
	}
	hits := make([]Hit, 0, len(rows))
	seen := map[string]bool{}
	for _, row := range rows {
		link := row.URL
		if link == "" {
			link = row.Link
		}
		if link == "" || seen[link] || Public(link) != nil {
			continue
		}
		seen[link] = true
		hits = append(hits, Hit{Title: row.Title, URL: link, Snippet: row.Snippet})
		if len(hits) == MaxHits {
			break
		}
	}
	return hits, nil
}

// Brief is a search plus the markdown of the public hits that could be read.
type Brief struct {
	Query string
	Hits  []Hit
	Pages []Page
}

// Research searches, then reads up to readN distinct public pages.
// searchKey is the OmniRoute key. fetchKey is the Firecrawl key.
func Research(ctx context.Context, client *http.Client, searchKey, fetchKey, query string, readN int) (Brief, error) {
	if readN < 1 {
		readN = 1
	}
	if readN > 4 {
		readN = 4
	}
	hits, err := Search(ctx, client, searchKey, query)
	if err != nil {
		return Brief{}, err
	}
	brief := Brief{Query: query, Hits: hits}
	for _, hit := range hits {
		if len(brief.Pages) == readN {
			break
		}
		page, err := Scrape(ctx, client, fetchKey, hit.URL)
		if err != nil || strings.TrimSpace(page.Markdown) == "" {
			continue
		}
		brief.Pages = append(brief.Pages, page)
	}
	return brief, nil
}
