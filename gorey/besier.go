package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
	ivl "github.com/gitchander/godigits/utils/interval"
)

// Bézier curve
// https://en.wikipedia.org/wiki/B%C3%A9zier_curve
// https://ru.wikipedia.org/wiki/%D0%9A%D1%80%D0%B8%D0%B2%D0%B0%D1%8F_%D0%91%D0%B5%D0%B7%D1%8C%D0%B5

// Quadratic Bézier
func BesierQuad(c *cairo.Canvas, p0, p1, p2 geom.Point2f) {

	const koef = 2.0 / 3.0

	var (
		p01 = geom.PtLerp(p0, p1, koef)
		p21 = geom.PtLerp(p2, p1, koef)
	)

	BesierCubic(c, p0, p01, p21, p2)
}

func BesierCubic(c *cairo.Canvas, p0, p1, p2, p3 geom.Point2f) {

	startNewPath := false

	if startNewPath {
		c.MoveTo(p0.X, p0.Y)
	} else {
		c.LineTo(p0.X, p0.Y)
	}

	c.CurveTo(
		p1.X, p1.Y,
		p2.X, p2.Y,
		p3.X, p3.Y,
	)
}

// ------------------------------------------------------------------------------
func besierQuadCustom(dc *cairo.Canvas, a, b, c geom.Point2f) {

	var (
		n = 30
		l = ivl.IntervalFloat{
			Min: 0,
			Max: 1,
		}
	)

	tv := ivl.NewIntervalValuer(l, n)

	for i := 0; i < n; i++ {
		t := tv.IndexToValue(i)

		var (
			ab = geom.PtLerp(a, b, t)
			bc = geom.PtLerp(b, c, t)

			abc = geom.PtLerp(ab, bc, t)

			p = abc
		)

		if i == 0 {
			dc.MoveTo(p.X, p.Y)
		} else {
			dc.LineTo(p.X, p.Y)
		}
	}
}

func besierCubicCustom(dc *cairo.Canvas, a, b, c, d geom.Point2f) {

	var (
		n = 30
		l = ivl.IntervalFloat{
			Min: 0,
			Max: 1,
		}
	)

	tv := ivl.NewIntervalValuer(l, n)

	for i := 0; i < n; i++ {
		t := tv.IndexToValue(i)

		var (
			ab = geom.PtLerp(a, b, t)
			bc = geom.PtLerp(b, c, t)
			cd = geom.PtLerp(c, d, t)

			abc = geom.PtLerp(ab, bc, t)
			bcd = geom.PtLerp(bc, cd, t)

			abcd = geom.PtLerp(abc, bcd, t)

			p = abcd
		)

		if i == 0 {
			dc.MoveTo(p.X, p.Y)
		} else {
			dc.LineTo(p.X, p.Y)
		}
	}
}
