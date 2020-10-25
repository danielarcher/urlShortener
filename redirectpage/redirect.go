package redirectpage

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

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request)  {
	code := r.URL.Path[len("/go/"):]
	url, err := h.storage.Load(code)
	if err != nil {
		h.logger.Fatal("Unable to load: ", err)
	}
	h.logger.Printf("Decoded from code %s redirected to url %s \n", code, url)

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *Handler) WithLogger(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		h.logger.Printf("request processed in %d ms \n", time.Now().Sub(startTime).Milliseconds())
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/go/", h.WithLogger(h.Redirect))
}
