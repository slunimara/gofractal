package gofractal

import "testing"

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
