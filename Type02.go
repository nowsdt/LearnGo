package main

import "fmt"

func main() {

	arr1 := new([5]int)
	arr1[0] = 15
	arr1[4] = 14

	// 值copy
	arr2 := *arr1
	fmt.Printf("%p-%p\n", arr1, &arr2)

	fmt.Println(arr1)

	modify1(*arr1)
	fmt.Println(arr1)

	modify2(arr1)
	fmt.Println(arr1)

	modify3((*arr1)[0:])
	fmt.Println(arr1)

}

func modify1(arr [5]int) {
	arr[0] = 80
}

// 指针 引用传递
func modify2(arr *[5]int) {
	(*arr)[0] = 81
}

// slice 引用传递
func modify3(arr []int) {
	arr[0] = 82
}
