package gorey

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type RegularPolygon struct {
	N       int
	Phase   float64 // [0..1]
	Content Object
}

var _ Object = RegularPolygon{}

func (RegularPolygon) IsObject() {}

func (v RegularPolygon) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	const (
		lineWidthRel = 5.0
	)

	var (
		marginRel = geom.MakeFrame1(5)
		//marginRel = MakeFrame4(2, 4, 6, 8)

		paddingRel = geom.MakeFrame1(5)
	)

	var (
		vm = vmin(r)

		lineWidthAbs = lineWidthRel * vm
		marginAbs    = marginRel.MulScalar(vm)
		paddingAbs   = paddingRel.MulScalar(vm)
	)

	var (
		r1     = r.Shrink(marginAbs)
		r1w    = minFloat64(r1.Dx(), r1.Dy())
		center = r1.Center()
	)

	var (
		deltaAngle = tau / float64(v.N)

		radius  = (r1w - lineWidthAbs) / 2
		radiuIn = radius * math.Cos(deltaAngle/2)

		r2 = geom.PointToRect2f(center).Grow(geom.MakeFrame1((radiuIn - (lineWidthAbs / 2)) / math.Sqrt2))

		r3 = r2.Shrink(paddingAbs)
	)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, r1)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(2))
		cairoCircle(c, center, radius+(lineWidthAbs/2))
		c.Fill()

		cairoSetSourceColor(c, levelToColor(3))
		cairoCircle(c, center, radiuIn-(lineWidthAbs/2))
		c.Fill()

		cairoSetSourceColor(c, levelToColor(4))
		cairoRectangle(c, r2)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(5))
		cairoRectangle(c, r3)
		c.Fill()
	}

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(lineWidthAbs)

	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.SetLineJoin(cairo.LINE_JOIN_ROUND)

	phaseTau := v.Phase * tau

	n := v.N

	getPoint := func(theta float64) geom.Point2f {
		p := geom.PolarToCartesian(geom.ShPolar(radius, theta))
		p = p.InvertAxisY()
		return center.Add(p)
	}

	angle := phaseTau
	if n > 0 {
		p := getPoint(angle)
		c.MoveTo(p.X, p.Y)
		angle += deltaAngle
	}

	for i := 0; i < n; i++ {
		p := getPoint(angle)
		c.LineTo(p.X, p.Y)
		angle += deltaAngle
	}

	c.Stroke()

	// Content
	if v.Content != nil {
		v.Content.Draw(c, r3, level+1)
	}
}
