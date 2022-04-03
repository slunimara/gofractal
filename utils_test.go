package gofractal

import "testing"

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

// TODO: Write better test
func TestGCD(t *testing.T) {
	gcdFunc := func(a, b uint64) uint64 {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}

		return a
	}

	a := uint64(10)
	b := uint64(20)

	if gcdFunc(a, b) != GCD(a, b) {
		t.Errorf(
			"Function GCD was incorrect, got value: %d, want: %d.",
			GCD(a, b),
			gcdFunc(a, b))
	}
}
