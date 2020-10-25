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
