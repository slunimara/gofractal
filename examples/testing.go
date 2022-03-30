package main

import (
	"fmt"
	"time"

	"github.com/slunimara/gofractal"
)

func main() {
	fmt.Println("--- Mandelbrot Set ---")

	canvas := gofractal.NewCanvas(25000, 20000)

	tStart := time.Now()
	gofractal.Mandelbrot(canvas, 1000, 0.0001)
	tEnd := time.Now()

	fmt.Printf("Time: %v\n", tEnd.Sub(tStart))

	canvas.Save("mandelbrot.png")

	fmt.Println("--- Done ---")
}
