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
	canvas := gofractal.NewCanvas(2500, 2000)

	tStart := time.Now()
	gofractal.Mandelbrot(canvas, 1000, 0.001)
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
