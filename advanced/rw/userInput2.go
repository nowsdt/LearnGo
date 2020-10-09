package main

import (
	"bufio"
	"fmt"
	"os"
)

// 读取用户输入

func main() {
	var inputReader *bufio.Reader
	var err error
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("enter some words")
	input, err := inputReader.ReadString('\n')

	if err == nil {
		fmt.Printf("input:%s\n", input)
	}
}
