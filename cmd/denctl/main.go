package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
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
		w.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(r.Context(), 8*time.Second)
		defer cancel()
		json.NewEncoder(w).Encode(console.Collect(ctx, repo))
	})
	server := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Printf("denctl %s", addr)
	return server.ListenAndServe()
}

var errPublic = publicAddr{}

type publicAddr struct{}

func (publicAddr) Error() string { return "denctl listens on 127.0.0.1 only" }
