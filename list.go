package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dirname := "I:\\pics\\WeiXin\\remove"
	dir, _ := ioutil.ReadDir(dirname)
	for _, info := range dir {
		fmt.Println("rm", info.Name())
	}
}
