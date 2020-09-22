package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	fmt.Println("sizeOf:", unsafe.Sizeof(x), " Alignof:", unsafe.Alignof(x))
	fmt.Println("a sizeOf:", unsafe.Sizeof(x.a), " a Alignof:", unsafe.Alignof(x.a))
	fmt.Println("b sizeOf:", unsafe.Sizeof(x.b), " b Alignof:", unsafe.Alignof(x.b))
	fmt.Println("c sizeOf:", unsafe.Sizeof(x.c), " c Alignof:", unsafe.Alignof(x.c))

	fmt.Println(x)

	pointer := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pointer = 15
	fmt.Println(x)

	// 错误
	// tmp对应的uintptr是一个非指针变量，也就是一个值垃圾回收时候如果指针变化，它不会更新
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 61

	fmt.Println(x)
}
