package main

import "net"
import "testing"

func TestLANAllowed(t *testing.T) {
	if !lanAllowed(net.ParseIP("192.168.1.20")) || !lanAllowed(net.ParseIP("127.0.0.1")) {
		t.Fatal("home and loopback addresses must pass")
	}
	if lanAllowed(net.ParseIP("1.1.1.1")) || lanAllowed(nil) {
		t.Fatal("public and empty addresses must fail")
	}
}
