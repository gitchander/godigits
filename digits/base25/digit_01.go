package base25

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit1 struct{}

var _ dgdr.DigitDrawer = Digit1{}

func (Digit1) Width(height float64) (width float64) {
	width = height / 2
	return width
}

func (Digit1) DrawDigit(c *gg.Context, x0, y0 float64, height float64, digit int) {

	var (
		nx = 8
		ny = 16

		w = height / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.8 * w

		circleRadius = 0.6
		//circleRadius = 0.5
	)

	c.Push()
	defer c.Pop()

	c.Translate(x0, y0)
	c.Scale(w, w)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var (
		x1 = 2.0
		x2 = 4.0
		x3 = 6.0
	)

	var (
		y1 = 2.0
		y2 = 4.0
		y3 = 6.0
		y4 = 8.0
		y5 = 14.0

		yl  = 8.0
		ydl = 1.5
	)

	{
		c.MoveTo(x1, y1)
		c.LineTo(x2, y2)
		c.QuadraticTo(x3, y3, x3, y4)
		c.LineTo(x3, y5)

		c.MoveTo(x3, y1)
		c.LineTo(x2, y2)
		c.QuadraticTo(x1, y3, x1, y4)
		c.LineTo(x1, y5)

		c.Stroke()
	}

	quoDigit, remDigit := quoRem(digit, 5)

	if (remDigit == 1) || (remDigit == 2) || (remDigit == 3) || (remDigit == 4) {
		c.DrawCircle(x2, y1, circleRadius)
	}
	if (remDigit == 3) || (remDigit == 4) {
		c.DrawCircle(x1, y2, circleRadius)
		c.DrawCircle(x3, y2, circleRadius)
	}
	if (remDigit == 2) || (remDigit == 4) {
		c.DrawCircle(x2, y3, circleRadius)
	}

	c.Fill()

	for i := 0; i < quoDigit; i++ {
		yp := yl + ydl*float64(i)
		c.MoveTo(x1+1, yp)
		c.LineTo(x3-1, yp)
	}
	c.Stroke()
}
