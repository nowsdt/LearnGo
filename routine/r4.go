package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	go func() {
		for {
			//ch1 <-rand.Intn(2)
			select {
			case ch1 <- 0:
			case ch1 <- 1:
			}
		}
	}()

	idx := 0

	go func() {
		for {
			fmt.Println(<-ch1)
			idx++
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println(idx)
}
