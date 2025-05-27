package base27

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

const AspectRatio = 2

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

	trits := calcTritsBase27(digit)

	var (
		// xs = []float64{10, 50, 90}
		// ys = []float64{20, 60, 100, 140, 180}

		xs = []float64{15, 50, 85}
		ys = []float64{30, 65, 100, 135, 170}
	)

	c.MoveTo(xs[1], ys[0])
	c.LineTo(xs[1], ys[4])

	switch d := trits[0]; d {
	case -1:
		c.MoveTo(xs[1], ys[0])
		c.LineTo(xs[0], ys[1])
	case +1:
		c.MoveTo(xs[1], ys[0])
		c.LineTo(xs[2], ys[1])
	default:
	}

	switch d := trits[1]; d {
	case -1:
		c.MoveTo(xs[1], ys[2])
		c.LineTo(xs[0], ys[1])
	case +1:
		c.MoveTo(xs[1], ys[2])
		c.LineTo(xs[2], ys[1])
	default:
	}

	switch d := trits[2]; d {
	case -1:
		c.MoveTo(xs[1], ys[2])
		c.LineTo(xs[0], ys[3])
	case +1:
		c.MoveTo(xs[1], ys[2])
		c.LineTo(xs[2], ys[3])
	default:
	}

	c.Stroke()
}
