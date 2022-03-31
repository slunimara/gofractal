package gofractal

import "testing"

func TestIsStable(t *testing.T) {
	numberInterations := uint(10)

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
		result, _ := isStable(table.c, numberInterations)

		if result != table.r {
			t.Errorf("Function absc was incorrect, with value %g, got: %t, want: %t.", table.c, result, table.r)
		}
	}
}

func TestArange(t *testing.T) {
	tables := []struct {
		start float64
		stop  float64
		step  float64
		val   []float64
	}{
		{1, 10, 1, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{-3, 3, 1, []float64{-3, -2, -1, 0, 1, 2, 3}},
		{5, 2, 1, []float64{5, 4, 3, 2}},
	}

	for _, table := range tables {
		result := arange(table.start, table.stop, table.step)

		for i, v := range result {
			if v != table.val[i] {
				t.Errorf("Function arange was incorrect, got value: %g on the index %d, want: %g.", table.val[i], i, v)
				break
			}
		}
	}
}

func TestMandelbrot(t *testing.T) {
	canvas := NewCanvas(250, 200)

	Mandelbrot(canvas, 10, 0.01)
}
