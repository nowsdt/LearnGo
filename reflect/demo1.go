package main

import (
	"fmt"
	"reflect"
	"strconv"
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
	quote := strconv.Quote("hello")
	fmt.Println(quote)

	type tom struct {
		name string
		age  int
	}

	t2 := tom{name: "tom"}
	of := reflect.ValueOf(t2)
	fmt.Println("NumField:", of.NumField())
	fmt.Println("0 Field:", of.Field(0))

	x1 := 2
	d := reflect.ValueOf(&x1).Elem()

	px := d.Addr().Interface().(*int)

	*px = 3

	fmt.Println("x1:", x1)

	d.Set(reflect.ValueOf(15))
	fmt.Println("x1:", x1)

}
