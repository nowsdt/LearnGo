package main

import (
	"log"
	"os"
)

var cwd string

func main() {
	log.SetFlags(log.Llongfile)

	s := make([]byte, 5)
	log.Print(len(s), cap(s))
	s = s[2:4]
	log.Print(len(s), cap(s))

	log.Print(len(s[2:2]), len(s[2:3]))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	log.Print(string(s2))
	s2[1] = 't'
	log.Print(string(s2))
	log.Print(len(s), cap(s))

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
