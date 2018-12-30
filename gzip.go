package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"io"
	"net"
	"net/http"
)

type gzipResponseWriter struct {
	io.WriteCloser
	http.ResponseWriter
}

func newGzipResponseWriter(w http.ResponseWriter) *gzipResponseWriter {
	gz := gzip.NewWriter(w)
	return &gzipResponseWriter{WriteCloser: gz, ResponseWriter: w}
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.WriteCloser.Write(b)
}

// Hijack satisfies the http.Hijacker interface. It is needed when switching
// protocols (say, to a websocket connection).
func (w gzipResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("gzipResponseWriter: ResponseWriter does not satisfy http.Hijacker interface")
	}
	return h.Hijack()
}
