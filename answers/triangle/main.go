package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.

type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

func init() {
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if nat(a, b, c) {
		return NaT
	} else if equ(a, b, c) {
		return Equ
	} else if iso(a, b, c) {
		return Iso
	} else if sca(a, b, c) {
		return Sca
	}

	return NaT
}

func equ(a, b, c float64) bool {
	return a == b && a == c && c == b
}

func iso(a, b, c float64) bool {
	return a == b || b == c || c == a
}

func sca(a, b, c float64) bool {
	return a != b || b != c || a != c
}

func nat(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0 ||
		a+b < c || a+c < b || b+c < a ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)
}
