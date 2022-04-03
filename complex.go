package gofractal

import "math"

// complexPow2 Second power of the complex number.
func complexPow2(c complex128) complex128 {
	x, y := real(c), imag(c)
	xy := x * y

	return complex((x*x - y*y), xy+xy)
}

// complexAbs Absolute value of the complex number.
func complexAbs(c complex128) uint {
	x, y := real(c), imag(c)

	return uint(math.Sqrt(x*x + y*y))
}
