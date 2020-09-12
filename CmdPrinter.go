package main

import (
	"flag"
	"fmt"
)

type People struct {
	name string
	age  int8
}

func (s *People) String() string {
	return fmt.Sprintf("name: %s, age: %d", s.name, s.age)
}

func (s *People) Set(name string) error {
	s.name = name
	s.age = 18
	return nil

}

func PeopleFlag(value *People, name, usage string) *People {
	flag.CommandLine.Var(value, name, usage)
	return nil
}

func main() {
	p := new(People)
	PeopleFlag(p, "name", "name usage")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(p)
}
