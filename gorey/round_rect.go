package gorey

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type RoundRect struct {
	Content Object
}

var _ Object = RoundRect{}

func (RoundRect) IsObject() {}

func (v RoundRect) Draw(c *cairo.Canvas, r Bounds, level int) {

	var (
		lineWidthRel = 5.0
		radiusRel    = 15.0

		marginRel  = geom.MakeFrame1(5)
		paddingRel = geom.MakeFrame1(5)
	)

	var (
		vm = vmin(r)

		lineWidthAbs = lineWidthRel * vm
		radiusAbs    = radiusRel * vm
		marginAbs    = marginRel.MulScalar(vm)
		paddingAbs   = paddingRel.MulScalar(vm)
	)

	var (
		r1 = r.Shrink(marginAbs)

		// r2 = r1.Shrink(MakeFrame1(radiusAbs + lineWidthAbs/2 - (radiusAbs-lineWidthAbs/2)/math.Sqrt2))
		r2 = r1.Shrink(geom.MakeFrame1((radiusAbs-lineWidthAbs/2)*(1.0-1.0/math.Sqrt2) + lineWidthAbs))

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
		var (
			useBesier = true

			ra = radiusAbs
			r2 = r1.Shrink(geom.MakeFrame1(radiusAbs + lineWidthAbs/2))

			x1, y1 = r2.Min.X, r2.Min.Y
			x2, y2 = r2.Max.X, r2.Max.Y
		)

		if useBesier {
			BesierQuad(c, geom.Pt2f(x1-ra, y1), geom.Pt2f(x1-ra, y1-ra), geom.Pt2f(x1, y1-ra))
		} else {
			c.Arc(x1, y1, ra, -tau/2, -tau/4)
		}
		c.LineTo(x2, y1-ra)

		if useBesier {
			BesierQuad(c, geom.Pt2f(x2, y1-ra), geom.Pt2f(x2+ra, y1-ra), geom.Pt2f(x2+ra, y1))
		} else {
			c.Arc(x2, y1, ra, tau*3/4, tau)
		}
		c.LineTo(x2+ra, y2)

		if useBesier {
			BesierQuad(c, geom.Pt2f(x2+ra, y2), geom.Pt2f(x2+ra, y2+ra), geom.Pt2f(x2, y2+ra))
		} else {
			c.Arc(x2, y2, ra, 0, tau/4)
		}
		c.LineTo(x1, y2+ra)

		if useBesier {
			BesierQuad(c, geom.Pt2f(x1, y2+ra), geom.Pt2f(x1-ra, y2+ra), geom.Pt2f(x1-ra, y2))
		} else {
			c.Arc(x1, y2, ra, tau/4, tau/2)
		}
		c.LineTo(x1-ra, y1)

		cairoSetSourceColor(c, Black)
		c.SetLineWidth(lineWidthAbs)

		c.Stroke()
	}

	// Content
	if v.Content != nil {
		v.Content.Draw(c, r3, level+1)
	}
}
