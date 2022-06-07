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
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

var boilDuration float32

var boilDurationInput widget.Editor

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

					inputString := boilDurationInput.Text()
					inputString = strings.TrimSpace(inputString)
					inputFloat, _ := strconv.ParseFloat(inputString, 32)
					boilDuration = float32(inputFloat)
					boilDuration = boilDuration / (1 - progress)

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
						d := image.Point{X: 400, Y: 400}
						return layout.Dimensions{Size: d}
					}),

					layout.Rigid(func(gtx C) D {
						bar := material.ProgressBar(th, progress)
						//fmt.Println("render", progress, boiling)
						return bar.Layout(gtx)
					}),

					// input
					layout.Rigid(func(gtx C) D {
						ed := material.Editor(th, &boilDurationInput, "sec")
						boilDurationInput.SingleLine = true
						boilDurationInput.Alignment = text.Middle

						if boiling && progress < 1 {
							boilRemain := (1 - progress) * boilDuration
							inputStr := fmt.Sprintf("%.1f", math.Round(float64(boilRemain)*10/10))
							boilDurationInput.SetText(inputStr)
						}

						margins := layout.Inset{
							Top:    unit.Dp(0),
							Right:  unit.Dp(170),
							Bottom: unit.Dp(40),
							Left:   unit.Dp(170),
						}
						border := widget.Border{
							Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
							CornerRadius: unit.Dp(3),
							Width:        unit.Dp(2),
						}

						return margins.Layout(gtx, func(gtx C) D {
							return border.Layout(gtx, ed.Layout)
						})
					}),

					// button
					layout.Rigid(func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(25),
							Left:   unit.Dp(25),
						}

						var btnText string

						if !boiling {
							btnText = "start"
						} else {
							btnText = "stop"
						}

						return margins.Layout(gtx, func(gtx C) D {
							btn := material.Button(th, &starButton, btnText)
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
			if boiling && progress < 1 {
				progress += p
				//fmt.Println("progress", progress)
				w.Invalidate()
			}

		}
	}

	return nil
}
