package gofractal

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
