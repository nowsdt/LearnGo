package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	port := os.Args[1]
	ln, err := net.Listen("tcp", ":8181")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn, port)
	}
}
func handleRequest(conn net.Conn, port string) {
	fmt.Println("new client")

	proxy, err := net.Dial("tcp", "127.0.0.1:"+port)
	if err != nil {
		panic(err)
	}

	fmt.Println("proxy connected")
	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
