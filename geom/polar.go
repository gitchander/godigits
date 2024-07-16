package geom

import (
	"image"
	"math"
)

// ρ - rho

// φ - phi
// θ - theta

// type Polar struct {
// 	r float64 // rho
// 	t float64 // phi
// }

type Polar struct {
	Rho float64
	Phi float64
}

func _() {
	image.Pt(0, 0)
}

// Pt is shorthand for Point{X, Y}.
// Pr is shorthand for Polar{Rho, Phi}.
// Pr, MakePolar, ShPolar
// ShPolar is shorthand for Polar{Rho, Phi}.
func ShPolar(r, φ float64) Polar {
	return Polar{
		Rho: r,
		Phi: φ,
	}
}

func (p Polar) ToCartesian() Point2f {
	return PolarToCartesian(p)
}

// https://en.wikipedia.org/wiki/Polar_coordinate_system

// The radial coordinate is often denoted by r or ρ.
// The angular coordinate by φ, θ, or t.

// r and φ

func polarToCartesian(r, φ float64) (x, y float64) {
	sin, cos := math.Sincos(φ)
	x = r * cos
	y = r * sin
	return
}

func cartesianToPolar(x, y float64) (r, φ float64) {
	r = math.Hypot(x, y)
	φ = math.Atan2(y, x)
	return
}

func PolarToCartesian(p Polar) Point2f {
	x, y := polarToCartesian(p.Rho, p.Phi)
	return Pt2f(x, y)
}

func CartesianToPolar(a Point2f) Polar {
	r, φ := cartesianToPolar(a.X, a.Y)
	return Polar{
		Rho: r,
		Phi: φ,
	}
}
