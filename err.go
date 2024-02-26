package main

import (
	"fmt"
	"time"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicing: %s \r\n", e)
		}
	}()
	//badCall()
	fmt.Println("after bad call \r\n")

	fmt.Println(time.Now().Add(time.Hour * 34))
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
