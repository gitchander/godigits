package gorey

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Rhombus struct {
	Content Object
}

var _ Object = Rhombus{}

func (Rhombus) IsObject() {}

func (v Rhombus) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	var (
		lineWidthRel = 5.0

		// marginRel  = MakeFrame1(0)
		// paddingRel = MakeFrame1(0)

		marginRel  = geom.MakeFrame1(5)
		paddingRel = geom.MakeFrame1(5)

		// marginRel  = MakeFrame4(2, 4, 6, 8)
		// paddingRel = MakeFrame4(3, 6, 8, 10)
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
		r2i = (r1w - 2*(lineWidthAbs*math.Sqrt2)) / 4

		r2 = geom.PointToRect2f(center).Grow(geom.MakeFrame1(r2i))

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
		d := (r1w - (lineWidthAbs * math.Sqrt2)) / 2

		c.MoveTo(center.X, center.Y-d)
		c.LineTo(center.X+d, center.Y)
		c.LineTo(center.X, center.Y+d)
		c.LineTo(center.X-d, center.Y)
		c.LineTo(center.X, center.Y-d)

		cairoSetSourceColor(c, Black)
		c.SetLineWidth(lineWidthAbs)
		c.SetLineCap(cairo.LINE_CAP_SQUARE)
		c.Stroke()
	}

	// Content
	if v.Content != nil {
		v.Content.Draw(c, r3, level+1)
	}
}
