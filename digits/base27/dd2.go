package base27

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawer2 struct{}

func (DigitDrawer2) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	b = geom.BoundsAspect(b, AspectRatio)

	const (
		koefDA = 0.4
		//koefDA = 0.5

		koefLineWidth = 0.75
		//koefLineWidth = 0.8
		//koefLineWidth = 0.9
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

	trits := make([]int, 3)
	calcTrits(digit, trits)

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

// func (DigitDrawer2) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

// 	b = geom.BoundsAspect(b, AspectRatio)
// 	v := b.Vmin()

// 	c.Push()
// 	defer c.Pop()

// 	c.Translate(b.Min.X, b.Min.Y)
// 	c.Scale(v, v)

// 	lw := 10.0
// 	c.SetLineWidth(lw * v)
// 	c.SetLineCap(gg.LineCapRound)
// 	c.SetRGB(0, 0, 0)

// 	bs := make([]int, 3)
// 	calcDigits(digit, bs)

// 	var (
// 		// xs = []float64{10, 50, 90}
// 		// ys = []float64{20, 60, 100, 140, 180}

// 		xs = []float64{15, 50, 85}
// 		ys = []float64{30, 65, 100, 135, 170}
// 	)

// 	drawLineSegment(c, pt(xs[1], ys[0]), pt(xs[1], ys[2]))
// 	drawLineSegment(c, pt(xs[1], ys[2]), pt(xs[1], ys[4]))

// 	switch d := bs[0]; d {
// 	case -1:
// 		drawLineSegment(c, pt(xs[1], ys[0]), pt(xs[0], ys[1]))
// 	case +1:
// 		drawLineSegment(c, pt(xs[1], ys[0]), pt(xs[2], ys[1]))
// 	default:
// 	}

// 	switch d := bs[1]; d {
// 	case -1:
// 		drawLineSegment(c, pt(xs[1], ys[2]), pt(xs[0], ys[1]))
// 	case +1:
// 		drawLineSegment(c, pt(xs[1], ys[2]), pt(xs[2], ys[1]))
// 	default:
// 	}

// 	switch d := bs[2]; d {
// 	case -1:
// 		drawLineSegment(c, pt(xs[1], ys[2]), pt(xs[0], ys[3]))
// 	case +1:
// 		drawLineSegment(c, pt(xs[1], ys[2]), pt(xs[2], ys[3]))
// 	default:
// 	}

// 	c.Stroke()
// }

// func pt(x, y float64) geom.Point2f {
// 	return geom.MakePoint2f(x, y)
// }

// func drawLineSegment(c *gg.Context, a, b geom.Point2f) {
// 	const k = 0.1
// 	var (
// 		a1 = geom.PtLerp(a, b, k)
// 		b1 = geom.PtLerp(a, b, 1-k)
// 	)
// 	c.MoveTo(a1.X, a1.Y)
// 	c.LineTo(b1.X, b1.Y)
// }
