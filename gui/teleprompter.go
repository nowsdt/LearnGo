package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var lines []string

func main() {

	f, err := ioutil.ReadFile("E:\\GoWS\\LearnGo\\gui\\speech.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(f), "\n")

	for i := 0; i < 10; i++ {
		lines = append(lines, "")
	}

	//fmt.Println(strings.Join(lines, ","))

	go func() {
		w := app.NewWindow(
			app.Title("Teleprompter"),
			app.Size(unit.Dp(350), unit.Dp(300)),
		)

		if err := draw(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func draw(w *app.Window) error {
	var scrollY int = 0
	var focusBarY int = 78
	var textWidth int = 300
	var fontSize int = 35
	var autoScroll bool = false
	var autoSpeed int = 1
	th := material.NewTheme(gofont.Collection())

	for e := range w.Events() {
		switch e := e.(type) {

		case system.FrameEvent:
			var ops op.Ops
			gtx := layout.NewContext(&ops, e)

			// https://git.sr.ht/~eliasnaur/gio-example/tree/b8e926ba1ef2/item/gophers/main.go
			key.InputOp{
				Tag: w,
				Keys: key.Set(strings.Join(
					[]string{key.NameSpace, key.NameShift, key.NameUpArrow, key.NameDownArrow, "J", "K", "F", "S", "N", "+", "-", "Shift-A"},
					"|")),
			}.Add(gtx.Ops)

			// https://git.sr.ht/~eliasnaur/gio-example/tree/b8e926ba1ef2/item/opengl/main.go
			pointer.InputOp{
				Tag:   w,
				Types: pointer.Scroll,
			}.Add(gtx.Ops)

			for _, event := range gtx.Events(w) {
				log.Println(event)
				switch event := event.(type) {
				case key.Event:
					log.Println(event.String())

					if event.State == key.Release {
						var stepSize int = 1
						if event.Modifiers == key.ModShift {
							stepSize = 10
						}

						if event.Name == key.NameDownArrow || event.Name == "J" {
							scrollY = scrollY + stepSize*4
						}

						if event.Name == key.NameUpArrow || event.Name == "K" {
							scrollY = scrollY - stepSize*4
							if scrollY < 0 {
								scrollY = 0
							}
						}

						if event.Name == key.NameSpace {
							autoScroll = !autoScroll
							if autoSpeed == 0 {
								autoScroll = true
								autoSpeed++
							}
						}

						if event.Name == "F" {
							autoScroll = true
							autoSpeed++
						}
						if event.Name == "S" {
							if autoSpeed > 0 {
								autoSpeed--
							}
						}

						if event.Name == "N" {
							textWidth = textWidth - stepSize*10
						}

						if event.Name == "+" {
							fontSize = fontSize + stepSize
						}

						if event.Name == "-" {
							fontSize = fontSize - stepSize
						}
						if event.Name == "U" {
							focusBarY = focusBarY - stepSize
						}
						if event.Name == "D" {
							focusBarY = focusBarY + stepSize
						}

						w.Invalidate()
					}
				case pointer.Event:
					log.Println(event.Type, event.Modifiers, event.Source, event.Scroll, event.PointerID, event.Position)

					if event.Type == pointer.Scroll {
						var stepSize int = 1

						if event.Modifiers == key.ModShift {
							stepSize = 3
						}

						thisScroll := int(event.Position.Y)
						log.Println(scrollY, thisScroll, stepSize)
						scrollY = scrollY + thisScroll*stepSize

						if scrollY < 0 {
							scrollY = 0
						}

						log.Println("scrollY:", scrollY)
						w.Invalidate()
					}
				default:
					log.Println("point.default", event)
				}
			}

			paint.Fill(&ops, color.NRGBA{R: 0xff, G: 0xfe, B: 0xe0, A: 0xff})

			if autoScroll {
				scrollY = scrollY + autoSpeed
				op.InvalidateOp{At: gtx.Now.Add(time.Second / 50)}.Add(&ops)
			}

			marginWidth := (gtx.Constraints.Max.X - textWidth) / 2
			margins := layout.Inset{
				Left:   unit.Dp(float32(marginWidth)),
				Right:  unit.Dp(float32(marginWidth)),
				Top:    unit.Dp(0),
				Bottom: unit.Dp(0),
			}

			visList := layout.List{
				Axis: layout.Vertical,
				Position: layout.Position{
					Offset: scrollY,
				},
			}

			margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

				return visList.Layout(gtx, len(lines),
					func(gtx layout.Context, index int) layout.Dimensions {
						paragraph := material.Label(th, unit.Sp(float32(fontSize)), lines[index])
						paragraph.Alignment = text.Middle
						//fmt.Println(index, lines[index])
						return paragraph.Layout(gtx)
					})
			})

			op.Offset(image.Pt(0, focusBarY)).Add(&ops)
			stack := clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, 50)}.Push(&ops)
			paint.ColorOp{Color: color.NRGBA{R: 0xff, A: 0x66}}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			stack.Pop()

			//fmt.Println(scrollY, focusBarY, textWidth, fontSize, autoScroll, autoSpeed)

			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}

func init() {
	log.SetFlags(log.Lshortfile)
}
