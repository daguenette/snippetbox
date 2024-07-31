package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /snippet/view/{id}", snippetViewHandler)
	mux.HandleFunc("GET /snippet/create", snippetFormHandler)
	mux.HandleFunc("POST /snippet/create", snippetCreateHandler)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
