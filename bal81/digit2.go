package bal81

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type Digit2 struct{}

func (Digit2) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	const (
		//koefDA = 0.4
		koefDA = 0.5

		//koefLineWidth = 0.8
		koefLineWidth = 0.9
		//koefLineWidth = 1.0

		//enableAll = true
		enableAll = false
	)

	var (
		w = min(b.Dx(), b.Dy()/2)

		dA = w / 2 * koefDA
		dB = (w/2 - dA) / 2

		xA = dA
		xB = dB

		yA = dA
		yB = dB

		lineWidth = dB * koefLineWidth
	)

	// c.Push()
	// defer c.Pop()

	c.SetLineCap(gg.LineCapRound)
	c.SetLineWidth(lineWidth)

	center := b.Center()

	// Draw vertical lines
	if true {
		c.MoveTo(center.X, center.Y-(2*yA+yB+yB/2))
		c.LineTo(center.X, center.Y-(yB/2))

		c.MoveTo(center.X, center.Y+(yB/2))
		c.LineTo(center.X, center.Y+(2*yA+yB+yB/2))
	}

	trits := calcTrits(digit, 4)

	var (
		x1 = xB
		x2 = xA + xB

		y0 = center.Y - (2*yA + yB + yB/2)
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
