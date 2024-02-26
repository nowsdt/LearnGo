package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	if len(args) > 2 {
		str := strings.TrimSpace(args[1])
		subLen, _ := strconv.Atoi(args[2])

		fmt.Println(str[:subLen])
		fmt.Println(str[subLen:])
	}

	environ := os.Environ()
	fmt.Println(environ)
}
