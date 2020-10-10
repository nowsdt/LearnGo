package main

import (
	"io"
	"log"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

func simpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func formServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case http.MethodGet:
		io.WriteString(w, form)
	case http.MethodPost:
		io.WriteString(w, req.FormValue("in"))
	}
}

type HandleFnc func(http.ResponseWriter, *http.Request)

func logPanics(function HandleFnc) HandleFnc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v ", req.RemoteAddr, x)
			}
		}()
		function(w, req)
	}
}

func main() {
	/*	http.HandleFunc("/test1", simpleServer)
		http.HandleFunc("/test2", formServer)*/
	http.HandleFunc("/test1", logPanics(simpleServer))
	http.HandleFunc("/test2", logPanics(formServer))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
