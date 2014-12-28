package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

var (
	testServer http.Handler
	addr       string
	once       sync.Once
)

func startServer() {
	testServer := httptest.NewServer(newServer())
	addr = testServer.Listener.Addr().String()
}

func TestHealthz(t *testing.T) {
	once.Do(startServer)
	resp, err := http.Get("http://" + addr + "/healthz")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code to be %d, got %d", http.StatusOK, resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expected no error reading response body, got %v", err)
	}
	expected := "ok"
	if strings.TrimSpace(string(b)) != expected {
		t.Errorf("expected response body to be %q, got %q", expected, b)
	}
}
