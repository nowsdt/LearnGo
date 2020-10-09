package main

import (
	fmt "fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go pump(ch)
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	time.Sleep(2 * time.Second)
}

func pump(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
}
