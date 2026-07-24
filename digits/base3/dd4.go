package base3

import (
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

const (
	pi  = math.Pi
	tau = 2 * math.Pi
)

type DigitDrawer4 struct{}

func (DigitDrawer4) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, DigitAspectRatio)
	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	var (
		x1 = 20.0
		x2 = 50.0
		x3 = 80.0

		dy = 30.0
	)

	lw := 10.0
	c.SetLineWidth(lw * v)
	c.SetLineCap(gg.LineCapRound)
	c.SetRGB(0, 0, 0)

	c.MoveTo(x2, 20)
	c.QuadraticTo(x1, 100-dy, 20, 100)
	// c.QuadraticTo(x1, 100+dy, x2, 100+dy)
	// c.QuadraticTo(x3, 100+dy, x3, 100)
	c.MoveTo(x3, 100)
	c.QuadraticTo(x3, 100+dy, x2, 180)

	c.Stroke()

	d := digit
	if d == 0 {
		c.DrawArc(x2, 100, 30, 0, tau)
		//c.DrawCircle(x2, 100, 30)
		c.Stroke()
	}

	if d == -1 {
		c.DrawArc(x2, 100, 30, 0, pi)
		c.Stroke()
	}

	if d == +1 {
		c.DrawArc(x2, 100, 30, pi, tau)
		c.Stroke()
	}
}
