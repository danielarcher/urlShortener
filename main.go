package main

import (
	"log"
	"net/http"
	"os"
	"urlShortener/encodepage"
	"urlShortener/redirectpage"
	"urlShortener/server"
	"urlShortener/storages"
)

var (
	serverAddress = os.Getenv("SERVER_PORT")
	storagePath = os.Getenv("STORAGE_PATH")
)

func main() {
	logger := log.New(os.Stdout, "urlShortener ", log.LstdFlags|log.Lshortfile)
	storage := storages.NewFileSystem("C:\\webserver") //"C:\\webserver"

	mux := http.NewServeMux()

	routes(logger, storage, mux)

	srv := server.New(mux, ":8080") //:8080
	logger.Println("Starting server")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server could not start: %v", err)
	}
}

func routes(logger *log.Logger, storage *storages.FileSystem, mux *http.ServeMux) {
	encode := encodepage.NewHandler(logger, storage)
	encode.SetupRoutes(mux)

	redirect := redirectpage.NewHandler(logger, storage)
	redirect.SetupRoutes(mux)
}
