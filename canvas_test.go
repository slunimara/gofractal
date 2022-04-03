package gofractal

import (
	"image/color"
	"testing"
)

// TODO: Write better test
func TestNextPixel(t *testing.T) {
}

func TestDrawing(t *testing.T) {
	canvas := NewCanvas(100, 100)

	for i := 0; i < 100; i++ {
		for y := 0; y < 100; y++ {
			if y%2 == 0 {
				canvas.DrawNextPixel(color.Black)
			} else {
				canvas.DrawNextPixel(color.White)
			}
		}
	}

	canvas.Save("test.png")
}
