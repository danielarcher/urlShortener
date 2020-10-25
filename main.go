package main

import (
	"log"
	"net/http"
	"time"
	"urlShortener/handlers"
	"urlShortener/storages"
)

func main() {
	srv := &http.Server{
		Addr: "",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
		Handler: routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server could not start: %v", err)
	}
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	storage := storages.FileSystem{
		Path: "C:\\webserver",
	}
	mux.Handle("/", handlers.Home())
	mux.Handle("/encode/", handlers.Encode(storage))
	mux.Handle("/go/", handlers.Redirect(storage))

	return mux
}

