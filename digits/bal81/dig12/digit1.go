package dig12

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
)

type Digit1 struct {
	A float64 // 0 <= A
	B float64 // 0 <= B <= (A/2)
	C float64 // 0 <= C <= (B/2)
}

func (d Digit1) DigitDrawer() dgdr.DigitDrawerB {
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

func MakeDigit1_p1() dgdr.DigitDrawerB {
	dA := 30.0
	return Digit1{
		A: dA,
		B: dA * 0.26,
		C: dA * 0.2,
	}.DigitDrawer()
}

//------------------------------------------------------------------------------

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

	// Draw vertical lines
	if true {
		c.MoveTo(dA, 0*dA+dB)
		c.LineTo(dA, 2*dA-dB)

		c.MoveTo(dA, 2*dA+dB)
		c.LineTo(dA, 4*dA-dB)
	}

	trits := calcTrits(digit, 4)

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
			case +1:
				drawSegment(1, i, (i%2) != 0)
			}
		}
	}

	c.Stroke()
}
