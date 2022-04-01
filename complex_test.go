package gofractal

import (
	"math/cmplx"
	"testing"
)

func TestComplexPow2(t *testing.T) {
	tables := []struct {
		c complex128
		v complex128
	}{
		{complex(1, 0), complex(1, 0)},
		{complex(1, 1), complex(0, 2)},
		{complex(5, 2), complex(21, 20)},
		{complex(7, 4), complex(33, 56)},
	}

	for _, table := range tables {
		result := complexPow2(table.c)

		if result != table.v {
			t.Errorf("Function complexPow2 was incorrect, with value %g, got: %g, want: %g.", table.c, result, table.v)
		}
	}
}

func TestComplexAbs(t *testing.T) {
	tables := []struct {
		c complex128
		v uint
	}{
		{complex(1, 0), uint(cmplx.Abs(complex(1, 0)))},
		{complex(1, 1), uint(cmplx.Abs(complex(1, 1)))},
		{complex(5, 2), uint(cmplx.Abs(complex(5, 2)))},
		{complex(7, 4), uint(cmplx.Abs(complex(7, 4)))},
	}

	for _, table := range tables {
		result := complexAbs(table.c)

		if result != table.v {
			t.Errorf("Function complexAbs was incorrect, with value %g, got: %d, want: %d.", table.c, result, table.v)
		}
	}
}
