package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "newLine")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("newLine printed")
	}
}