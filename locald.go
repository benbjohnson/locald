package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	version = "0.1.0"
)

var port uint
func init() {
	flag.UintVar(&port, "port", 8000, "the port to listen on")
	flag.UintVar(&port, "p", 8000, "the port to listen on")
}

func main() {
	flag.Parse()

	// Retrieve the current working directory.
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Start the server.
	fmt.Printf("locald v%s\n", version)
	fmt.Printf("Serving %s on http://localhost:%d\n", path, port)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(path))))
}
