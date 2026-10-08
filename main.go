package main

import (
	"golang-url-shortner/handlers"
	"golang-url-shortner/storage"
	"net/http"
)

func main() {
	// initialise the store
	store := storage.NewStore()

	// initialise the handler
	handler := handlers.NewHandler(store)

	// Define routes
	http.HandleFunc("/shorten", handler.ShortenUrl) //post shorten
	http.HandleFunc("/", handler.RedirectURL)

	// Start the server
	port :=":8080"
	println("Serving on port " , port)

	if err := http.ListenAndServe(port, nil); err != nil {
		println("Error starting server:", err)
	}
}
