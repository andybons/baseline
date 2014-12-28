package main

import (
	"flag"
	"net/http"
)

func main() {
	var port = flag.String("port", ":8080", "HTTP service port")
	http.ListenAndServe(*port, newServer())
}
