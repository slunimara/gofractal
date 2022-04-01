package gofractal

import (
	"image/color"
	"math"
	"sync"
)

type FractalConfig interface {
	IsStable(c, z complex128) (bool, uint)
	MaxIterations() uint
	Density() float64
	Draw(canvas *Canvas)
}

func Fractal(canvas *Canvas, config FractalConfig) {
	view := NewView(
		complex(0.5, 1),
		complex(-2, -1))

	xRange := arange(real(view.bl), real(view.tr), config.Density())
	yRange := arange(imag(view.tr), imag(view.bl), config.Density())

	var mutex sync.Mutex
	waitGroup := NewWaitGroup()

	for y, _imag := range yRange {

		for waitGroup.Length() >= 64 {
			continue
		}

		waitGroup.Add(1)

		go fractalLineComputation(waitGroup, xRange, _imag, config, &mutex, canvas, y)
	}

	waitGroup.Wait()
}

func fractalLineComputation(
	waitGroup *WaitGroup,
	xRange []float64,
	_imag float64,
	config FractalConfig,
	mutex *sync.Mutex,
	canvas *Canvas,
	y int,
) {
	defer waitGroup.Done()

	stableArray := make([]bool, len(xRange))

	for x, _real := range xRange {
		c := complex(_real, _imag)

		stable, _ := config.IsStable(c, complex(0, 0))

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
