package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Printf("received request %s %s [%s] \n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *req

	if clientIp, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIp = strings.Join(prior, ","+", "+clientIp)
		}

		outReq.Header.Set("X-Forwarded-For", clientIp)

		res, err := transport.RoundTrip(outReq)
		if err != nil {
			rw.WriteHeader(http.StatusBadGateway)
			return
		}

		for key, value := range res.Header {
			for _, v := range value {
				rw.Header().Add(key, v)
			}
		}
		rw.WriteHeader(res.StatusCode)
		io.Copy(rw, res.Body)
		res.Body.Close()
	}
}

func main() {
	log.Println("server on :8080")
	http.Handle("/", &Pxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
