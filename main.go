package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8080", "HTTP service address")
	flag.Parse()
	s, err := newServer(useStdLibOptions())
	if err != nil {
		log.Fatalf("newServer: %v", err)
	}
	log.Fatal(http.ListenAndServe(*addr, s))
}
