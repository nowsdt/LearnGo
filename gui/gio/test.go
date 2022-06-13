package main

// A simple Gio program. See https://gioui.org for more information.

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
)

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			// Gather and print all events captured by our input area since the previous frame.
			for _, event := range gtx.Events(w) {
				// Perform event handling here instead of in the outer type switch.
				log.Printf("%#+v", event)
			}

			// Create a clip area the size of the window.
			areaStack := clip.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Push(gtx.Ops)
			// Register for all pointer inputs on the current clip area.
			pointer.InputOp{
				Types: pointer.Enter | pointer.Leave | pointer.Drag | pointer.Press | pointer.Release | pointer.Scroll | pointer.Move,
				Tag:   w, // Use the window as the event routing tag. This means we can call gtx.Events(w) and get these events.
			}.Add(gtx.Ops)
			// Register for keyboard input on the current clip area.
			key.InputOp{
				Tag: w, // Use the window as the event routing tag. This means we can call gtx.Events(w) and get these events.
			}.Add(gtx.Ops)
			// Request keyboard focus to the current clip area.
			key.FocusOp{
				Tag: w, // Focus the input area with our window as the tag.
			}.Add(gtx.Ops)
			// Pop the clip area to finalize it.
			areaStack.Pop()

			e.Frame(gtx.Ops)
		}
	}
}
