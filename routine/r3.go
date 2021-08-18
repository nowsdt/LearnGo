package main

import "fmt"

func main() {

	ch1 := make(chan int)

	done := make(chan bool)

	go func(count int, out chan<- int) {
		for i := 0; i < count; i += 10 {
			ch1 <- i
		}
		close(out)
	}(100000, ch1)

	go func(in <-chan int, done chan<- bool) {
		for num := range in {
			fmt.Println(num)
		}
		done <- true
	}(ch1, done)

	<-done
	fmt.Println("done")
}
