package main

import (
	"net/http"
	"urlShortener/handlers"
	"urlShortener/storages"
)

func main() {
	routes()
	serve()
}

func routes() {
	storage := storages.FileSystem{
		Path: "C:\\webserver",
	}
	http.Handle("/", handlers.Home())
	http.Handle("/encode/", handlers.Encode(storage))
	http.Handle("/go/", handlers.Redirect(storage))
}

func serve() {
	_ = http.ListenAndServe(":8080", nil)
}

