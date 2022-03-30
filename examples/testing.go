package main

import (
	"image/color"

	"github.com/slunimara/gofractal"
)

func main() {
	canvas := gofractal.NewCanvas(25, 25)

	canvas.DrawNextPixel(color.RGBAModel.Convert(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
	canvas.DrawNextPixel(color.RGBAModel.Convert(color.RGBA{R: 0, G: 255, B: 0, A: 255}))
	canvas.DrawNextPixel(color.RGBAModel.Convert(color.RGBA{R: 0, G: 0, B: 255, A: 255}))

	canvas.Save("testing.png")
}
