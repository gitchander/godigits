package base3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

const AspectRatio = 2 // dy / dx

type DigitDrawer1 struct{}

func (DigitDrawer1) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	var (
		lw = 10.0

		x0 = 10.0
		x1 = 50.0
		x2 = 90.0

		y0 = 30.0
		y1 = 70.0
		y2 = 130.0
		y3 = 170.0
	)

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)
	c.SetRGB(0, 0, 0)

	c.MoveTo(x1, y0)
	c.LineTo(x1, y3)

	var (
		rd   = Base3Bal()
		_, d = rd.RestDigit(digit)
	)

	// Positive
	if (d == 0) || (d == +1) {
		c.MoveTo(x1, y0)
		c.LineTo(x0, y1)
	}

	// Negative
	if (d == 0) || (d == -1) {
		c.MoveTo(x1, y3)
		c.LineTo(x2, y2)
	}

	c.Stroke()
}
