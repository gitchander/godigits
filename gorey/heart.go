package gorey

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
	ivl "github.com/gitchander/godigits/utils/interval"
)

type Heart struct{}

var _ Object = Heart{}

func (Heart) IsObject() {}

func (v Heart) Draw(c *cairo.Canvas, r Bounds, level int) {

	const lineWidthRel = 5

	var (
		marginRel = geom.MakeFrame1(5)
		//paddingRel = MakeFrame1(5)
	)

	var (
		vm = vmin(r)

		lineWidthAbs = lineWidthRel * vm
		marginAbs    = marginRel.MulScalar(vm)
		//paddingAbs   = paddingRel.MulScalar(vm)
	)

	var (
		//size   = Pt2f(1, 1)
		r1     = r.Shrink(marginAbs)
		r1w    = minFloat64(r1.Dx(), r1.Dy())
		center = r1.Center()
	)

	w := r1w

	c.SetSourceRGB(0, 0, 0)
	c.SetLineWidth(lineWidthAbs)

	ti := ivl.IntervalFloat{
		Min: 0,
		Max: tau,
	}

	scale := 0.027 * w
	n := 87

	var (
		t0 = ti.Min
		dt = ti.Width() / float64(n-1)
	)

	for i := 0; i < n; i++ {

		t := t0 + (dt * float64(i))

		p := HeartCurve(t)
		p = p.MulScalar(scale)
		p = p.InvertAxisY()
		p = p.Add(center)

		if i == 0 {
			c.MoveTo(p.X, p.Y)
		} else {
			c.LineTo(p.X, p.Y)
		}

		// c.Arc(p.X, p.Y, 0.003, 0, tau)
		// c.Fill()
	}
	if true {
		c.Stroke()
	} else {
		c.SetSourceRGB(1, 0, 0)
		c.Fill()
	}
}

// Heart Curve
// https://mathworld.wolfram.com/HeartCurve.html
func HeartCurve(t float64) geom.Point2f {

	sinT, cosT := math.Sincos(t)

	var (
		x = 16 * cube(sinT)
		y = (13 * cosT) - (5 * math.Cos(2*t)) - (2 * math.Cos(3*t)) - math.Cos(4*t)
	)

	return geom.Point2f{
		X: x,
		Y: y,
	}
}
