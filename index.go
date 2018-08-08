package main

import (
	"log"
	"net/http"
	"strings"
)

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "text/html")
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
	e := "testetag"
	w.Header().Set("Etag", e)
	http.ServeFile(w, r, "./static/")
	log.Println("This is reaching sendHTTP")

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", serveHTTP)
	http.ListenAndServe(":8080", nil)
}
