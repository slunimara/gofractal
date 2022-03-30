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
		lastPixel: PointZero,
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
	c.lastPixel.AddXY(1, 0)

	if c.lastPixel.X() == c.Width() {
		c.lastPixel.SetX(1)
		c.lastPixel.AddXY(0, 1)
	}

	if c.lastPixel.Y() == c.Height() {
		c.SetLastPixel(PointZero)
	}
}

// SetLastPixel sets the current pixel position.
func (c *Canvas) SetLastPixel(p Point) {
	c.lastPixel = p
}

// DrawNextPixel draws the next pixel in the current direction.
func (c *Canvas) DrawNextPixel(color color.Color) {
	c.DrawPixel(color)
	c.NextPixel()
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
