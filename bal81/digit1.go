package bal81

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/numbers"
)

type DigitDrawer interface {
	DrawDigit(c *gg.Context, b geom.Bounds, digit int)
}

type Digit1 struct {
	A float64 // 0 <= A
	B float64 // 0 <= B <= (A/2)
	C float64 // 0 <= C <= (B/2)
}

func (d Digit1) DigitDrawer() DigitDrawer {
	var (
		dA = d.A
		dB = clamp(d.B, 0, dA/2)
		dC = clamp(d.C, 0, dB)
	)
	return &digitDrawer1{
		a: dA,
		b: dB,
		c: dC,
	}
}

type digitDrawer1 struct {
	a, b, c float64
}

func (p *digitDrawer1) DrawDigit(c *gg.Context, b geom.Bounds, digit int) {

	var (
		dA = p.a
		dB = p.b
		dC = p.c
	)

	c.Push()
	defer c.Pop()

	c.SetLineCap(gg.LineCapRound)
	c.SetLineWidth(dC)
	c.Translate(b.Min.X, b.Min.Y)

	drawSegment := func(x, y int, rev bool) {
		var (
			x0 = float64(x + 0)
			x1 = float64(x + 1)

			y0 = float64(y + 0)
			y1 = float64(y + 1)
		)
		if rev {
			c.MoveTo(dA*x0+dB, dA*y1-dB)
			c.LineTo(dA*x1-dB, dA*y0+dB)
		} else {
			c.MoveTo(dA*x0+dB, dA*y0+dB)
			c.LineTo(dA*x1-dB, dA*y1-dB)
		}
	}

	// Draw vertical lines
	if true {
		c.MoveTo(dA, 0*dA+dB)
		c.LineTo(dA, 2*dA-dB)

		c.MoveTo(dA, 2*dA+dB)
		c.LineTo(dA, 4*dA-dB)
	}

	trits := calcTrits(digit, 4)

	// Segments:
	if false {
		for i := range trits {
			drawSegment(0, i, (i%2) == 0)
			drawSegment(1, i, (i%2) != 0)
		}
	} else {
		for i, trit := range trits {
			switch trit {
			case -1:
				drawSegment(0, i, (i%2) == 0)
			case 1:
				drawSegment(1, i, (i%2) != 0)
			}
		}
	}

	c.Stroke()
}

func clamp(x float64, min, max float64) float64 {
	if min > max {
		return 0
	}
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func calcTrits(x int, n int) []int {
	b := numbers.Bal3
	trits := make([]int, n)
	for i := range trits {
		x, trits[i] = numbers.RestDigit(b, x)
	}
	return trits
}
