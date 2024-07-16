package main

import (
	"math"
)

const tau = 2 * math.Pi

// Lerp - linear interpolation.
// t: [0..1]
// (t = 0) -> v0
// (t = 1) -> v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1-t)*v0 + t*v1
}

func normalize(x float64, min, max float64) float64 {
	return (x - min) / (max - min)
}

func uint16GetBit(x uint16, i int) uint {
	return uint((x >> i) & 1)
}

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
