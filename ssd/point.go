package main

import (
	"github.com/gitchander/godigits/geom"
)

type Point2f = geom.Point2f

func Pt(x, y float64) Point2f {
	return Point2f{
		X: x,
		Y: y,
	}
}

//------------------------------------------------------------------------------

// Shrink polygon
func ShrinkPolygon(ps []geom.Point2f) []geom.Point2f {

	n := len(ps)

	vs := make([]geom.Point2f, n)

	const (
		//shrinkWidth = 0.0
		//shrinkWidth = 0.1
		shrinkWidth = 0.25
	)

	for i := 0; i < n; i++ {

		var (
			a = ps[mod(i-1, n)]
			b = ps[i]
			c = ps[mod(i+1, n)]
		)

		var (
			n1 = lineNorm(a, b).MulScalar(shrinkWidth)

			p1 = a.Sub(n1)
			p2 = b.Sub(n1)
		)

		var (
			n2 = lineNorm(b, c).MulScalar(shrinkWidth)

			p3 = b.Sub(n2)
			p4 = c.Sub(n2)
		)

		p, _ := intersection(p1.X, p1.Y, p2.X, p2.Y, p3.X, p3.Y, p4.X, p4.Y)

		vs[i] = p
	}

	return vs
}

// L1:{(x1,y1)-(x2,y2)}, L2:{(x3,y3)-(x4,y4)}
func intersection(x1, y1, x2, y2, x3, y3, x4, y4 float64) (geom.Point2f, bool) {

	d := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)

	if d == 0 {
		return geom.Point2f{}, false
	}

	v1 := x1*y2 - y1*x2
	v2 := x3*y4 - y3*x4

	vx := (v1 * (x3 - x4)) - ((x1 - x2) * v2)
	vy := (v1 * (y3 - y4)) - ((y1 - y2) * v2)

	return geom.Point2f{
		X: vx / d,
		Y: vy / d,
	}, true
}
