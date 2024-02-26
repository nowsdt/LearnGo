package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"sync"
)

func main() {
	sem := sync.WaitGroup{}

	// 请求数量
	requestCnt := 2
	sem.Add(requestCnt)

	url := ""
	go func() {
		defer sem.Done()
		param := ``
		values := url2.Values{}
		values.Add("orderId", param)
		//get(url + "/backdoor/refundCallBack?" + values.Encode())
		//get(url + "/test/testPayCallback?" + values.Encode())
		get(url + "/?" + values.Encode())
		//post(url+"/backdoor/returnCallback", param)
	}()

	go func() {
		defer sem.Done()
		param := ``
		values := url2.Values{}
		values.Add("msg", param)
		get(url + "" + values.Encode())
	}()

	sem.Wait()
}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	result := responseToString(resp)
	log.Printf("url:%s, resp:%s\n", url, result)

}

func post(url, params string) {
	resp, err := http.Post(url, "application/json; charset=UTF-8", bytes.NewBuffer([]byte(params)))
	if err != nil {
		log.Fatal(err)
	}

	result := responseToString(resp)
	log.Printf("url:%s, resp:%s\n", url, result)
}

func responseToString(response *http.Response) string {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
