package main

import (
	"context"
	_ "embed"
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/tcpsystems/den/internal/initd"
)

//go:embed openapi.json
var openAPI []byte

// denapi is the web publisher. It dials the daemon socket and serves HTTP
// on loopback. Browsers, dashboards, and later monitors use this process.
// They do not dial the socket themselves.
func loopback(addr string) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	if host != "127.0.0.1" && host != "::1" {
		return errPublic
	}
	return nil
}

var errPublic = publicAddrError{}

type publicAddrError struct{}

func (publicAddrError) Error() string { return "denapi listens on 127.0.0.1 only" }

func main() {
	addr := flag.String("addr", "127.0.0.1:8793", "loopback listen address")
	sock := flag.String("sock", initd.DefaultSock(), "daemon socket")
	flag.Parse()
	if err := loopback(*addr); err != nil {
		log.Fatal(err)
	}
	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = "den"
		},
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
				var d net.Dialer
				return d.DialContext(ctx, "unix", *sock)
			},
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(openAPI)
	})
	mux.Handle("/", proxy)
	log.Printf("denapi %s -> %s", *addr, *sock)
	srv := &http.Server{Addr: *addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Fatal(srv.ListenAndServe())
}
