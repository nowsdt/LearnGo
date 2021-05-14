package main

import (
	"fmt"
	"os"
)

var (
	sharding = 64
)

func main() {
	if len(os.Args) < 1 {
		return
	}
	str := os.Args[1]

	hash := strHash(str)

	fmt.Println(hash % int32(sharding))
}

func strHash(str string) int32 {
	hash := int32(0)
	for _, c := range []rune(str) {
		hash = 31*hash + c
	}
	return hash
}
