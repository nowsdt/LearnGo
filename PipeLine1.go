package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for {
			if x, ok := <-naturals; ok {
				squares <- x * x
			} else {
				break
			}
		}

		close(squares)

	}()

	for t := range squares {
		fmt.Println(t)
	}
}
