package dig11

import (
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit11_v6 struct{}

var _ dgdr.DigitDrawer = Digit11_v6{}

func (Digit11_v6) Width(height float64) (width float64) {
	width = height * 8 / 24
	return
}

func (d Digit11_v6) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 8
		ny = 24

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.5 * w
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var dx, dy float64
	dx, dy = 1, 1

	radius := dx * math.Sqrt2

	var (
		drawT = func(x, y float64) {
			c.MoveTo(x, y)
			c.LineTo(x+2*dx, y+2*dy)
			c.MoveTo(x-dx, y+3*dy)
			c.DrawArc(x, y+2*dy, radius, DegToRad(135), DegToRad(315))
			c.MoveTo(x-1*dx, y+3*dy)
			c.LineTo(x, y+4*dy)
		}
		draw1 = func(x, y float64) {
			c.MoveTo(x, y)
			c.LineTo(x+dx, y+dy)
			c.DrawArc(x, y+2*dy, radius, DegToRad(-45), DegToRad(135))
			c.MoveTo(x-2*dx, y+2*dy)
			c.LineTo(x, y+4*dy)
		}

		draw0 = func(x, y float64) {
			drawT(x, y)
			draw1(x, y)
		}
	)

	bs := make([]int, 4)
	calcDigits(digit, bs)
	digits := bs

	digits = trimLast0(digits)

	var (
		x0 = 4.0
		y0 = 4.0
	)
	//c.MoveTo(x0-1*dx, y0-1*dy)
	c.MoveTo(x0-2*dx, y0-2*dy)
	c.LineTo(x0, y0)
	for _, digit := range digits {
		switch digit {
		case -1:
			drawT(x0, y0)
		case 0:
			draw0(x0, y0)
		case 1:
			draw1(x0, y0)
		}
		y0 += 4 * dy
	}
	//c.LineTo(x0+1*dx, y0+1*dy)
	c.LineTo(x0+2*dx, y0+2*dy)

	c.Stroke()
}
