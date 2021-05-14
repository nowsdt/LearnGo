package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		for {
			go func() {
				fmt.Println("hello world")
			}()
		}
	}()

	time.Sleep(10 * time.Second)
}
