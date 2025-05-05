package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

// rune: '^'
type Caret struct {
	Content Object
}

var _ Object = Caret{}

func (Caret) IsObject() {}

func (v Caret) Draw(c *cairo.Canvas, r Bounds, level int) {

	const (
		signHeightFactor = 0.25

		lineWidthFactor = 0.25
		dxFactor        = 0.5
		dyFactor        = 0.5
	)

	var (
		w = minFloat64(r.Dx(), r.Dy())

		signHeight = w * signHeightFactor

		lineWidth = signHeight * lineWidthFactor
		dx        = signHeight * dxFactor
		dy        = signHeight * dyFactor
	)

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(lineWidth)

	//c.SetLineJoin(cairo.LINE_JOIN_BEVEL)
	c.SetLineJoin(cairo.LINE_JOIN_MITER)
	//c.SetLineJoin(cairo.LINE_JOIN_ROUND)

	center := r.Center()

	y1 := r.Min.Y + (signHeight-dy)/2

	c.MoveTo(center.X-dx, y1+dy)
	c.LineTo(center.X, y1)
	c.LineTo(center.X+dx, y1+dy)

	c.Stroke()

	// Content
	{
		var (
			margin = geom.MakeFrame4(0.05, 0.05, signHeightFactor, 0.05)
			cr     = r.Shrink(margin.MulScalar(w))
		)
		if DrawLevelArea {
			cairoSetSourceColor(c, levelToColor(level))
			cairoRectangle(c, cr)
			c.Fill()
		}
		if v.Content != nil {
			v.Content.Draw(c, cr, level+1)
		}
	}
}
