package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicing: %s \r\n", e)
		}
	}()
	badCall()
	fmt.Println("after bad call \r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
