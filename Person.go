package main

import "fmt"

type Person struct {
	name, address string
	age           int
}

func main() {
	person := &Person{"tom", "beijing", 20}
	fmt.Println(person.name)
	fmt.Println(person)

	const table = "TRUNCATE TABLE user_score_history"

	for i := 0; i < 50; i++ {
		fmt.Printf("%s%v;\n", table, i)
	}

}
