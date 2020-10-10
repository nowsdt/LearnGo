package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com",
	"https://www.sogou.com",
	"https://www.qq.com",
	"http://www.google.com/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
