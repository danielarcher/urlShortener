package encodepage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"urlShortener/storages"
)

type Handler struct {
	logger  *log.Logger
	storage storages.Storage
}

type EncodeRequest struct {
	Url string `json:"url"`
}
type EncodeResponse struct {
	RedirectUrl string `json:"redirect_url"`
}

func NewHandler(l *log.Logger, s storages.Storage) *Handler {
	return &Handler{
		logger:  l,
		storage: s,
	}
}

func (h *Handler) Encode(w http.ResponseWriter, r *http.Request)  {
	var encReq EncodeRequest
	data,_ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(data, &encReq)
	h.logger.Println(encReq.Url)

	if encReq.Url == "" {
		_, _ = w.Write([]byte("empty url"))
		return
	}

	code, err := h.storage.Save(encReq.Url)
	if err != nil {
		h.logger.Fatal("Unable to save: ", err)
	}

	h.logger.Println("Encoded url:" + encReq.Url + " to code:" + code)

	var encRes = EncodeResponse{
		RedirectUrl: "http://localhost:8080/go/"+code,
	}

	data, _ = json.Marshal(encRes)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
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
