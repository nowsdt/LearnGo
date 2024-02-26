package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	bsh = "/bin/sh"
	url = "http://www.haeea.cn/GovSearch/News/Index?words=%E6%B2%B3%E5%8D%97%E7%9C%81%E6%99%AE%E9%80%9A%E9%AB%98%E6%A0%A1%E6%8B%9B%E7%94%9F%E5%BD%95%E5%8F%96%E6%8E%A7%E5%88%B6%E5%88%86%E6%95%B0%E7%BA%BF&page=1"
	c   = "-c"

	c0 = `osascript -e 'display notification "分数线出来了" with title "分数线出来了"'`
	c1 = `osascript -e 'display notification "出错了" with title "出错了"'`
	c2 = `osascript -e 'display alert "分数线已出！！" as critical'`
)

func main() {
	timer := time.Tick(time.Second * 5)
	runtime.GOMAXPROCS(1)

	var count int32 = 0

	first := true
	for range timer {

		resp, err := http.Get(url)

		log.Println("query count:", count)

		count++

		if err != nil {
			exec.Command(bsh, c, c1).Start()
			log.Println("error:", err)
		}

		if resp.StatusCode != http.StatusOK {
			exec.Command(bsh, c, c1).Start()
			log.Println("errorCode:", resp.StatusCode)
		}

		if resp.StatusCode == http.StatusOK {
			bodyData, _ := ioutil.ReadAll(resp.Body)
			if strings.Contains(string(bodyData), "2022") {
				fmt.Println("2022", string(bodyData))
				if first {
					exec.Command(bsh, c, c2).Start()
				}
				first = false
				for {
					exec.Command(bsh, c, c0).Start()
					time.Sleep(time.Second * 5)
				}
			}
		}
	}

}

func init() {
	log.SetFlags(log.LstdFlags)
}
