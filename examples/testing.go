package main

import (
	"fmt"

	"github.com/slunimara/gofractal"
)

func main() {
	fmt.Print("--- Mandelbrot Set ---\n")
	gofractal.Mandelbrot()

	fmt.Print("\n--- Done ---")
}
