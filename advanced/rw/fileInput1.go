package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:\\IntelliJ_Code\\daojia\\o2o-group-shopping\\o2o-group-shopping-open-service\\src\\main\\assemble\\package.xml")

	if err != nil {
		fmt.Println("an error occurred on open file")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		readString, err := reader.ReadString('\n')
		fmt.Printf("%s\n", readString)

		if err == io.EOF {
			return
		}
	}
}
