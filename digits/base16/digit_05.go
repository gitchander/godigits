package base16

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

// The digits that was created my little son Sasha.
// He also wanted to come up with something. :)

type Digit5 struct{}

var _ dgdr.DigitDrawer = Digit5{}

func (Digit5) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit5) DrawDigit(c *gg.Context, x0, y0 float64, digitHeight float64, digit int) {

	//func drawDigitSasha(c *gg.Context, x0, y0 float64, size float64, digit int) error {

	var (
		nx = 6
		ny = 12

		dx = digitHeight / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.5 * dx

		circleRadius = 0.5
	)

	c.Push()
	defer c.Pop()

	c.Translate(x0, y0)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var (
		x1 = 1.0
		x2 = 3.0
		x3 = 5.0
	)

	var (
		y1 = 2.5
		y2 = 4.0
		yl = 6.0
		y3 = 10.0

		cy1 = 2.0
		cy2 = 4.0
	)

	// switch 0 {
	// case 1:
	// 	y1 = 2.5
	// 	y2 = 4.0
	// 	yl = 6.0
	// 	y3 = 10.0

	// 	cy1 = 2.0
	// 	cy2 = 4.0
	// }

	if true {
		c.MoveTo(x1, y2)
		c.CubicTo(x1, y1, x2, y1, x2, y2)
		//c.QuadraticTo(x2, y12, x2, y2)

		c.MoveTo(x3, y2)
		c.CubicTo(x3, y1, x2, y1, x2, y2)
		//c.QuadraticTo(x2, y12, x2, y2)

		c.MoveTo(x2, y2)
		c.LineTo(x2, y3)

		c.Stroke()
	}

	quoDigit, remDigit := quoRem(digit, 4)

	if (remDigit == 1) || (remDigit == 3) {
		c.DrawCircle(x2, cy1, circleRadius)
	}
	if (remDigit == 2) || (remDigit == 3) {
		c.DrawCircle(x2-1, cy2, circleRadius)
		c.DrawCircle(x2+1, cy2, circleRadius)
	}
	c.Fill()

	for i := 0; i < quoDigit; i++ {
		yp := yl + float64(i)
		c.MoveTo(x1, yp)
		c.LineTo(x3, yp)
	}
	c.Stroke()
}
