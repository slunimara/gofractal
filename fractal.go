package gofractal

import (
	"image/color"
	"math"
	"math/cmplx"
)

func is_stable(c complex128, max_iterations int) bool {
	z := complex(0, 0)
	i := 1
	for i <= max_iterations {
		z = cmplx.Pow(z, 2) + c
		i += 1
	}
	return cmplx.Abs(c) <= 2
}

func Mandelbrot() {
	var width, height uint64 = 250, 200
	canvas := NewCanvas(width, height)
	view := NewView(complex(0.5, 1), complex(-2, -1))

	max_iterations := 10
	density := 0.01

	xRange := arange(real(view.bl), real(view.tr), density)
	yRange := arange(imag(view.tr), imag(view.bl), density)

	for _, im := range yRange {
		for _, re := range xRange {
			c := complex(re, im)

			if is_stable(c, max_iterations) {
				canvas.DrawNextPixel(color.Black)
			} else {
				canvas.DrawNextPixel(color.White)
			}
		}
	}

	canvas.Save("mandel.png")
}

// Neefektivní! Step musí být > 0 !!
func arange(start, stop, step float64) []float64 {
	N := int(math.Ceil((math.Abs(stop-start))/step)) + 1
	arr := make([]float64, N)

	if start > stop {
		step *= -1
	}

	for i, _ := range arr {
		arr[i] = start + (float64(i) * step)
	}

	return arr
}
