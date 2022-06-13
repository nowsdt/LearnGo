package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image/color"
	"time"
)

func main() {

	a := app.New()

	w1 := a.NewWindow("Canvas")
	myCanvas := w1.Canvas()
	//rectangle(myCanvas)
	circle(myCanvas)

	w1.Resize(fyne.NewSize(100, 100))
	w1.ShowAndRun()
}

func rectangle(myCanvas fyne.Canvas) {
	blue := color.NRGBA{B: 180, A: 255}
	rectangle := canvas.NewRectangle(blue)

	myCanvas.SetContent(rectangle)

	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		rectangle.FillColor = green
		rectangle.Refresh()
	}()
}

func circle(myCanvas fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}

	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 3
	circle.StrokeColor = red

	myCanvas.SetContent(circle)

	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		circle.FillColor = green
		circle.Refresh()
	}()
}
