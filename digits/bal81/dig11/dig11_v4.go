package dig11

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit11_v4 struct{}

var _ dgdr.DigitDrawer = Digit11_v4{}

func (Digit11_v4) Width(height float64) (width float64) {
	width = height * 8 / 24
	return
}

func (d Digit11_v4) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

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
	//dx, dy = 0.5, 0.5

	var (
		drawT = func(x, y float64) {
			c.LineTo(x+2*dx, y+2*dy)
			c.CubicTo(x-1*dx, y-1*dy, x-3*dx, y+1*dy, x, y+4*dy)
		}
		draw0 = func(x, y float64) {
			c.MoveTo(x+2*dx, y+2*dy)
			c.CubicTo(x-1*dx, y-1*dy, x-3*dx, y+1*dy, x, y+4*dy)

			c.MoveTo(x, y)
			c.CubicTo(x+3*dx, y+3*dy, x+1*dx, y+5*dy, x-2*dx, y+2*dy)

			c.MoveTo(x, y+4*dy)
		}
		draw1 = func(x, y float64) {
			c.CubicTo(x+3*dx, y+3*dy, x+1*dx, y+5*dy, x-2*dx, y+2*dy)
			c.LineTo(x, y+4*dy)
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
