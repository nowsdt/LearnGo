package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	w1 := myApp.NewWindow("Widget")
	w1.SetContent(widget.NewEntry())

	w1.ShowAndRun()
}
