package main

import "fmt"

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func getInteger() int {
	return <-resume
}

func main() {
	resume = integers()
	fmt.Println(getInteger())
	fmt.Println(getInteger())
	fmt.Println(getInteger())
	fmt.Println(getInteger())
}