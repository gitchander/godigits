package hexagone

import (
	"math"

	"github.com/gitchander/godigits/geom"
)

const tau = 2.0 * math.Pi

func HexagoneVertexes(center geom.Point2f, radius float64) []geom.Point2f {
	ps := make([]geom.Point2f, 6)
	theta := -tau / 4
	dt := tau / float64(len(ps))
	for i := range ps {
		p := center.Add(geom.PolarToCartesian(geom.ShPolar(radius, theta)))
		ps[i] = p
		theta += dt
	}
	return ps
}

func ShrinkPoints(a, b geom.Point2f, t float64) (a1, b1 geom.Point2f) {
	a1 = geom.PtLerp(a, b, t)
	b1 = geom.PtLerp(b, a, t)
	return a1, b1
}

func clampFloat64(x float64, min, max float64) float64 {
	if min > max { // empty interval
		return 0 // default value
	}
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func clampFloat64Min(x float64, min float64) float64 {
	if x < min {
		return min
	}
	return x
}

func clampFloat64Max(x float64, max float64) float64 {
	if x > max {
		return max
	}
	return x
}

func SegmentPolygone(a, b geom.Point2f, lineWidth float64, angle float64) []geom.Point2f {

	//-------------------------------------------------------
	angle = clampFloat64(angle, 0, math.Pi)
	da := angle / 2

	if da < math.Pi/2 {
		d := geom.Distance(a, b)
		wmax := d * math.Tan(da)
		lineWidth = clampFloat64(lineWidth, 0, wmax)
	} else {
		lineWidth = clampFloat64Min(lineWidth, 0)
	}

	radius := (lineWidth / (2 * math.Sin(da)))

	//-------------------------------------------------------

	// Point a:
	var (
		deltaA = b.Sub(a)
		angleA = math.Atan2(deltaA.Y, deltaA.X)

		a1 = a.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleA+da)))
		a2 = a.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleA-da)))
	)

	// Point b:
	var (
		deltaB = a.Sub(b)
		angleB = math.Atan2(deltaB.Y, deltaB.X)

		b1 = b.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleB+da)))
		b2 = b.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleB-da)))
	)

	ps := []geom.Point2f{
		geom.Pt2f(a.X, a.Y),
		geom.Pt2f(a1.X, a1.Y),
		geom.Pt2f(b2.X, b2.Y),
		geom.Pt2f(b.X, b.Y),
		geom.Pt2f(b1.X, b1.Y),
		geom.Pt2f(a2.X, a2.Y),
	}
	return ps
}
