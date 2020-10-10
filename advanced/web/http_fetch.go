package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url = "https://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", url, err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", url, err)
	}
	fmt.Printf("Got: %q", string(data))
}
