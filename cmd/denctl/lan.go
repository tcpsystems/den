package main

import (
	"net"
	"net/http"
)

func lanAllowed(addr net.IP) bool {
	if addr == nil || addr.IsLoopback() {
		return addr != nil && addr.IsLoopback()
	}
	return addr.IsPrivate() && !addr.IsLoopback()
}

func lanOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			host = r.RemoteAddr
		}
		if !lanAllowed(net.ParseIP(host)) {
			http.Error(w, "home network only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
