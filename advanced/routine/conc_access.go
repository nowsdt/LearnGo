package main

import (
	"fmt"
	"strconv"
)

type person struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *person {
	p := &person{name, salary, make(chan func())}
	go p.backEnd()
	return p
}

func (p *person) backEnd() {
	for f := range (*p).chF {
		f()
	}
}

func (p *person) SetSalary(sal float64) {
	p.chF <- func() {
		p.salary = sal
	}
}

func (p *person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() {
		fChan <- p.salary
	}
	return <-fChan
}

func (p *person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)
}
