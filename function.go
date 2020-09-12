package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"time"
)

func main() {

	var arr = []int{1, 2, 3}
	one, two, three := top5(arr)
	fmt.Println(one, two, three)

	fmt.Println(time.Now())

	str := "tom"
	for i := 0; i < 2; i++ {
		fmt.Printf("%*s\n", i*2, "", str)
	}
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus ":  {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures ", "computer organization"},
		"programming languages": {"data structures", "compute organization"},
	}

	for i, course := range topSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	fmt.Println()
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
	fmt.Println(sum(1, 2))

	s := sum
	fmt.Println(s(1, 2))

	i := 0
	defer fmt.Println("defer", i)
	i++

	f1 := func(i, j int) { fmt.Println("匿名函数1", i, j) }
	f1(1, 2)
	// 直接调用的匿名函数  js中叫立即函数
	func(i, j, k int) { fmt.Println("匿名函数2", i, j, k) }(3, 6, 9)
	// 匿名函数的值 是一个内存地址
	fmt.Println(f1)

	f2 := func() (ret int) {
		defer func() {
			ret++
		}()
		return 1
	}

	fmt.Println(f2())

	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	where := func() {
		caller, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%s:%d", caller, file, line)
	}

	where()

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	fmt.Printf("%T\n", a)

	items := [...]int{10, 20, 30, 40, 50}

	itemSlice := items[:]
	fmt.Println("itemSlice before:", itemSlice)
	itemSlice = append(itemSlice[:1], itemSlice[2:]...)
	fmt.Println("itemSlice after:", itemSlice)
	for idx, item := range items {
		item *= 2
		if idx == len(items)-1 {
			defer fmt.Println("item", item)
		}
	}
	fmt.Println(items)

	ints := items[0:5]
	fmt.Println(len(ints))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func top5(arr []int) (one int, two int, three int) {
	one = arr[0]
	two = arr[1]
	three = arr[2]
	return
}

func squares() func() int {
	var x int
	fmt.Printf("value: %v\n", x)
	return func() int {
		x++
		return x * x
	}
}

func topSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

func sum(num ...int) (sum int) {
	for _, n := range num {
		sum += n
	}
	return
}
