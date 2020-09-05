package main

import (
	"fmt"
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
