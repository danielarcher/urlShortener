package main

import (
	"fmt"
	"log"
	"net/http"
	"urlShortener/handlers"
	"urlShortener/server"
	"urlShortener/shortener"
	"urlShortener/storages"
)

func main() {
	srv := server.New(routes(),":8080")
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
	handler := shortener.NewHandler(storage)
	mux.Handle("/", handlers.Home())
	mux.HandleFunc("/encode/", handler.Encode)
	mux.HandleFunc("/go/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/go/"):]
		url,err := handler.Decode(code)
		if err != nil {
			_, _ = fmt.Fprintln(w, "not found")
			return
		}
		http.Redirect(w, r, url, 302)
	})

	return mux
}
