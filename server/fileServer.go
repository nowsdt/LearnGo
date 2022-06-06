package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("/files/path")))
	http.Handler("/", http.FileServer(http.Dir("/tmp")))
	http.HandleFunc("/static/", func(rw http.ResponseWriter, rq *http.Rquest) {
		http.ServeFile(rw, rq, rq.URL.Path[1:])
	})
}
