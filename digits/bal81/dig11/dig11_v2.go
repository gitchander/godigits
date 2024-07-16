package dig11

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit11_v2 struct{}

var _ dgdr.DigitDrawer = Digit11_v2{}

func (Digit11_v2) Width(height float64) (width float64) {
	width = height * 18 / 50
	return
}

func (d Digit11_v2) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 18
		ny = 50

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 1.0 * w
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var dx, dy float64
	dx, dy = 1, 1

	var (
		drawT = func(x, y float64) {
			c.MoveTo(x, y)
			c.LineTo(x+3*dx, y+3*dy)
			c.CubicTo(x-1*dx, y-1*dy, x-5*dx, y+3*dy, x-1*dx, y+7*dy)
			c.LineTo(x, y+8*dy)
		}

		draw1 = func(x, y float64) {
			c.MoveTo(x, y)
			c.LineTo(x+1*dx, y+1*dy)
			c.CubicTo(x+5*dx, y+5*dy, x+1*dx, y+9*dy, x-3*dx, y+5*dy)
			c.LineTo(x, y+8*dy)
		}

		draw0 = func(x, y float64) {
			drawT(x, y)
			draw1(x, y)
		}
	)

	bs := make([]int, 4)
	calcDigits(digit, bs)
	digits := bs

	//digits = []int{0, 0, -1, 1}
	digits = trimLast0(digits)

	var (
		x0 = 9.0
		y0 = 9.0
	)
	//c.MoveTo(x0-1*dx, y0-1*dy)
	c.MoveTo(x0-3*dx, y0-3*dy)
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
		y0 += 8 * dy
	}
	//c.LineTo(x0+1*dx, y0+1*dy)
	c.LineTo(x0+3*dx, y0+3*dy)

	c.Stroke()
}
