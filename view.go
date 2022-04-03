package gofractal

import "math"

type View struct {
	topRight   complex128
	bottomLeft complex128
}

// Rectangle view of the complex plane.
func NewView(topRight, bottomLeft complex128) *View {
	return &View{
		topRight:   topRight,
		bottomLeft: bottomLeft,
	}
}

// Returns the top right coordinate.
func (v View) TopRight() complex128 {
	return v.topRight
}

// Returns the bottom left coordinate.
func (v View) BottomLeft() complex128 {
	return v.bottomLeft
}

// Returns ratio of the lengths of the x and y intervals.
func (v View) ViewRatio() (float64, float64) {
	x, y := v.XDistance(), v.YDistance()

	return x, y
}

// Returns the lengths of the x interval.
func (v View) XDistance() float64 {
	x1, x2 := real(v.topRight), real(v.bottomLeft)

	return math.Abs(x1 - x2)
}

// Returns the lengths of the y interval.
func (v View) YDistance() float64 {
	y1, y2 := imag(v.topRight), imag(v.bottomLeft)

	return math.Abs(y1 - y2)
}
