package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {

	a := app.New()

	w := a.NewWindow("Container")

	green := color.NRGBA{G: 180, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))

	content := container.NewWithoutLayout(text1, text2)

	w.SetContent(content)

	w.Resize(fyne.NewSize(200, 200))

	w.ShowAndRun()
}
