package gorey

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Circle struct {
	Content Object
}

var _ Object = Circle{}

func (Circle) IsObject() {}

func (v Circle) Draw(c *cairo.Canvas, r Bounds, level int) {

	const (
		// Rel - Relative
		lineWidthRel = 5
	)

	var (
		// marginRel  = geom.MakeFrame1(0)
		// paddingRel = geom.MakeFrame1(0)

		marginRel  = geom.MakeFrame1(5)
		paddingRel = geom.MakeFrame1(5)

		// marginRel  = MakeFrame4(2, 4, 6, 8)
		// paddingRel = MakeFrame4(8, 6, 4, 2)
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
		radius = (r1w - lineWidthAbs) / 2
		r2i    = (r1w/2 - lineWidthAbs) / math.Sqrt2

		r2 = geom.Point2fToBounds(center).Grow(geom.MakeFrame1(r2i))
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
		cairoRectangle(c, r2)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(3))
		cairoRectangle(c, r3)
		c.Fill()
	}

	// draw frame
	{
		cairoSetSourceColor(c, Black)
		c.SetLineWidth(lineWidthAbs)
		cairoCircle(c, center, radius)
		c.Stroke()
	}

	// Content
	if v.Content != nil {
		v.Content.Draw(c, r3, level+1)
	}
}
