package main

import (
	"net"
	"net/http"
)

func lanAllowed(addr net.IP) bool {
	if addr == nil {
		return false
	}
	addr = addr.To16()
	if addr == nil {
		return false
	}
	// IPv4-mapped IPv6 must be judged as IPv4. IsPrivate treats those
	// addresses as public, which would lock a home client out or let a
	// mapped public address through a later check.
	if v4 := addr.To4(); v4 != nil {
		addr = v4
	}
	if addr.IsLoopback() {
		return true
	}
	if addr.IsLinkLocalUnicast() || addr.IsLinkLocalMulticast() || addr.IsUnspecified() || addr.IsMulticast() {
		return false
	}
	return addr.IsPrivate()
}

func lanOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "home network only", http.StatusForbidden)
			return
		}
		if !lanAllowed(net.ParseIP(host)) {
			http.Error(w, "home network only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
