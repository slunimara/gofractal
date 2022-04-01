package gofractal

import (
	"testing"
)

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
