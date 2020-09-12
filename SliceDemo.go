package main

import "fmt"

func main() {
	var str = "hello"
	p1, p2 := splitStr(str, 2)

	fmt.Println(p1, p2)

	reverse(&str)
	fmt.Println("reverse1:", str)

	s := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(s)
}

func splitStr(s string, i int) (p1, p2 string) {
	bytes := []byte(s)

	p1 = string(bytes[:i])
	p2 = string(bytes[i:])
	return
}

func reverse(s *string) {
	bytes := []byte(*s)
	for i := 0; i <= len(bytes)/2; i++ {
		p := bytes[i]
		idx := len(bytes) - 1 - i
		fmt.Println(i, idx)
		bytes[i] = bytes[idx]
		bytes[idx] = p
	}
	fmt.Println("reverse0:", string(bytes))

}
