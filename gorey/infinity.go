package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Infinity struct{}

var _ Object = Infinity{}

func (Infinity) IsObject() {}

func (v Infinity) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	var (
		marginRel = geom.MakeFrame1(5)
		//marginRel = geom.MakeFrame4(15, 10, 20, 20)
	)

	var (
		vm = vmin(r)

		marginAbs = marginRel.MulScalar(vm)
	)

	r1 := r.Shrink(marginAbs)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, r1)
		c.Fill()
	}

	//drawInfinityByBesierCurve(c)
	drawLemniscate(c, r1)
}

func drawInfinityByBesierCurve(c *cairo.Canvas) {

	if false {

		margin := 0.06

		x1, y1 := margin, margin
		xc, yc := 0.5, 0.5
		x2, y2 := (1 - margin), (1 - margin)

		c.MoveTo(xc, yc)
		c.CurveTo(xc, yc, x1, y2, x1, yc)
		c.CurveTo(x1, yc, x1, y1, xc, yc)
		c.CurveTo(xc, yc, x2, y1, x2, yc)
		c.CurveTo(x2, yc, x2, y2, xc, yc)

	} else {

		margin := -0.06

		x1, y1 := margin, margin
		xc, yc := 0.5, 0.5
		x2, y2 := (1 - margin), (1 - margin)

		c.MoveTo(xc, yc)
		c.CurveTo(x1, y2, x1, y1, xc, yc)
		c.CurveTo(x2, y2, x2, y1, xc, yc)
	}

	c.Stroke()
}
