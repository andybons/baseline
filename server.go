package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type server struct {
	mux    mux
	log    logger
	client *http.Client
}

func useStdLibOptions() func(*server) error {
	return func(s *server) error {
		s.mux = http.NewServeMux()
		s.log = newStdLogger()
		s.client = &http.Client{Timeout: 30 * time.Second}
		return nil
	}
}

// newServer allocates and returns a new server.
func newServer(options ...func(*server) error) (*server, error) {
	s := &server{}
	for _, o := range options {
		if err := o(s); err != nil {
			return nil, err
		}
	}
	if s.mux == nil {
		return nil, fmt.Errorf("must provide an option func that specifies a mux")
	}
	if s.log == nil {
		return nil, fmt.Errorf("must provide an option func that specifies a logger")
	}
	if s.client == nil {
		return nil, fmt.Errorf("must provide an option func that specifies an *http.Client")
	}
	s.init()
	return s, nil
}

// init sets up a server by performing tasks like mapping
// path endpoints to handler functions.
func (s *server) init() {
	s.mux.HandleFunc("/healthz", s.handleHealthz)
}

func (s *server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "ok")
}

// ServeHTTP satisfies the http.Handler interface. It will compress all
// responses if the appropriate request headers are set.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		s.mux.ServeHTTP(w, r)
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gzw := newGzipResponseWriter(w)
	defer gzw.Close()
	s.mux.ServeHTTP(gzw, r)
}
