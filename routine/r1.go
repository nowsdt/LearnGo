package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go pump(ch1)
	go pump(ch2)

	go suck(ch1, ch2)

	time.Sleep(time.Second * 3)

}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1:", v)
		case v := <-ch2:
			fmt.Println("ch2:", v)
		}
	}
}
