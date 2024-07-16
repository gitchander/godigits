package gorey

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Rectangle struct {
	Content Object
}

var _ Object = Rectangle{}

func (Rectangle) IsObject() {}

func (v Rectangle) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	const (
		lineWidthRel = 5
	)

	var (
		// marginRel  = MakeFrame1(0)
		// paddingRel = MakeFrame1(0)

		marginRel  = geom.MakeFrame1(5)
		paddingRel = geom.MakeFrame1(5)

		// marginRel  = MakeFrame4(12, 4, 16, 8)
		// paddingRel = MakeFrame4(3, 6, 8, 1)
	)

	var (
		vm = vmin(r)

		lineWidthAbs = lineWidthRel * vm
		marginAbs    = marginRel.MulScalar(vm)
		paddingAbs   = paddingRel.MulScalar(vm)
	)

	var (
		r1 = r.Shrink(marginAbs)
		r2 = r1.Shrink(geom.MakeFrame1(lineWidthAbs / 2))
		r3 = r1.Shrink(geom.MakeFrame1(lineWidthAbs))
		r4 = r3.Shrink(paddingAbs)
	)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, r1)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(2))
		cairoRectangle(c, r3)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(3))
		cairoRectangle(c, r4)
		c.Fill()
	}

	// draw frame
	if true {
		cairoSetSourceColor(c, Black)
		c.SetLineWidth(lineWidthAbs)
		cairoRectangle(c, r2)
		c.Stroke()
	}

	// Content
	if v.Content != nil {
		v.Content.Draw(c, r4, level+1)
	}
}

// ------------------------------------------------------------------------------
func Square(o Object) Object {
	return SquareAspector(Rectangle{o})
}

// ------------------------------------------------------------------------------
// Standard ISO216 (A0, A1, ... A4, ... A7)
func ISO216(o Object) Object {
	return Aspector{
		AspectRatio: math.Sqrt2,
		Content: Rectangle{
			o,
		},
	}
}
