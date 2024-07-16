package gorey

import (
	"image"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type DigitV1 struct {
	Value int
}

var _ Object = DigitV1{}

func (DigitV1) IsObject() {}

func (d DigitV1) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	const (
		lineWidthRel = 0.5
	)

	var (
		marginRel = geom.MakeFrame1(0)
		//marginRel = geom.MakeFrame1(5)
	)

	var (
		vm = vmin(r)

		marginAbs = marginRel.MulScalar(vm)
	)

	dSize := image.Pt(8, 8)
	aspectRatio := float64(dSize.Y) / float64(dSize.X)

	var (
		r1 = r.Shrink(marginAbs)

		r2 = subRectByAspectRatio(r1, aspectRatio)

		cellSize = r2.Dx() / float64(dSize.X)

		wLineWidth = cellSize * lineWidthRel
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
	}

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(wLineWidth)

	dc := NewICanvas(c, r2, dSize.X, dSize.Y)

	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.SetLineJoin(cairo.LINE_JOIN_ROUND)

	// Draw base
	{
		dc.MoveTo(4, 1)
		dc.LineTo(4, 7)
		dc.LineTo(6, 7)
		c.Stroke()
	}

	//c.SetLineCap(cairo.LINE_CAP_ROUND)
	//c.SetLineCap(cairo.LINE_CAP_SQUARE)

	const maxValue = 20
	v := ((d.Value + maxValue) % (2*maxValue + 1)) - maxValue

	var negative bool
	if v < 0 {
		negative = true
		v = -v
	}

	x0, dx := 5, 1
	if negative {
		x0, dx = 3, -1
	}

	quo, rem := quoRem(v, 4)

	var (
		xi1 = x0 + 0*dx
		xi2 = x0 + 2*dx
	)

	yi := 2
	for i := 0; i < quo; i++ {
		dc.MoveTo(xi1, yi)
		dc.LineTo(xi2, yi)
		yi++
	}
	c.Stroke()

	xi := x0
	for i := 0; i < rem; i++ {
		dc.Circle(xi, yi, wLineWidth/2)
		xi += dx
	}
	c.Fill()
}

func Dig1(x int) Object {
	return DigitV1{
		Value: x,
	}
}
