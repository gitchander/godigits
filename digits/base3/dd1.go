package base3

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawer1 struct{}

func (DigitDrawer1) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)
	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	lw := 10.0
	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)
	c.SetRGB(0, 0, 0)

	c.MoveTo(50, 20)
	c.LineTo(50, 180)

	d := digit
	var (
		drawAll = false
		// drawAll = (d == 0)
	)

	// Positive
	if drawAll || (d == +1) {
		c.MoveTo(50, 20)
		c.LineTo(10, 60)
	}

	// Negative
	if drawAll || (d == -1) {
		c.MoveTo(50, 180)
		c.LineTo(90, 140)
	}

	c.Stroke()
}
