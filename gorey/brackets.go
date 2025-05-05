package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

// Curly brackets or braces
type CurlyBrackets struct {
	Content Object
}

var _ Object = CurlyBrackets{}

func (CurlyBrackets) IsObject() {}

func (v CurlyBrackets) Draw(c *cairo.Canvas, r Bounds, level int) {

	const (
		lineWidthRelative = 0.05
		radiusRelative    = 0.1
	)

	var (
		// margin  = MakeFrame1(0)
		// padding = MakeFrame1(0)

		margin  = geom.MakeFrame1(0.05)
		padding = geom.MakeFrame1(0.05)

		// margin  = MakeFrame2(0.05, 0)
		// padding = MakeFrame2(0.05, 0)
	)

	var (
		w = minFloat64(r.Dx(), r.Dy())

		lineWidth = w * lineWidthRelative
		radius    = w * radiusRelative
	)

	var (
		r1 = r.Shrink(margin.MulScalar(w))
		r2 = r1.Shrink(geom.MakeFrame2(radius, 0).MulScalar(2))
		r3 = r2.Shrink(padding.MulScalar(w))
	)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()
		//---------------------------------------------
		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, r1)
		c.Fill()
		//---------------------------------------------
		cairoSetSourceColor(c, levelToColor(2))
		cairoRectangle(c, r2)
		c.Fill()
		//---------------------------------------------
		cairoSetSourceColor(c, levelToColor(3))
		cairoRectangle(c, r3)
		c.Fill()
	}

	// draw frame
	{
		center := r1.Center()

		var (
			x1 = r1.Min.X + radius
			y1 = r1.Min.Y + radius + lineWidth/2

			x2 = r1.Max.X - radius
			y2 = r1.Max.Y - radius - lineWidth/2
		)

		r := radius

		yc := center.Y

		// left
		{
			c.MoveTo(x1+r, y1-r)
			BesierQuad(c, geom.Pt2f(x1+r, y1-r), geom.Pt2f(x1, y1-r), geom.Pt2f(x1, y1))
			//c.LineTo(x1, yc-r)
			BesierQuad(c, geom.Pt2f(x1, yc-r), geom.Pt2f(x1, yc), geom.Pt2f(x1-r, yc))
			BesierQuad(c, geom.Pt2f(x1-r, yc), geom.Pt2f(x1, yc), geom.Pt2f(x1, yc+r))
			//c.LineTo(x1, y2)
			BesierQuad(c, geom.Pt2f(x1, y2), geom.Pt2f(x1, y2+r), geom.Pt2f(x1+r, y2+r))
		}

		// right
		{
			c.MoveTo(x2-r, y1-r)
			BesierQuad(c, geom.Pt2f(x2-r, y1-r), geom.Pt2f(x2, y1-r), geom.Pt2f(x2, y1))
			//c.LineTo(x2, yc-r)
			BesierQuad(c, geom.Pt2f(x2, yc-r), geom.Pt2f(x2, yc), geom.Pt2f(x2+r, yc))
			BesierQuad(c, geom.Pt2f(x2+r, yc), geom.Pt2f(x2, yc), geom.Pt2f(x2, yc+r))
			//c.LineTo(x2, y2)
			BesierQuad(c, geom.Pt2f(x2, y2), geom.Pt2f(x2, y2+r), geom.Pt2f(x2-r, y2+r))
		}

		cairoSetSourceColor(c, Black)
		c.SetLineWidth(lineWidth)
		c.Stroke()
	}

	if v.Content != nil {
		v.Content.Draw(c, r3, level+1)
	}
}
