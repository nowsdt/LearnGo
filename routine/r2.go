package main

import "fmt"

func main() {
	ch1 := make(chan int)

	go func(a, b int) {
		ch1 <- a + b
	}(1, 2)

	fmt.Println(<-ch1)

}
