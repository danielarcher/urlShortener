package shortener

import (
	"log"
	"net/http"
	"urlShortener/storages"
)

type Handler struct {
	storage storages.Storage
}

func NewHandler(s storages.Storage) *Handler {
	return &Handler{
		storage: s,
	}
}

func (h Handler) Encode(w http.ResponseWriter, r *http.Request) {
	url := findUrl(r)
	if url == "" {
		_, _ = w.Write([]byte("empty url"))
		return
	}
	code := saveUrl(h.storage, url)
	writeToUser(w, url, code)
}

func (h Handler) Decode(code string) (string,error) {
	url, err := h.storage.Load(code)
	if err != nil {
		log.Fatal("Unable to load: ", err)
	}
	log.Println("Decoded from code:"+code+" to url:"+url)
	return url,err
}


func writeToUser(w http.ResponseWriter, url string, code string) {
	log.Println("Encoded url:" + url + " to code:" + code)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/go/" + code))
}

func saveUrl(s storages.Storage, url string) string {
	code, err := s.Save(url)
	if err != nil {
		log.Fatal("Unable to save: ", err)
	}
	return code
}

func findUrl(r *http.Request) string {
	keys, _ := r.URL.Query()["url"]
	url := keys[0]

	return url
}
