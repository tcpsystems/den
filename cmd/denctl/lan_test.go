package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLANAllowed(t *testing.T) {
	pass := []string{
		"192.168.1.20", "10.1.2.3", "172.16.0.4",
		"127.0.0.1", "::1", "fd12:3456::1",
		"::ffff:192.168.1.20", "::ffff:127.0.0.1",
	}
	fail := []string{
		"1.1.1.1", "8.8.8.8", "169.254.1.1", "fe80::1",
		"224.0.0.1", "0.0.0.0", "::", "::ffff:1.1.1.1",
		"2001:db8::1", "100.64.0.1",
	}
	for _, s := range pass {
		if !lanAllowed(net.ParseIP(s)) {
			t.Fatalf("%s must pass", s)
		}
	}
	for _, s := range fail {
		if lanAllowed(net.ParseIP(s)) {
			t.Fatalf("%s must fail", s)
		}
	}
	if lanAllowed(nil) {
		t.Fatal("empty address must fail")
	}
}

func TestWebBind(t *testing.T) {
	if err := webBind("127.0.0.1:8793"); err != nil {
		t.Fatal(err)
	}
	if err := webBind("192.168.1.20:8793"); err != nil {
		t.Fatal(err)
	}
	if err := webBind("[fd12:3456::1]:8793"); err != nil {
		t.Fatal(err)
	}
	for _, addr := range []string{"0.0.0.0:8793", "[::]:8793", "1.1.1.1:8793", ":8793", "not-an-addr"} {
		if err := webBind(addr); err == nil {
			t.Fatalf("%s must be refused", addr)
		}
	}
}

func TestLANOnlyRejectsPublicPeer(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	h := lanOnly(next)
	pub := httptest.NewRequest(http.MethodGet, "/", nil)
	pub.RemoteAddr = "1.1.1.1:40000"
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, pub)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("public peer %d", rec.Code)
	}
	bad := httptest.NewRequest(http.MethodGet, "/", nil)
	bad.RemoteAddr = "not-an-addr"
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, bad)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("unparsed peer %d", rec.Code)
	}
	home := httptest.NewRequest(http.MethodGet, "/", nil)
	home.RemoteAddr = "192.168.1.20:40000"
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, home)
	if rec.Code != http.StatusOK {
		t.Fatalf("home peer %d", rec.Code)
	}
}
