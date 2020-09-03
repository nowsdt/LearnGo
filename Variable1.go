package main

import (
	"log"
	"os"
)

var cwd string

func main() {
	log.Print(cwd)
	var err error
	cwd, err = os.Getwd()
	log.Print(cwd, err)

	const (
		a = 1
		b
		c = 2
		d
	)
	log.Println(a, b, c, d)

	const (
		e = iota + 100
		f
		g
		h
	)
	log.Println(e, f, g, h)
}

func init() {
	// 声明方法内局部变量，该声明未改变全局的cwd
	cwd, err := os.Getwd()

	/*	var err error
		cwd, err = os.Getwd()*/

	if err != nil {
		log.Fatalf("os.getWd failed: %v", err)
	}
	log.Printf("os.getWd :%s", cwd)

}
