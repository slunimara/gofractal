package gofractal

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

const (
	max_goroutines = 64
	debug          = true
)

type FractalConfig interface {
	IsStable(c, z complex128) (bool, uint)
	MaxIterations() uint
	Draw(canvas *Canvas)
	View() *View
}

// TODO: Documenation
func Fractal(canvas *Canvas, config FractalConfig) {
	if debug {
		fmt.Println("orig wid: ", canvas.Width(), " orig hei: ", canvas.Height())
	}

	var (
		mutex     sync.Mutex
		waitGroup = NewWaitGroup()

		bottomLeft = config.View().bottomLeft
		topRight   = config.View().topRight

		density = canvasDensity(canvas, config.View())

		xRange = arange(real(bottomLeft), real(topRight), density)
		yRange = arange(imag(topRight), imag(bottomLeft), density)
	)

	if debug {
		fmt.Println("density: ", density)
		fmt.Println("bl: ", bottomLeft, " tr: ", topRight)
		fmt.Println("wid: ", canvas.Width(), " hei: ", canvas.Height())
	}

	for y, _imag := range yRange {

		for waitGroup.Length() >= max_goroutines {
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

// TODO: Documenation
func canvasDensity(canvas *Canvas, view *View) float64 {
	keepRes := true

	if keepRes {
		return keepResolution(canvas, view)
	} else {
		return keepView(canvas, view)
	}
}

// TODO: Documenation
func keepResolution(canvas *Canvas, view *View) float64 {
	var density float64
	tr, bl := view.topRight, view.bottomLeft
	w, h := canvas.ResolutionRatio()
	x, y := view.ViewRatio()

	if x >= y {
		newRatioY := CrossMultiplication(float64(w), float64(h), x)
		ratioDifference := newRatioY - y
		sideExtension := (ratioDifference / 2) * y
		tr += complex(0, sideExtension)
		bl -= complex(0, sideExtension)

		if debug {
			fmt.Println("x: ", x, " y: ", y, " w: ", w, " h: ", h)
			fmt.Println("newRatioY: ", newRatioY)
			fmt.Println("difference: ", ratioDifference)
			fmt.Println("sideExtension: ", sideExtension)
		}

		*view = *NewView(tr, bl)
		density = view.XDistance() / float64(canvas.Width())
	} else {
		newRatioX := CrossMultiplication(float64(h), float64(w), y)
		ratioDifference := newRatioX - x
		sideExtension := (ratioDifference / 2) * x
		tr += complex(sideExtension, 0)
		bl -= complex(sideExtension, 0)

		if debug {
			fmt.Println("x: ", x, " y: ", y, " w: ", w, " h: ", h)
			fmt.Println("newRatioY: ", newRatioX)
			fmt.Println("difference: ", ratioDifference)
			fmt.Println("sideExtension: ", sideExtension)
		}

		*view = *NewView(tr, bl)
		density = view.YDistance() / float64(canvas.Height())
	}

	return density
}

// TODO: Documenation
func keepView(canvas *Canvas, view *View) float64 {
	width, height := canvas.Width(), canvas.Height()

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

// Determine the value of a variable.
// a : b = c : return
func CrossMultiplication(a, b, c float64) float64 {
	return b * c / a
}

// TODO: Documenation
func IntervalDistribution(start, stop, step float64) float64 {
	return math.Ceil((math.Abs(stop - start)) / step)
}

// TODO: Documenation
func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
