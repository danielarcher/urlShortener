package handlers

import (
	"fmt"
	"log"
	"net/http"
	"urlShortener/storages"
)

func Home() http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		_,_ = fmt.Fprint(w, "welcome")
	}

	return http.HandlerFunc(handleFunc)
}

func Encode(s storages.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		url := findUrl(r)
		if url == "" {
			_, _ = w.Write([]byte("empty url"))
			return
		}
		code := saveUrl(s, url)
		writeToUser(w, url, code)
	}

	return http.HandlerFunc(handleFunc)
}

func writeToUser(w http.ResponseWriter, url string, code string) {
	log.Println("Encoded url:" + url + " to code:" + code)
	_, _ = w.Write([]byte("/go/" + code))
}

func saveUrl(s storages.Storage, url string) string {
	code, err := s.Save(url)
	if err != nil {
		log.Fatal("Unable to save: ", err)
	}
	return code
}

func findUrl(r *http.Request) string {
	keys,_ := r.URL.Query()["url"]
	url := keys[0]

	return url
}

func Redirect(s storages.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/go/"):]
		url, err := s.Load(code)
		if err != nil {
			_,_ = fmt.Fprintln(w, "not found")
			log.Fatal("Unable to load: ", err)
			return
		}

		log.Println("Redirecting code:"+code+" url:"+url)
		http.Redirect(w, r, url, 302)
	}

	return http.HandlerFunc(handleFunc)
}
