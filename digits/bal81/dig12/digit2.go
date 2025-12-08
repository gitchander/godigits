package dig12

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type Digit2 struct{}

func (dd Digit2) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	const AspectRatio = 2
	b = geom.BoundsAspect(b, AspectRatio)

	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	var (
		digitSize = geom.MakePoint2f(100, 200)
		center    = digitSize.DivScalar(2)
	)

	var (
		// dA        = 20.0
		// dB        = 20.0
		// lineWidth = 20.0

		// dA        = 20.0
		// dB        = (50 - dA) / 2
		// lineWidth = dB * 0.8

		dA        = 20.0
		dB        = 15.0
		lineWidth = 12.0
	)

	var (
		xA = dA
		xB = dB

		yA = dA
		yB = dB
	)

	c.SetLineCap(gg.LineCapRound)
	c.SetLineWidth(lineWidth * v) // (* v) - because LineWidth is not scaling.

	yH := 2*yA + yB + yB/2

	// Draw vertical lines
	if true {
		c.MoveTo(center.X, center.Y-yH)
		c.LineTo(center.X, center.Y-(yB/2))

		c.MoveTo(center.X, center.Y+(yB/2))
		c.LineTo(center.X, center.Y+yH)
	}

	var (
		// enableAll = true
		// trits     = []int{0, 0, 0, 0}

		enableAll = false
		trits     = calcTritsBal81(digit)
	)

	var (
		x1 = xB
		x2 = xA + xB

		y0 = center.Y - yH
		y  = y0
		dy = yA + yB
	)
	for i, trit := range trits {
		if (i % 2) == 0 {
			if enableAll || (trit == -1) {
				c.MoveTo(center.X-x1, y+0*yA)
				c.LineTo(center.X-x2, y+1*yA)
			}
			if enableAll || (trit == +1) {
				c.MoveTo(center.X+x1, y+0*yA)
				c.LineTo(center.X+x2, y+1*yA)
			}
		} else {
			if enableAll || (trit == -1) {
				c.MoveTo(center.X-x2, y+0*yA)
				c.LineTo(center.X-x1, y+1*yA)
			}
			if enableAll || (trit == +1) {
				c.MoveTo(center.X+x2, y+0*yA)
				c.LineTo(center.X+x1, y+1*yA)
			}
		}
		y += dy
	}

	c.Stroke()
}
