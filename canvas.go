package gofractal

import (
	"github.com/fogleman/gg"
)

type Canvas struct {
	ctx *gg.Context
}

func NewCanvas(width, height uint64) *Canvas {
	return &Canvas{
		ctx: gg.NewContext(int(width), int(height)),
	}
}

type Color struct {
	r, g, b uint8
}

func NewColor(r, g, b uint8) Color {
	return Color{
		r: r,
		g: g,
		b: b,
	}
}

func (c Color) R() uint8 {
	return c.r
}

func (c Color) G() uint8 {
	return c.g
}

func (c Color) B() uint8 {
	return c.b
}

func (c Color) RGB() (uint8, uint8, uint8) {
	return c.r, c.g, c.b
}
