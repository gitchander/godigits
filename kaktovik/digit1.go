package main

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
)

type Digit1 struct{}

var _ dgdr.DigitDrawer = Digit1{}

func (Digit1) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit1) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		// nx = 4
		// ny = 8

		nx = 6
		ny = 12

		dx = digitHeight / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.45 * dx
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var d1, d2 int

	digit, d1 = utils.QuoRem(digit, 5)
	digit, d2 = utils.QuoRem(digit, 5)

	switch d2 {
	case 1:
		{
			var (
				y0 = 5.0
				//dy = -1.0
				dy = -2.0

				x1, x2 float64 = 1, 5
			)
			c.MoveTo(x1, y0+(0*dy))
			c.LineTo(x2, y0+(1*dy))
		}
	case 2:
		{
			var (
				y0 = 5.0
				dy = -1.5
				//dy = -2.0

				x1, x2 float64 = 1, 5
			)
			c.MoveTo(x1, y0+(0*dy))
			c.LineTo(x2, y0+(1*dy))
			c.LineTo(x1, y0+(2*dy))
		}
	case 3:
		{
			var (
				y0 = 5.0
				dy = -1.0
				//dy = -4.0 / 3.0

				x1, x2 float64 = 1, 5
			)
			c.MoveTo(x1, y0+(0*dy))
			c.LineTo(x2, y0+(1*dy))
			c.LineTo(x1, y0+(2*dy))
			c.LineTo(x2, y0+(3*dy))
		}
	}

	switch d1 {
	case 1:
		{
			var (
				x0 = 1.0
				dx = 2.0

				y1, y2 float64 = 5, 9
			)
			c.MoveTo(x0+(0*dx), y1)
			c.LineTo(x0+(1*dx), y2)
		}
	case 2:
		{
			var (
				x0 = 1.0
				dx = 2.0

				y1, y2 float64 = 5, 9
			)
			c.MoveTo(x0+(0*dx), y1)
			c.LineTo(x0+(1*dx), y2)
			c.LineTo(x0+(2*dx), y1)
		}
	case 3:
		{
			var (
				x0 = 1.0
				dx = 4.0 / 3.0

				y1, y2 float64 = 5, 9
			)
			c.MoveTo(x0+(0*dx), y1)
			c.LineTo(x0+(1*dx), y2)
			c.LineTo(x0+(2*dx), y1)
			c.LineTo(x0+(3*dx), y2)
		}
	case 4:
		{
			var (
				x0 = 1.0
				dx = 1.0

				y1, y2 float64 = 5, 9
			)
			c.MoveTo(x0+(0*dx), y1)
			c.LineTo(x0+(1*dx), y2)
			c.LineTo(x0+(2*dx), y1)
			c.LineTo(x0+(3*dx), y2)
			c.LineTo(x0+(4*dx), y1)
		}
	}

	c.Stroke()
}
