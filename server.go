package main

import (
	"log"
	"flag"
	"net/http"
)

var (
	port = flag.String("port", ":8080", "listen address")
	dir = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *port)
	log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir(*dir))))
}