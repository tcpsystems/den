package main

import (
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/tcpsystems/den/internal/initd"
)

// defaultSock is the socket the daemon presents. TUI and other local
// clients dial this. The browser cannot dial a Unix socket, so the same
// handlers are also served on 127.0.0.1 for the web UI only.
func defaultSock() string {
	return initd.DefaultSock()
}

func listenUnix(path string, h http.Handler) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	_ = os.Remove(path)
	ln, err := net.Listen("unix", path)
	if err != nil {
		return err
	}
	if err := os.Chmod(path, 0o600); err != nil {
		ln.Close()
		return err
	}
	srv := &http.Server{Handler: h, ReadHeaderTimeout: 5 * time.Second}
	return srv.Serve(ln)
}

func listenTCP(addr string, h http.Handler) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := &http.Server{Handler: h, ReadHeaderTimeout: 5 * time.Second}
	return srv.Serve(ln)
}
