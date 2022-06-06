package main

import (
	"fmt"
	"io"
	"net/http"
)

const SINA_URL = "http://hq.sinajs.cn/%s"

type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	outReq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(SINA_URL, req.RequestURI), nil)
	outReq.Header.Add("referer", "http://finance.sina.com.cn")

	fmt.Println(req.URL)

	res, err := http.DefaultTransport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	defer res.Body.Close()

	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)

}

func main() {
	fmt.Println("server on :8081")
	http.ListenAndServe("0.0.0.0:8081", &Pxy{})
}
