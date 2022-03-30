package gofractal

import (
	"image/color"

	"github.com/fogleman/gg"
)

type Canvas struct {
	ctx       *gg.Context
	lastPixel Point
}

// NewCanvas returns a new canvas with the given width and height.
func NewCanvas(width, height uint64) *Canvas {
	return &Canvas{
		ctx:       gg.NewContext(int(width), int(height)),
		lastPixel: NewPoint(1, 1),
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

// NextPixel increments the current pixel position.
func (c *Canvas) NextPixel() {
	c.lastPixel.AddX(1)

	if c.lastPixel.X() == c.Width()+1 {
		c.lastPixel.SetX(1)
		c.lastPixel.AddY(1)
	}

	if c.lastPixel.Y() == c.Height()+1 {
		c.lastPixel = NewPoint(1, 1)
	}
}

// SetLastPixel sets the current pixel position.
func (c *Canvas) SetLastPixel(p Point) {
	c.lastPixel = p
}

// DrawNextPixel draws the next pixel in the current direction.
func (c *Canvas) DrawNextPixel(color color.Color) {
	c.NextPixel()
	c.DrawPixel(color)
}

// DrawPixel draws a pixel at the current position.
func (c *Canvas) DrawPixel(color color.Color) {
	c.DrawPixelAt(
		c.lastPixel.X(),
		c.lastPixel.Y(),
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
func (c *Canvas) Save(filename string) error {
	return c.ctx.SavePNG(filename)
}
