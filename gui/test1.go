package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"log"
	"os"
	"time"
)

type C = layout.Context
type D = layout.Dimensions

var ops op.Ops
var starButton widget.Clickable

var th = material.NewTheme(gofont.Collection())

var progressIncrmenter chan float32

var progress float32

var boiling bool

func main() {

	progressIncrmenter = make(chan float32)

	go func() {
		for {
			time.Sleep(time.Second / 25)
			progressIncrmenter <- 0.04
		}
	}()

	go func() {
		w := app.NewWindow(
			app.Title("Egg timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)

	}()

	app.Main()
}

func draw(w *app.Window) error {

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				if starButton.Clicked() {
					fmt.Println("click", progress)
					boiling = !boiling
				}

				layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						circle := clip.Ellipse{
							Min: image.Pt(80, 0),
							Max: image.Pt(320, 240),
						}.Op(gtx.Ops)

						nrgba := color.NRGBA{B: 255, A: 100}
						paint.FillShape(gtx.Ops, nrgba, circle)
						d := image.Point{Y: 400}
						return layout.Dimensions{Size: d}
					}),

					layout.Rigid(func(gtx C) D {
						bar := material.ProgressBar(th, progress)
						fmt.Println("render", progress, boiling)
						return bar.Layout(gtx)
					}),
					layout.Rigid(func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(25),
							Left:   unit.Dp(25),
						}

						var text string

						if !boiling {
							text = "start"
						} else {
							text = "stop"
						}

						return margins.Layout(gtx, func(gtx C) D {
							btn := material.Button(th, &starButton, text)
							return btn.Layout(gtx)
						})

					}),
					layout.Rigid(
						layout.Spacer{Height: unit.Dp(25)}.Layout),
				)

				e.Frame(gtx.Ops)

			case system.DestroyEvent:
				return e.Err
			}

		case p := <-progressIncrmenter:
			fmt.Println("progressIncrmenter", p)

			if boiling && progress < 1 {
				progress += p
				fmt.Println(progress)
				w.Invalidate()
			}

		}
	}

	return nil
}
