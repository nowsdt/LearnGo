package main

import (
	"fmt"
	"sync"
	"time"
)

/*var (
	mu sync.Mutex
	mapping = make(map[string]string)
)

func Lookup(key string) string  {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}*/

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func print(s string) {
	fmt.Println(s)
}

type Rocket struct {
}

func (r *Rocket) Launch() {
	fmt.Println("launch")
}

func main() {
	fmt.Println(cache)
	time.AfterFunc(5*time.Second, func() {
		print("go")
	})

	r := new(Rocket)
	time.AfterFunc(7*time.Second, func() {
		r.Launch()
	})

	fmt.Printf("%T", r.Launch) // 方法变量
	time.AfterFunc(9*time.Second, r.Launch)

	/*	p := Point{1, 2}
		distance := Point.Distance
		// 方法表达式， 第一个参数是接受者
		fmt.Println(distance(p, p))

		distanceP := (*Point).Distance
		// 方法表达式， 第一个参数是接受者
		fmt.Println(distanceP(&p, p)) */

	// 方法变量
	var op func(p, q Point) Point

	fmt.Println(op)
	time.Sleep(20 * time.Second)

}
