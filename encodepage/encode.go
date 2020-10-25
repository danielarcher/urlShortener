package encodepage

import (
	"log"
	"net/http"
	"time"
	"urlShortener/storages"
)

type Handler struct {
	logger  *log.Logger
	storage storages.Storage
}

func NewHandler(l *log.Logger, s storages.Storage) *Handler {
	return &Handler{
		logger:  l,
		storage: s,
	}
}

func (h *Handler) Encode(w http.ResponseWriter, r *http.Request)  {
	url := findUrl(r)
	if url == "" {
		_, _ = w.Write([]byte("empty url"))
		return
	}

	code, err := h.storage.Save(url)
	if err != nil {
		h.logger.Fatal("Unable to save: ", err)
	}

	h.logger.Println("Encoded url:" + url + " to code:" + code)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("http://localhost:8080/go/" + code))
}

func (h *Handler) WithLogger(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		h.logger.Printf("request processed in %d ms \n", time.Now().Sub(startTime).Milliseconds())
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/encode", h.WithLogger(h.Encode))
}

func findUrl(r *http.Request) string {
	keys, _ := r.URL.Query()["url"]
	url := keys[0]

	return url
}
