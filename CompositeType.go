package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])

	for _, v := range a {
		fmt.Printf("index:%d \n", v)
	}

	var b = [3]int{1, 2}
	for _, v := range b {
		fmt.Printf("value:%v \n", v)
	}

	var c = [...]int{1, 2}

	for _, v := range c {
		fmt.Printf("value:%v \n", v)
	}

	type Currency int
	const (
		USD Currency = iota
		EUR
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "@", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	r := [...]int{10: -1}
	fmt.Println(r)

	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q \n", runes)

	var names []string

	fmt.Println(names)

	ages := make(map[string]int)
	ages["tom"] = 20
	ages["lucy"] = 21

	fmt.Println(ages)

	status := map[string]int{
		"run":  1,
		"stop": -1,
	}

	fmt.Println(status)

	delete(status, "run")
	status["stop"]++
	fmt.Println(status)

	st, ok := status["run"]

	fmt.Println(st, ok)

	fmt.Println()

	type Point struct {
		x, y int
	}

	type Circle struct {
		point  Point
		radius int
	}

	var circle Circle
	circle.point.x = 10
	circle.point.y = 15
	circle.radius = 5

	fmt.Println(circle)

	type ACircle struct {
		Point
		radius int
	}

	var aCircle ACircle
	aCircle.radius = 10
	aCircle.x = 15
	aCircle.y = 20

	aCircle = ACircle{Point{10, 15}, 10}

	aCircle = ACircle{Point: Point{10, 15}, radius: 5}

	fmt.Printf("%#v", aCircle)

}

/**
数组是传值
*/
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zLen := len(x) + 1

	if zLen < cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}

		z = make([]int, zLen, zCap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}
