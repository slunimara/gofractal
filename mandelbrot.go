package gofractal

type mandelbrot struct {
	maxIterations uint
	density       float64
}

// TODO: Documenation
func NewMandelbrot(maxIterations uint, density float64) *mandelbrot {
	return &mandelbrot{
		maxIterations: maxIterations,
		density:       density,
	}
}

// TODO: Documenation
func (m mandelbrot) MaxIterations() uint {
	return m.maxIterations
}

// TODO: Documenation
func (m mandelbrot) Density() float64 {
	return m.density
}

// TODO: Documenation
func (m mandelbrot) IsStable(c complex128, z complex128) (bool, uint) {
	i := uint(0)

	for i <= m.MaxIterations() && complexAbs(z) <= 2 {
		z = complexPow2(z) + c
		i += 1
	}

	return complexAbs(z) <= 2, i
}

// TODO: Documenation
func (m mandelbrot) Draw(canvas *Canvas) {
	Fractal(canvas, m)
}
