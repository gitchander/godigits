package base27

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawer2 struct{}

func (dd DigitDrawer2) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {
	//dd.drawDigit_v1(c, b, digit)
	dd.drawDigit_v2(c, b, digit)
}

// func (DigitDrawer2) drawDigit_v1(c *gg.Context, b geom.Bounds, digit int) {

// 	b = geom.BoundsAspect(b, AspectRatio)

// 	const (
// 		koefDA = 0.4
// 		//koefDA = 0.5

// 		koefLineWidth = 0.75
// 		//koefLineWidth = 0.8
// 		//koefLineWidth = 0.9
// 		//koefLineWidth = 1.0

// 		//enableAll = true
// 		enableAll = false
// 	)

// 	var (
// 		w = min(b.Dx(), b.Dy()/2)

// 		dA = w / 2 * koefDA
// 		dB = (w/2 - dA) / 2

// 		xA = dA
// 		xB = dB

// 		yA = dA
// 		yB = dB

// 		lineWidth = dB * koefLineWidth
// 	)

// 	c.SetLineCap(gg.LineCapRound)
// 	c.SetLineWidth(lineWidth)

// 	center := b.Center()

// 	// Draw vertical lines
// 	if true {
// 		c.MoveTo(center.X, center.Y-(2*yA+yB+yB/2))
// 		c.LineTo(center.X, center.Y-(yB/2))

// 		c.MoveTo(center.X, center.Y+(yB/2))
// 		c.LineTo(center.X, center.Y+(2*yA+yB+yB/2))
// 	}

// 	trits := calcTritsBase27(digit)

// 	var (
// 		x1 = xB
// 		x2 = xA + xB

// 		y0 = center.Y - (2*yA + yB + yB/2)
// 		y  = y0
// 		dy = yA + yB
// 	)
// 	for i, trit := range trits {
// 		if (i % 2) == 0 {
// 			if enableAll || (trit == -1) {
// 				c.MoveTo(center.X-x1, y+0*yA)
// 				c.LineTo(center.X-x2, y+1*yA)
// 			}
// 			if enableAll || (trit == +1) {
// 				c.MoveTo(center.X+x1, y+0*yA)
// 				c.LineTo(center.X+x2, y+1*yA)
// 			}
// 		} else {
// 			if enableAll || (trit == -1) {
// 				c.MoveTo(center.X-x2, y+0*yA)
// 				c.LineTo(center.X-x1, y+1*yA)
// 			}
// 			if enableAll || (trit == +1) {
// 				c.MoveTo(center.X+x2, y+0*yA)
// 				c.LineTo(center.X+x1, y+1*yA)
// 			}
// 		}
// 		y += dy
// 	}

// 	c.Stroke()
// }

func (DigitDrawer2) drawDigit_v2(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)

	v := b.Vmin()

	c.Push()
	defer c.Pop()

	c.Translate(b.Min.X, b.Min.Y)
	c.Scale(v, v)

	// size 100x200 : dx = 100, dy = 200

	var (
		// dA              = 20.0
		// dB              = 20.0
		// lineWidthFactor = 1.0

		// dA            = 20.0
		// dB            = 15.0
		// lineWidthFactor = 0.8

		dA              = 20.0
		dB              = (50 - dA) / 2
		lineWidthFactor = 0.8
	)

	var (
		xA = dA
		xB = dB

		yA = dA
		yB = dB

		lineWidth = dB * lineWidthFactor

		//enableAll = true
		enableAll = false
	)

	c.SetLineCap(gg.LineCapRound)
	c.SetLineWidth(lineWidth * v)

	center := geom.MakePoint2f(50, 100)

	yH := 2*yA + yB + yB/2

	// Draw vertical lines
	if true {
		c.MoveTo(center.X, center.Y-yH)
		c.LineTo(center.X, center.Y-(yB/2))

		c.MoveTo(center.X, center.Y+(yB/2))
		c.LineTo(center.X, center.Y+yH)
	}

	trits := calcTritsBase27(digit)

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
