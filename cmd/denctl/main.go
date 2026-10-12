package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/tcpsystems/den/internal/console"
)

//go:embed web/index.html
var page embed.FS

func main() {
	sock := flag.String("sock", defaultSock(), "unix socket presented to local clients")
	repo := flag.String("repo", ".", "git repository to read")
	flag.Parse()
	if err := serve(*sock, *repo); err != nil {
		log.Fatal(err)
	}
}

func serve(sock, repo string) error {
	sub, err := fs.Sub(page, "web")
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(sub)))
	mux.HandleFunc("/api/state", func(w http.ResponseWriter, r *http.Request) {
		writeState(w, r, repo)
	})
	mux.HandleFunc("/api/events", eventsHandler(repo))
	log.Printf("denctl sock %s", sock)
	return listenUnix(sock, mux)
}

func eventsHandler(repo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "stream unsupported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		// MDN event stream: text/event-stream, fields split by a blank line.
		// A leading colon is a comment and keeps the connection alive when
		// no state event is due. retry is the reconnect delay in milliseconds.
		fmt.Fprintf(w, "retry: 3000\n\n")
		flusher.Flush()
		tick := time.NewTicker(200 * time.Millisecond)
		keep := time.NewTicker(15 * time.Second)
		defer tick.Stop()
		defer keep.Stop()
		var last string
		var id int
		send := func() bool {
			payload, err := stateJSON(r.Context(), repo)
			if err != nil {
				return false
			}
			current := string(payload)
			if current == last {
				return true
			}
			last = current
			id++
			fmt.Fprintf(w, "id: %d\nevent: state\ndata: %s\n\n", id, payload)
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
			case <-keep.C:
				fmt.Fprintf(w, ": keepalive\n\n")
				flusher.Flush()
			case <-tick.C:
				if !send() {
					return
				}
			}
		}
	}
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
