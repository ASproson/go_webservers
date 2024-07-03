package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const filePathRoot = "."

	mux := http.NewServeMux() // http request multiplexer (router)

	fileServer := http.FileServer(http.Dir(filePathRoot))
	mux.Handle("/app/*", http.StripPrefix("/app", fileServer))
	mux.HandleFunc("/healthz", healthzHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Listening on port %s", port)

	log.Fatal(server.ListenAndServe())
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	const contentType = "text/plain; charset=utf-8"
	const statusCode = 200
	const bodyText = "OK"

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write([]byte(bodyText))

}
