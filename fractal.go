package gofractal

import (
	"image/color"
	"math"
	"math/cmplx"
)

// TODO: Documenation
func is_stable(c complex128, maxIterations int) bool {
	i := 0
	z := 0 + 0i

	for i < maxIterations {
		z = cmplx.Pow(z, 2) + c
		i++
	}

	return cmplx.Abs(z) <= 2
}

// TODO: Documenation
func Mandelbrot(canvas *Canvas, maxIterations int, density float64) {
	view := NewView(
		complex(0.5, 1),
		complex(-2, -1))

	xRange := arange(real(view.bl), real(view.tr), density)
	yRange := arange(imag(view.tr), imag(view.bl), density)

	for _, im := range yRange {
		for _, re := range xRange {
			c := complex(re, im)

			if is_stable(c, maxIterations) {
				canvas.DrawNextPixel(color.Black)
			} else {
				canvas.DrawNextPixel(color.White)
			}
		}
	}
}

// TODO: Documenation
// Neefektivní! Step musí být > 0 !!
func arange(start, stop, step float64) []float64 {
	N := int(math.Ceil((math.Abs(stop - start)) / step))
	arr := make([]float64, N)

	if start > stop {
		step *= -1
	}

	for i := range arr {
		arr[i] = start + (float64(i) * step)
	}

	return arr
}
