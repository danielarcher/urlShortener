package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
	"urlShortener/handlers"
	"urlShortener/storages"
)

func main() {
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default cipher suite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}
	srv := &http.Server{
		Addr: "",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
		TLSConfig: tlsConfig,
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

