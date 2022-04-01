package gofractal

import "testing"

func TestMandelbrot(t *testing.T) {
	canvas := NewCanvas(250, 200)
	mandel := NewMandelbrot(10, 0.01)

	mandel.Draw(canvas)
}

func TestMandelbrotIsStable(t *testing.T) {
	mandel := NewMandelbrot(10, 0.01)

	tables := []struct {
		c complex128
		r bool
	}{
		{complex(1, 0), false},
		{complex(0, 0), true},
		{complex(-2, 3), false},
		{complex(-0.5, -0.5), true},
		{complex(0.5, 0.5), false},
		{complex(-2, 1), false},
	}

	for _, table := range tables {
		result, _ := mandel.IsStable(table.c, 0+0i)

		if result != table.r {
			t.Errorf("Function absc was incorrect, with value %g, got: %t, want: %t.", table.c, result, table.r)
		}
	}
}
