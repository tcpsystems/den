package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOpenAPIIsLoopbackOnly(t *testing.T) {
	if err := loopback("127.0.0.1:8793"); err != nil {
		t.Fatal(err)
	}
	if err := loopback("0.0.0.0:8793"); err == nil {
		t.Fatal("public address accepted")
	}
	req := httptest.NewRequest(http.MethodGet, "/openapi.json", nil)
	rec := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(openAPI)
	}).ServeHTTP(rec, req)
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "openapi") {
		t.Fatalf("%d %s", rec.Code, rec.Body.String())
	}
}

