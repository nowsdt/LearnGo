package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {

	}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {

	}
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println(" buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}
