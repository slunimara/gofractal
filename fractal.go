package gofractal

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

// TODO: Documenation
func isStable(c complex128, z complex128, maxIterations uint) (bool, uint) {
	i := uint(0)

	for i <= maxIterations && complexAbs(z) <= 2 {
		z = complexPow2(z) + c
		i += 1
	}

	return complexAbs(z) <= 2, i
}

// TODO: Documenation
func Mandelbrot(canvas *Canvas, view *View, maxIterations int) {
	fmt.Print("orig wid: ", canvas.Width(), " orig hei: ", canvas.Height(), "\n")
	density := canvasDensity(canvas, view)

	xRange := arange(real(view.bottomLeft), real(view.topRight), density)
	yRange := arange(imag(view.topRight), imag(view.bottomLeft), density)

	fmt.Print("density: ", density, "\n")
	// fmt.Print("xRange: ", len(xRange), " yRange: ", len(yRange), "\n")
	fmt.Print("bl: ", view.bottomLeft, " tr: ", view.topRight, "\n")
	fmt.Print("wid: ", canvas.Width(), " hei: ", canvas.Height(), "\n")

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
				stable, _ := isStable(c, complex(0, 0), uint(maxIterations))
				//stable, _ := isStable(0.25+0i, c, uint(maxIterations))
				stableArray[x] = stable
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
	N := int(IntervalDistribution(start, stop, step))
	arr := make([]float64, N)

	if start > stop {
		step *= -1
	}

	for i := range arr {
		arr[i] = start + (float64(i) * step)
	}

	return arr
}

func canvasDensity(canvas *Canvas, view *View) float64 {
	width, height := canvas.Width(), canvas.Height()
	keepResolution := true

	if keepResolution {
		tr, bl := view.topRight, view.bottomLeft
		w, h := canvas.resolutionRatio()
		x, y := view.viewRatio()

		if x >= y {
			newRatioY := CrossMultiplication(float64(w), float64(h), x)
			ratioDifference := newRatioY - y
			sideExtension := (ratioDifference / 2) * y
			tr += complex(0, sideExtension)
			bl -= complex(0, sideExtension)

			fmt.Print("x: ", x, " y: ", y, " w: ", w, " h: ", h, "\n")
			fmt.Print("newRatioY: ", newRatioY, "\n")
			fmt.Print("difference: ", ratioDifference, "\n")
			fmt.Print("sideExtension: ", sideExtension, "\n\n")

			*view = *NewView(tr, bl)
			density := view.XDistance() / float64(width)

			return density
		} else {
			newRatioX := CrossMultiplication(float64(h), float64(w), y)
			ratioDifference := newRatioX - x
			sideExtension := (ratioDifference / 2) * x
			tr += complex(sideExtension, 0)
			bl -= complex(sideExtension, 0)

			fmt.Print("x: ", x, " y: ", y, " w: ", w, " h: ", h, "\n")
			fmt.Print("newRatioY: ", newRatioX, "\n")
			fmt.Print("difference: ", ratioDifference, "\n")
			fmt.Print("sideExtension: ", sideExtension, "\n\n")

			*view = *NewView(tr, bl)
			density := view.YDistance() / float64(height)

			return density
		}
	} else {
		if width >= height {
			density := view.XDistance() / float64(width)
			y1, y2 := imag(view.topRight), imag(view.bottomLeft)
			height := uint64(IntervalDistribution(y1, y2, density))

			*canvas = *NewCanvas(width, height)

			return density
		} else {
			density := view.YDistance() / float64(height)
			x1, x2 := real(view.topRight), real(view.bottomLeft)
			width := uint64(IntervalDistribution(x1, x2, density))

			*canvas = *NewCanvas(width, height)

			return density
		}
	}
}

// Determine the value of a variable.
// a : b = c : return
func CrossMultiplication(a, b, c float64) float64 {
	return b * c / a
}

func IntervalDistribution(start, stop, step float64) float64 {
	return math.Ceil((math.Abs(stop - start)) / step)
}

func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
