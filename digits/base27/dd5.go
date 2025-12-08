package base27

import (
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawer5 struct{}

func (DigitDrawer5) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	lw := 10.0
	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)
	c.SetRGB(0, 0, 0)

	const tau = 2 * math.Pi
	var (
		vs = make([]geom.Point2f, 6)

		angle      = -tau / 4
		deltaAngle = tau / 6

		center = geom.MakePoint2f(50, 100)
	)
	for i := range vs {
		var (
			p  = geom.MakePolar(40, angle)
			c  = geom.PolarToCartesian(p)
			cc = center.Add(c)
		)
		vs[i] = cc
		angle += deltaAngle
	}

	gd := newGeomDrawer(c)

	if true {
		// for _, v := range vs {
		// 	gd.DrawLine(center, v)
		// }

		gd.DrawLine(center, vs[0])
		gd.DrawLine(center, vs[2])
		gd.DrawLine(center, vs[4])

		// gd.DrawLine(center, vs[1])
		// gd.DrawLine(center, vs[3])
		// gd.DrawLine(center, vs[5])

		c.Stroke()
	}

	trits := calcTritsBal27(digit)

	//--------------------------------------------------------------------------

	switch t := trits[0]; t {
	case -1:
		gd.DrawLine(vs[5], vs[0])
	case 0:
	case +1:
		gd.DrawLine(vs[0], vs[1])
	}

	switch t := trits[1]; t {
	case -1:
		gd.DrawLine(vs[1], vs[2])
	case 0:
	case +1:
		gd.DrawLine(vs[2], vs[3])
	}

	switch t := trits[2]; t {
	case -1:
		gd.DrawLine(vs[3], vs[4])
	case 0:
	case +1:
		gd.DrawLine(vs[4], vs[5])
	}

	//--------------------------------------------------------------------------

	c.Stroke()
}
