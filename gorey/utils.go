package gorey

import (
	"strconv"

	"github.com/gitchander/godigits/geom"
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

// ------------------------------------------------------------------------------
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------------------------------------------------------------
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

// ------------------------------------------------------------------------------
// pow2, a^2, a*a
func square(a float64) float64 {
	return a * a
}

// pow3, a^3, a*a*a
func cube(a float64) float64 {
	return a * a * a
}

// ------------------------------------------------------------------------------
func vmin(b geom.Bounds) float64 {
	return minFloat64(b.Dx(), b.Dy()) / 100.0
}

func vmax(b geom.Bounds) float64 {
	return maxFloat64(b.Dx(), b.Dy()) / 100.0
}

//------------------------------------------------------------------------------
