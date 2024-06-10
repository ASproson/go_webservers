package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux() // http request multiplexer (router)

	mux.Handle("/", http.FileServer(http.Dir("."))) // Server files on root dir

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Listening on port %s", port)

	log.Fatal(server.ListenAndServe())
}
