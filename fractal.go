package gofractal

import (
	"fmt"
	"image/color"
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

// TODO: Documenation
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

		stableArray[x], _ = config.IsStable(c, complex(0, 0))
	}

	mutex.Lock()
	defer mutex.Unlock()

	for x, isStable := range stableArray {

		x := uint64(x)
		y := uint64(y)

		if isStable {
			canvas.DrawPixelAt(x, y, color.Black)
		} else {
			canvas.DrawPixelAt(x, y, color.White)
		}
	}
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
	topRight := view.topRight
	bottomLeft := view.bottomLeft

	w, h := canvas.ResolutionRatio()
	x, y := view.ViewRatio()

	if x >= y {
		newRatioY := CrossMultiplication(float64(w), float64(h), x)
		ratioDifference := newRatioY - y
		sideExtension := (ratioDifference / 2) * y

		topRight += complex(0, sideExtension)
		bottomLeft -= complex(0, sideExtension)

		if debug {
			fmt.Println("x: ", x, " y: ", y, " w: ", w, " h: ", h)
			fmt.Println("newRatioY: ", newRatioY)
			fmt.Println("difference: ", ratioDifference)
			fmt.Println("sideExtension: ", sideExtension)
		}

		*view = *NewView(topRight, bottomLeft)
		return view.XDistance() / float64(canvas.Width())
	}

	newRatioX := CrossMultiplication(float64(h), float64(w), y)
	ratioDifference := newRatioX - x
	sideExtension := (ratioDifference / 2) * x

	topRight += complex(sideExtension, 0)
	bottomLeft -= complex(sideExtension, 0)

	if debug {
		fmt.Println("x: ", x, " y: ", y, " w: ", w, " h: ", h)
		fmt.Println("newRatioY: ", newRatioX)
		fmt.Println("difference: ", ratioDifference)
		fmt.Println("sideExtension: ", sideExtension)
	}

	*view = *NewView(topRight, bottomLeft)
	return view.YDistance() / float64(canvas.Height())
}

// TODO: Documenation
func keepView(canvas *Canvas, view *View) float64 {
	width := canvas.Width()
	height := canvas.Height()

	if width >= height {
		density := view.XDistance() / float64(width)
		y1 := imag(view.topRight)
		y2 := imag(view.bottomLeft)
		height := uint64(IntervalDistribution(y1, y2, density))

		*canvas = *NewCanvas(width, height)

		return density
	}

	density := view.YDistance() / float64(height)
	x1 := real(view.topRight)
	x2 := real(view.bottomLeft)
	width = uint64(IntervalDistribution(x1, x2, density))

	*canvas = *NewCanvas(width, height)

	return density
}

