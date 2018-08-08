package main

import (
	"net/http"
	"log"
)


func serveHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	e := "testetag"
	http.ServeFile(w, r, "./static/")
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
	log.Println('This is reaching sendHTTP')

	//if match := r.Header.Get("If-None-Match"); match != "" {
		//if strings.Contains(match, e) {
			//w.WriteHeader(http.StatusNotModified)
			//return
		//}
	//}
}

func main() {
	http.HandleFunc("/", serveHTTP)
	http.ListenAndServe(":8080", nil)
}
