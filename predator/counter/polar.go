package main

import (
	"math"
)

// ρ - rho

// φ - phi
// θ - theta

type Polar struct {
	Rho float64 // Radius
	Phi float64 // Angle
}

// https://en.wikipedia.org/wiki/Polar_coordinate_system

// The radial coordinate is often denoted by r or ρ.
// The angular coordinate by φ, θ, or t.

// r and φ

func PolarToCartesian(r, φ float64) (x, y float64) {
	sin, cos := math.Sincos(φ)
	x = r * cos
	y = r * sin
	return
}

func CartesianToPolar(x, y float64) (r, φ float64) {
	r = math.Hypot(x, y)
	φ = math.Atan2(y, x)
	return
}
