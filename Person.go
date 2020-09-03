package main

import "fmt"

type Person struct {
	name, address string
	age           int
}

func main() {
	person := &Person{"tom", "beijing", 20}
	fmt.Println(person.name)
	fmt.Println(person)
}
