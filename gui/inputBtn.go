package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
	"image/color"
	"os"
	"time"
)

func main() {

	go func() {
		window := app.NewWindow(
			app.Title("Btn"),
			app.Size(unit.Dp(400), unit.Dp(800)),
		)
		var ops = new(op.Ops)

		changes := time.NewTicker(time.Second)
		defer changes.Stop()

		btnOffset := 0

		for {
			select {
			case e := <-window.Events():
				switch e := e.(type) {
				case system.DestroyEvent:
					fmt.Println(e.Err)
					os.Exit(0)
				case system.FrameEvent:

					//gtx := layout.NewContext(ops, e)

					ops.Reset()

					op.Offset(image.Pt(btnOffset, 50)).Add(ops)

					doButton(ops, e.Queue)

					e.Frame(ops)

				}
			case t := <-changes.C:
				btnOffset = int(t.Second() % 3 * 100)
				window.Invalidate()
			}
		}
	}()

	app.Main()
}

var tag = new(bool)
var pressed = false

func doButton(ops *op.Ops, q event.Queue) {
	for _, ev := range q.Events(tag) {
		if x, ok := ev.(pointer.Event); ok {
			switch x.Type {
			case pointer.Press:
				pressed = true
			case pointer.Release:
				pressed = false
			}
		}
	}

	area := clip.Rect(image.Rect(0, 0, 500, 500)).Push(ops)

	pointer.InputOp{
		Tag:   tag,
		Types: pointer.Press | pointer.Release,
	}.Add(ops)

	area.Pop()

	defer clip.Rect{Max: image.Pt(200, 200)}.Push(ops).Pop()

	var c color.NRGBA

	if pressed {
		c = color.NRGBA{R: 255, A: 255}
	} else {
		c = color.NRGBA{G: 255, A: 255}
	}

	//clip.Rect{Max: image.Pt(200, 200)}.Push(ops)
	paint.ColorOp{Color: c}.Add(ops)
	paint.PaintOp{}.Add(ops)

}
