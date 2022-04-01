package gofractal

import (
	"testing"
)

func TestPixelAdd(t *testing.T) {
	type tableTest struct {
		start, input, expected Pixel
	}

	tables := []tableTest{
		{Pixel{0, 0}, Pixel{1, 1}, Pixel{1, 1}},
		{Pixel{1, 1}, Pixel{1, 1}, Pixel{2, 2}},
		{Pixel{1, 2}, Pixel{5, 2}, Pixel{6, 4}},
		{Pixel{1, 4}, Pixel{7, 4}, Pixel{8, 8}},
		{Pixel{0, 0}, Pixel{1, 0}, Pixel{1, 0}},
		{Pixel{0, 0}, Pixel{0, 1}, Pixel{0, 1}},
		{Pixel{0, 0}, Pixel{0, 1}, Pixel{0, 1}},
	}

	for _, table := range tables {
		result := table.start
		result.Add(table.input)

		if result.X() != table.expected.X() && result.Y() != table.expected.Y() {
			t.Errorf("Function Pixel.Add was incorrect, with value %v, got: %v, want: %v.", table.start, result, table.expected)
		}
	}
}
