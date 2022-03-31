package gofractal

import (
	"image/color"
	"math"
	"math/cmplx"
	"sync"
)

// TODO: Documenation
func is_stable(c complex128, maxIterations int) bool {
	i := 1
	z := 0 + 0i

	for i <= maxIterations {
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

	var mutex sync.Mutex
	waitGroup := NewWaitGroup()

	for y, im := range yRange {

		for waitGroup.Length() >= 64 {
			continue
		}

		waitGroup.Add(1)

		go func(y int, im float64) {
			defer waitGroup.Done()

			stableArray := make([]bool, len(xRange))

			for x, re := range xRange {
				c := complex(re, im)

				stableArray[x] = is_stable(c, maxIterations)
			}

			mutex.Lock()
			for x, stable := range stableArray {

				if stable {
					canvas.DrawPixelAt(uint64(x), uint64(y), color.Black)
				} else {
					canvas.DrawPixelAt(uint64(x), uint64(y), color.White)
				}

			}
			mutex.Unlock()

		}(y, im)
	}

	waitGroup.Wait()
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
