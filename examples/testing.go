package main

import (
	"fmt"

	"github.com/slunimara/gofractal"
)

func main() {
	fmt.Println("--- Mandelbrot Set ---")

	canvas := gofractal.NewCanvas(250, 200)
	gofractal.Mandelbrot(canvas, 10, 0.01)
	canvas.Save("mandelbrot.png")

	fmt.Println("--- Done ---")
}
