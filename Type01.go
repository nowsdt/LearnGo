package main

import (
	"log"
	"math"
)

func main() {
	i := 16
	j := -12
	log.Printf("%v", i%j)

	var x uint8 = 1
	var y uint8 = 14

	log.Printf("X:%08b\n", x)
	log.Printf("Y:%08b\n", y)

	log.Printf("X&^:%08b", x&^y)

	medals := []string{"gold", "silver", "bronze"}
	log.Println("==========")
	for k := len(medals) - 1; k >= 0; k-- {
		log.Println(medals[k])
	}
	log.Println(math.MaxInt16)

	var z complex64 = 10 + 15i
	log.Printf("z:%v %T", z, z)

	s := "helloworld"
	log.Printf("hello: %v %T", s[2], s)
	log.Printf("hello: %v %T", s[:2], s)

	const html = `go it tool for dev
					hello`

	log.Println(html)

}
