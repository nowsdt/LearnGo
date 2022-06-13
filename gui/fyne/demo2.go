package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {

	myApp := app.New()

	w := myApp.NewWindow("hello")

	clock := widget.NewLabel("")

	updateTime(clock)

	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

func updateTime(clock *widget.Label) {
	format := time.Now().Format("Time: 03:04:05")
	clock.SetText(format)
}
