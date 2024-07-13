package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Size struct {
	heighth float64
	width   float64
}

type Item struct {
	Name string
	Size Size
}
type Offset struct {
	value float64
}

type Textout struct {
	sout *canvas.Text
	pout *canvas.Text
	iout *canvas.Text
	oout *canvas.Text
}

var a fyne.App
var w fyne.Window
var textout Textout
var fsize, psize, isize Item


func main() {

	fsize = Item{Name: "Frame"}
	psize = Item{Name: "Paper"}
	isize = Item{Name: "Image"}

	sizes := []string{"18 x 24", "13 x 19", "11 x 17", "11 x 14", "10 x 12",
		"8.5 x 11", "8 x 10", "Custom"}

	flabel := widget.NewLabel(fsize.Name)
	plabel := widget.NewLabel(psize.Name)
	ilabel := widget.NewLabel(isize.Name)

	a = app.New()

	w = a.NewWindow("Mat Calclulator")
	w.Resize(fyne.NewSize(800, 400))

	frame := container.New(layout.NewVBoxLayout(), flabel, widget.NewSelect(sizes, fsize.parseInput))
	paper := container.New(layout.NewVBoxLayout(), plabel, widget.NewSelect(sizes, psize.parseInput))
	image := container.New(layout.NewVBoxLayout(), ilabel, widget.NewSelect(sizes, isize.parseInput))

	getSizes := container.New(layout.NewHBoxLayout())
	getSizes.Add(frame)
	getSizes.Add(paper)
	getSizes.Add(image)

	offsetSelect := container.New(layout.NewHBoxLayout())

	var offset Offset
	oLabel := widget.NewLabel("Offset (negative is whitespace)")
	oBox := container.New(layout.NewVBoxLayout(), oLabel, widget.NewSelect([]string{"0.5", "0.25", "0", "-0.25", "-.5"}, offset.parseOffset))

	offsetSelect.Add(oBox)

	var dOffset Offset
	dBox := container.New(layout.NewVBoxLayout(), widget.NewLabel("Double Mat Offset"), widget.NewSelect([]string{"0.5", "0.25"}, dOffset.parseOffset))

	var dMat bool
	double := widget.NewCheck("Double Mat", func(value bool) {
		dMat = value
		if value {
			dBox.Show()
		}
		if !value {
			dBox.Hide()
		}
	})

	calc := widget.NewButton("Calculate", func() {

		textout.calculate(fsize, psize, isize, offset, dOffset, dMat)
	})

	calc.Importance = widget.MediumImportance

	doubleSelect := container.New(layout.NewHBoxLayout())

	doubleSelect.Add(double)
	doubleSelect.Add(dBox)
	dBox.Hide()
	doubleSelect.Add(layout.NewSpacer())
	doubleSelect.Add(calc)
	doubleSelect.Add(layout.NewSpacer())

	textout.pout = canvas.NewText("", color.NRGBA{R: 0, G: 180, B: 180, A: 255})
	textout.pout.TextSize = 18
	textout.iout = canvas.NewText("", color.NRGBA{R: 255, G: 60, B: 0, A: 255})
	textout.iout.TextSize = 18
	textout.oout = canvas.NewText("", color.NRGBA{R: 180, G: 180, B: 0, A: 255})
	textout.oout.TextSize = 18

	outputFields := container.New(layout.NewVBoxLayout())
	
	outputFields.Add(textout.pout)
	outputFields.Add(textout.iout)
	outputFields.Add(textout.oout)

	textout.sout = canvas.NewText("", color.NRGBA{R: 0, G: 180, B: 0, A: 255})
	status := container.New(layout.NewVBoxLayout())
	status.Add(textout.sout)
	spacer := layout.NewSpacer()

	textout.sout.TextSize = 18

	c := container.New(layout.NewVBoxLayout(), getSizes, offsetSelect, doubleSelect, outputFields, spacer, status)
	textout.sout.Text = fmt.Sprintf("Frame Size: %.3f x %.3f   Paper Size: %.3f x %.3f   Image Size: %.3f x %.3f",
		fsize.Size.heighth, fsize.Size.width,
		psize.Size.heighth, psize.Size.width,
		isize.Size.heighth, isize.Size.width)
	w.SetContent(c)

	w.ShowAndRun()

}


