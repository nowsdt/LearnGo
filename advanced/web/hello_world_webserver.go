package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/favicon.ico", Favicon)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside helloServer handler")
	fmt.Fprintf(w, "hello, "+req.URL.Path[1:])
}

func Favicon(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside Favicon handler")
	fmt.Fprintf(w, "Favicon, "+req.URL.Path)
}
