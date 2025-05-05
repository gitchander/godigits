package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Invert struct {
	Content Object
}

var _ Object = Invert{}

func (Invert) IsObject() {}

func (v Invert) Draw(c *cairo.Canvas, r Bounds, level int) {

	const (
		lineWidthFactor = 0.05
	)

	var (
		rs = r.Size()
		w  = minFloat64(rs.X, rs.Y)

		lineWidth = w * lineWidthFactor
	)

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(lineWidth)

	margin := geom.MakeFrame4(0.05, 0.05, 0.15, 0.05)

	cr := r.Shrink(margin.MulScalar(w))

	y1 := (r.Min.Y + cr.Min.Y) / 2
	c.MoveTo(cr.Min.X, y1)
	c.LineTo(cr.Max.X, y1)

	c.Stroke()

	// Content
	{
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
