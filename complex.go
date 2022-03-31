package gofractal

import "math"

// Second power of the complex number.
func cPow(c complex128) complex128 {
	x, y := real(c), imag(c)
	xy := x * y
	return complex((x*x - y*y), xy+xy)
}

// Absolute value of the complex number.
func cAbs(c complex128) uint {
	x, y := real(c), imag(c)
	return uint(math.Sqrt(x*x + y*y))
}
