package gofractal

type mandelbrot struct {
	maxIterations uint
	view          View
}

// TODO: Documenation
func NewMandelbrot(maxIterations uint, view View) *mandelbrot {
	return &mandelbrot{
		maxIterations: maxIterations,
		view:          view,
	}
}

// TODO: Documenation
func (m mandelbrot) MaxIterations() uint {
	return m.maxIterations
}

// TODO: Documenation
func (m mandelbrot) View() *View {
	return &m.view
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
