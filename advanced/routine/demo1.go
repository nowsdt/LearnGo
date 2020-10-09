package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go send(ch)
	go receive(ch)

	time.Sleep(2 * time.Second)
}

func send(ch chan<- string) {
	ch <- "tom"
	ch <- "cat"
	ch <- "lucy"
	ch <- "lily"
}

func receive(ch <-chan string) {
	var input string

	for {
		input = <-ch
		fmt.Println("input:", input)
	}
}
