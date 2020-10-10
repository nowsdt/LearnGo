package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8000")
	checkError(err)
	fmt.Println("server started")
	for {
		conn, err := server.Accept()
		checkError(err)
		fmt.Printf("client:%s connected\n", conn.RemoteAddr().String())
		go processRequest(conn)
	}
}

func processRequest(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		checkError(err)
		fmt.Println("client:", conn.RemoteAddr(), " receive from client:", strings.Trim(string(buf[:n]), "\r\n"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error ", err.Error())
		panic(err.Error())
	}
}
