package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

type testLogger struct {
	t *testing.T
}

func newTestLogger(t *testing.T) *testLogger {
	return &testLogger{t: t}
}

func (l *testLogger) Criticalf(format string, args ...interface{}) {
	l.t.Logf("CRITICAL: "+format, args...)
}
func (l *testLogger) Debugf(format string, args ...interface{}) {
	l.t.Logf("DEBUG: "+format, args...)
}
func (l *testLogger) Errorf(format string, args ...interface{}) {
	l.t.Logf("ERROR: "+format, args...)
}
func (l *testLogger) Infof(format string, args ...interface{}) {
	l.t.Logf("INFO: "+format, args...)
}
func (l *testLogger) Warningf(format string, args ...interface{}) {
	l.t.Logf("WARNING: "+format, args...)
}

var _ logger = (*testLogger)(nil)

func useTestDefaults(t *testing.T) func(*server) error {
	return func(s *server) error {
		s.log = newTestLogger(t)
		s.mux = &http.ServeMux{}
		s.client = http.DefaultClient
		return nil
	}
}

var (
	testServer http.Handler
	addr       string
	once       sync.Once
)

func startServer(t *testing.T) {
	once.Do(func() {
		s, err := newServer(useTestDefaults(t))
		if err != nil {
			t.Fatalf("newServer(): %v", err)
		}
		testServer := httptest.NewServer(s)
		addr = testServer.Listener.Addr().String()
	})
}

func TestHealthz(t *testing.T) {
	startServer(t)
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
