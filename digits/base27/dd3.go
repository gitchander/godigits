package base27

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawer3 struct{}

func (DigitDrawer3) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

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

	if true {
		c.MoveTo(10, 40)
		c.QuadraticTo(50, 10, 50, 50)
	}
	c.MoveTo(50, 50)
	c.LineTo(50, 160)

	// Draw horizontal lines
	{
		c.MoveTo(20, 80)
		c.LineTo(80, 80)

		c.MoveTo(20, 120)
		c.LineTo(80, 120)
	}

	c.Stroke()

	trits := calcTritsBase27(digit)

	var (
		radius = 8.0

		ys = []float64{60, 100, 140}
	)
	for i, trit := range trits {
		switch trit {
		case -1:
			c.DrawCircle(30, ys[i], radius)
		case +1:
			c.DrawCircle(70, ys[i], radius)
		}
	}

	c.Fill()
}
