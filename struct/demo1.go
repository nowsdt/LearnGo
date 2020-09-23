package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	i int
}

func main() {
	var v myStruct
	var p *myStruct

	//p.i = 15
	p = &v
	(*p).i = 15
	(&v).i = 20
	//无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的 选择器符（selector-notation） 来引用结构体的字段：
	fmt.Println(v.i)
	fmt.Println(p.i)

	// 初始化
	//1. &struct1{a, b, c} 是一种简写，底层仍然会调用 new ()
	// 类型指针
	var ptr = &myStruct{15}
	my := myStruct{16}

	fmt.Println(ptr, my)

	//如果 File 是一个结构体类型，那么表达式 new(File) 和 &File{} 是等价的。
	//如果想知道结构体类型T的一个实例占用了多少内存，可以使用：size := unsafe.Sizeof(T{})。
}
