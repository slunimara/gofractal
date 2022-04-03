package gofractal

import (
	"image/color"

	"github.com/fogleman/gg"
)

type Canvas struct {
	ctx          *gg.Context
	currentPixel Pixel
}

// NewCanvas returns a new canvas with the given width and height.
func NewCanvas(width, height uint64) *Canvas {
	return &Canvas{
		ctx:          gg.NewContext(int(width), int(height)),
		currentPixel: PixelZero(),
	}
}

// Width returns the width of the canvas.
func (c Canvas) Width() uint64 {
	return uint64(c.ctx.Width())
}

// Height returns the height of the canvas.
func (c Canvas) Height() uint64 {
	return uint64(c.ctx.Height())
}

func (c Canvas) ResolutionRatio() (uint64, uint64) {
	d := GCD(c.Height(), c.Width())

	return uint64(c.Width() / d), uint64(c.Height() / d)
}

// NextPixel increments the current pixel position.
func (c *Canvas) NextPixel() {
	c.currentPixel.AddX(1)

	if c.currentPixel.X() == c.Width() {
		c.currentPixel.SetX(0)
		c.currentPixel.AddY(1)
	}

	if c.currentPixel.Y() == c.Height() {
		c.SetCurrentPixel(PixelZero())
	}
}

// SetCurrentPixel sets the current pixel position.
func (c *Canvas) SetCurrentPixel(p Pixel) {
	c.currentPixel = p
}

func (c *Canvas) GetCurrentPixel() *Pixel {
	return &c.currentPixel
}

// DrawNextPixel draws the next pixel in the current direction.
func (c *Canvas) DrawNextPixel(color color.Color) {
	c.DrawPixel(color)
	c.NextPixel()
}

// DrawPixel draws a pixel at the current position.
func (c *Canvas) DrawPixel(color color.Color) {
	c.DrawPixelAt(
		c.currentPixel.X(),
		c.currentPixel.Y(),
		color)
}

// DrawPixelAt draws a pixel at the given point.
func (c *Canvas) DrawPixelAt(x, y uint64, color color.Color) {
	c.ctx.SetColor(color)
	c.ctx.SetPixel(int(x), int(y))
}

// Clear clears the canvas.
func (c *Canvas) Clear() {
	c.ctx.Clear()
}

// Save saves the canvas to the given file.
func (c Canvas) Save(filename string) error {
	return c.ctx.SavePNG(filename)
}
