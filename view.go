package gofractal

import "math"

var (
	Default = NewView(complex(0.5, 1), complex(-2, -1))
)

type View struct {
	topRight   complex128
	bottomLeft complex128
}

// TODO: Documenation
func NewView(topRight, bottomLeft complex128) *View {
	return &View{
		topRight:   topRight,
		bottomLeft: bottomLeft,
	}
}

// TODO: Documenation
func (v View) TopRight() complex128 {
	return v.topRight
}

// TODO: Documenation
func (v View) BottomLeft() complex128 {
	return v.bottomLeft
}

// TODO: Documenation
func (v View) ViewRatio() (float64, float64) {
	x, y := v.XDistance(), v.YDistance()

	return x, y
}

// TODO: Documenation
func (v View) XDistance() float64 {
	x1, x2 := real(v.topRight), real(v.bottomLeft)

	return math.Abs(x1 - x2)
}

// TODO: Documenation
func (v View) YDistance() float64 {
	y1, y2 := imag(v.topRight), imag(v.bottomLeft)

	return math.Abs(y1 - y2)
}
