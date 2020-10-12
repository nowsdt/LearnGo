package main

import (
	"fmt"
	"os"
	"text/template"
)

type person struct {
	Name           string
	nonExportedAge string
}

func main() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{html .Name}} {{.Name}} {{.Name|html}} {{.}} !")
	p := person{"lucy>", "ages"}

	fmt.Println(p)

	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("there was an error:", err.Error())
	}
}
