package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image/color"
	"log"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("keyboard checker"),
			app.Size(unit.Dp(300), unit.Dp(350)),
		)

		if err := draw(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	margins := layout.Inset{
		Left:   unit.Dp(10),
		Right:  unit.Dp(10),
		Top:    unit.Dp(35),
		Bottom: unit.Dp(10),
	}

	var keyDesc = "press any key"
	var statusDesc = "press any key"

	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			title := material.H5(th, "keyboard checker")
			maroon := color.NRGBA{R: 127, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			//fmt.Println("FrameEvent")

			key.InputOp{
				Tag: w, // Use the window as the event routing tag. This means we can call gtx.Events(w) and get these events.
			}.Add(gtx.Ops)
			// Request keyboard focus to the current clip area.
			key.FocusOp{
				Tag: w, // Focus the input area with our window as the tag.
			}.Add(gtx.Ops)

			for _, event := range gtx.Events(w) {
				switch event := event.(type) {
				case key.Event:
					fmt.Println(event)
					keyDesc = event.Name

					if event.State == key.Press {
						statusDesc = "Press"
					}

					if event.State == key.Release {
						statusDesc = "Release"
					}

					log.Println(keyDesc, statusDesc)
					w.Invalidate()
				}
			}

			margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				labelKey := material.Label(th, unit.Sp(20), keyDesc)
				labelKey.Color = color.NRGBA{B: 255, A: 128}
				labelKey.Alignment = text.Start
				return labelKey.Layout(gtx)
			})

			margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				labelStatus := material.Label(th, unit.Sp(20), statusDesc)
				labelStatus.Color = color.NRGBA{B: 255, A: 128}
				labelStatus.Alignment = text.End
				return labelStatus.Layout(gtx)
			})

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}

func init() {
	log.SetFlags(log.LstdFlags)
}
