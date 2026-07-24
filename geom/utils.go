package geom

import (
	"strconv"
)

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

//------------------------------------------------------------------------------

// Lerp - linear interpolation.
// t: [0..1]
// (t = 0) -> v0
// (t = 1) -> v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

func Lerp(v0, v1 float64, t float64) float64 {
	return lerp(v0, v1, t)
}

//------------------------------------------------------------------------------

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func formatInt(a int) string {
	return strconv.Itoa(a)
}

//------------------------------------------------------------------------------

// GCD - Greatest Common Denominator: largest number that can devide two numbers.
// GCD - Greatest Common Divisor
// https://en.wikipedia.org/wiki/Greatest_common_divisor

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//------------------------------------------------------------------------------
