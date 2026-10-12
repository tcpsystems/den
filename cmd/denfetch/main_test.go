package main

import "testing"

func TestFetchRefusesPrivateBeforeRead(t *testing.T) {
	if err := refusePrivate("http://127.0.0.1/secret"); err == nil {
		t.Fatal("loopback accepted")
	}
	if err := refusePrivate("https://example.com/page"); err != nil {
		t.Fatal(err)
	}
}
