package main

import (
	"fmt"
	"time"

	"github.com/slunimara/gofractal"
)

func main() {
	fmt.Println("--- Mandelbrot Set ---")

	mandelbrot()

	fmt.Println("--- Done ---")
}

func mandelbrot() {
	canvas := gofractal.NewCanvas(1920, 1080)

	tStart := time.Now()
	view := gofractal.NewView(
		complex(0.5, 1),
		complex(-2, -1))

	// view := gofractal.NewView(
	// 	complex(0, 1),
	// 	complex(-2, 0))

	gofractal.Mandelbrot(canvas, view, 100)
	tEnd := time.Now()

	fmt.Printf("Time: %v\n", tEnd.Sub(tStart))
	canvas.Save("mandelbrot.png")
}

func performanceTesting() {
	fmt.Println("--- Time testing")

	// c := 528649457 + 38573538347i

	// tStart := time.Now()
	// cmplx.Abs(c)
	// tEnd := time.Now()
	// fmt.Printf("Time: %v\n", tEnd.Sub(tStart))

	// tStart = time.Now()
	// gofractal.CAbs(c)
	// tEnd = time.Now()
	// fmt.Printf("Time: %v\n", tEnd.Sub(tStart))
}
