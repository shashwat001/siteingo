package main

import (
	"net/http"
	"strings"
)

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	e := "testetag"
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
	http.ServeFile(w, r, "./static/")
}

func main() {
	http.HandleFunc("/", serveHTTP)
	http.ListenAndServe(":8080", nil)
}
