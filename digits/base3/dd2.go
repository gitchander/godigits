package base3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawer2 struct{}

func (DigitDrawer2) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	var (
		lw = 10.0
	)

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)
	c.SetRGB(0, 0, 0)

	c.MoveTo(10, 50)
	c.QuadraticTo(50, 20, 50, 50)
	c.LineTo(50, 170)

	var (
		rd   = Base3Bal()
		_, d = rd.RestDigit(digit)
	)

	// Negative or Positive
	if (d == -1) || (d == +1) {
		c.MoveTo(30, 100)
		c.LineTo(70, 100)
	}

	c.Stroke()

	// Negative
	if d == -1 {
		c.DrawCircle(31, 60, 8)
		c.Fill()
	}
}
