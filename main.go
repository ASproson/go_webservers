package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux() // http request multiplexer (router)

	// Catch all route to return 404 for any path
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Listening on port %s", port)

	log.Fatal(server.ListenAndServe())
}
