package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)

	go suck(ch1, ch2)

	time.Sleep(1 * time.Second)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case i := <-ch1:
			fmt.Println("ch1:", i)
		case i := <-ch2:
			fmt.Println("ch2:", i)
		}
	}

}
