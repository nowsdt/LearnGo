package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		panic("error" + err.Error())
	}

	buf := make([]byte, 512)
	go processData(conn, buf)

	for {
		reader := bufio.NewReader(os.Stdin)
		readString, _ := reader.ReadString('\n')
		_, _ = conn.Write([]byte(readString))
	}
}

func processData(conn net.Conn, buf []byte) {
	for {
		n, err := conn.Read(buf)
		if err != io.EOF && n > 0 {
			fmt.Println("receive from server:", string(buf[:n]))
		}
	}
}
