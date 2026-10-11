package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/tcpsystems/den/internal/console"
)

//go:embed web/index.html
var page embed.FS

func main() {
	addr := flag.String("addr", "127.0.0.1:8793", "loopback listen address")
	repo := flag.String("repo", ".", "git repository to read")
	flag.Parse()
	if err := serve(*addr, *repo); err != nil {
		log.Fatal(err)
	}
}

func serve(addr, repo string) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	if host != "127.0.0.1" && host != "::1" {
		return errPublic
	}
	sub, err := fs.Sub(page, "web")
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(sub)))
	mux.HandleFunc("/api/state", func(w http.ResponseWriter, r *http.Request) {
		writeState(w, r, repo)
	})
	mux.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "stream unsupported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		tick := time.NewTicker(5 * time.Second)
		defer tick.Stop()
		send := func() bool {
			payload, err := stateJSON(r.Context(), repo)
			if err != nil {
				return false
			}
			fmt.Fprintf(w, "data: %s\n\n", payload)
			flusher.Flush()
			return true
		}
		if !send() {
			return
		}
		for {
			select {
			case <-r.Context().Done():
				return
			case <-tick.C:
				if !send() {
					return
				}
			}
		}
	})
	server := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Printf("denctl %s", addr)
	return server.ListenAndServe()
}

func writeState(w http.ResponseWriter, r *http.Request, repo string) {
	payload, err := stateJSON(r.Context(), repo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func stateJSON(parent context.Context, repo string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(parent, 8*time.Second)
	defer cancel()
	return json.Marshal(console.Collect(ctx, repo))
}

var errPublic = publicAddr{}

type publicAddr struct{}

func (publicAddr) Error() string { return "denctl listens on 127.0.0.1 only" }
