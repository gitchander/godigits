package base16

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit4 struct{}

var _ dgdr.DigitDrawer = Digit4{}

func (Digit4) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit4) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 4
		ny = 8

		dx = digitHeight / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.5 * dx

		circleRadius = 0.35
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var (
		x1 = 1.0
		x2 = 2.0
		x3 = 3.0
	)

	var (
		y1  = 1.0
		y12 = 3.0
		y2  = 3.0
		yl  = 4.0
		y3  = 7.0

		cy1 = 1.0
		cy2 = 3.0
	)

	switch 1 {
	case 1:
		y1 = 1.5
		y12 = 2.0
		y2 = 3.0
		yl = 4.0
		y3 = 7.0

		cy1 = 1.0
		cy2 = 2.8

	case 2:
		y1 = 1.0
		y12 = 2.5
		y2 = 2.5
		yl = 4.0
		y3 = 7.0

		cy1 = 1.0
		cy2 = 2.5

	case 3:
		y1 = 1.0
		y12 = 2.0
		y2 = 2.0
		yl = 4.0
		y3 = 7.0

		cy1 = 1.0
		cy2 = 2.0
	}

	if true {
		c.MoveTo(x1, y1)
		c.QuadraticTo(x2, y12, x2, y2)

		c.MoveTo(x3, y1)
		c.QuadraticTo(x2, y12, x2, y2)

		c.MoveTo(x2, y2)
		c.LineTo(x2, y3)

		c.Stroke()
	}

	quoDigit, remDigit := quoRem(digit, 4)

	if (remDigit == 1) || (remDigit == 3) {
		c.DrawCircle(x2, cy1, circleRadius)
	}
	if (remDigit == 2) || (remDigit == 3) {
		c.DrawCircle(x1, cy2, circleRadius)
		c.DrawCircle(x3, cy2, circleRadius)
	}
	c.Fill()

	for i := 0; i < quoDigit; i++ {
		yp := yl + float64(i)
		c.MoveTo(x1, yp)
		c.LineTo(x3, yp)
	}
	c.Stroke()
}
