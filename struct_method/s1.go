package main

import (
	. "LearnGo/structpack"
	"fmt"
	"strings"
)

type BookAnonymous struct {
	book Book
	string
	int
}

func main() {
	//混合字面量语法（composite literal syntax）&struct1{a, b, c} 是一种简写，底层仍然会调用 new ()
	var bk1 Book
	bk1 = Book{Id: 2, Name: "lucy"}
	bk1 = Book{3, "lucy"}
	fmt.Println(bk1)

	bk2 := &Book{
		Id:   1,
		Name: "tom",
	}
	fmt.Println(bk2)
	fmt.Println(*bk2)

	bk3 := Book{1, "java se book"}
	fmt.Println(bk3)

	bk4 := new(Book)

	bk4.Name = "golang book"
	fmt.Println(bk4)
	fmt.Println(*bk4)
	toUpper(bk4)
	fmt.Println(bk4)

	bk5 := Book{12, "golang api"}
	toUpper(&bk5)
	fmt.Println(bk5)

	bka := &BookAnonymous{bk5, "author", 9}
	fmt.Println(*bka)
	fmt.Println(bka.string)

}

func toUpper(bk *Book) {
	(*bk).Name = strings.ToUpper(bk.Name)
}
