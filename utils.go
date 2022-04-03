package gofractal

import "math"

// Creates an array from start to the end divided by n segments that's given by interval distribution based on step parameter.
func arange(start, stop, step float64) []float64 {
	N := int(IntervalDistribution(start, stop, step))
	arr := make([]float64, N)

	if start > stop {
		step *= -1
	}

	for i := range arr {
		arr[i] = start + (float64(i) * step)
	}

	return arr
}

// Determine the value of a variable. Formula: a : b = c : return
func CrossMultiplication(a, b, c float64) float64 {
	return b * c / a
}

// Computes the number of segments from start to end with given step.
func IntervalDistribution(start, stop, step float64) float64 {
	if step < 0 {
		return 0
	} else if step == 0 {
		return 1
	}

	return math.Ceil((math.Abs(stop - start)) / step)
}

// GCD Greatest Common Divisor of two numbers. Implemented with binary algorithm.
func GCD(a, b uint64) uint64 {
	var shift uint64

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	// shift <-- log K, where K is the greatest power of 2 dividing a and b
	for ; ((a | b) & 1) == 0; shift++ {
		a >>= 1
		b >>= 1
	}

	// Continue dividing a unit its odd
	for (a & 1) == 0 {
		a >>= 1
	}

	ok := true
	for ok {
		for (b & 1) == 0 {
			b >>= 1
		}

		if a > b {
			b, a = a, b
		}

		b -= a

		ok = (b != 0)
	}

	return a << shift
}
