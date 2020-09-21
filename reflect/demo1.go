package main

import (
	"fmt"
	"reflect"
)

func main() {

	t := reflect.TypeOf(3)

	fmt.Println(t.String())
	fmt.Println(t)
	fmt.Println("===================")

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	kind := v.Kind()
	fmt.Println("kind", kind)

	x := v.Interface()
	i := x.(int)
	fmt.Println(i)
	fmt.Println(x)

}
