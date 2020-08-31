package main

import (
	"fmt"
	"os"
)

const appName = "tomcat"
const AppName = "tomcat"

func main() {
	const (
		name = "tom"
		age  = 18
	)

	var (
		k = 15
		m = 20
	)

	fmt.Println(name, age)
	fmt.Printf("%T,%T\n", name, age)
	fmt.Printf("%v,%v\n", name, age)
	fmt.Println(k, m)
	fmt.Println(appName, AppName)

	fmt.Println(os.Args)

	var str0, str1, str2, str3 string
	fmt.Println(str0, str1, str2, str3)

	f, err := os.Open("tom")
	if err != nil {
		fmt.Printf("error :%s", err)
	}

	fmt.Println(f)

	pname := "tomcatpname"
	fmt.Println(pname)

	i, j := 15, 16
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)

	tom := "lucy"
	fmt.Println(tom)

	fmt.Println(&tom)
	p := &tom

	*p = "lucy" + "1"
	fmt.Println(*p)
	fmt.Println(tom)
	fmt.Println(tom == *p)

	q := new(int)
	*q = 15
	fmt.Println(*q)


}
