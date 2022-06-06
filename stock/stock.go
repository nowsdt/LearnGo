package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	sina_url    = "http://hq.sinajs.cn/list=%s"
	tencent_url = "http://qt.gtimg.cn/q=%s"

	shanghai = "sh"
	shenzhen = "sz"

	czbk = "003032"
	lbgj = "600559"
)

const (
	name = iota
	open
	preClose
	cur
	high
	low
	buy1
	sale1
	dealCount
	dealAmount
)

func main() {
	var err error

	url := fmt.Sprint(sina_url, shenzhen+czbk)
	url = "http://hq.sinajs.cn/list=sh600559,sh600905"
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("referer", "http://finance.sina.com.cn")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	all, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes.Runes(all)))
	fmt.Println(resp.Header)

}
