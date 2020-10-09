package main

import "fmt"

// 读取用户输入

func main() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "6.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)
	fmt.Println(" enter your full name")
	fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s", &firstName, &lastName)

	fmt.Printf("hi %s %s !\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("from the string we read:", f, i, s)
}
