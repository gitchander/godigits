package base16

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit3 struct{}

func (Digit3) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit3) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 8
		ny = 16

		w = digitHeight / float64(ny)

		relativeSize = 0.5

		greedWidth = 0.025 * w
		lineWidth  = relativeSize * w

		radius = relativeSize
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	//----------------------------------------------------------------------

	var (
		// x1, x2 float64 = 2, 6
		x1, x2 float64 = 3, 5

		//y1, y2 float64 = 2, 14
	)

	var (
		val      = digit
		negative = false
	)

	const maxValue = 15

	if (digit < -maxValue) || (maxValue < digit) {
		return
	}

	val = ((val + maxValue) % (2*maxValue + 1)) - maxValue

	if val < 0 {
		val = -val
		negative = true
	}

	switch k := val % 4; k {

	case 0:
		{
			c.MoveTo(3, 2)
			c.LineTo(3, 14)

			c.MoveTo(5, 2)
			c.LineTo(5, 14)
		}

	case 1:
		{
			c.MoveTo(3, 2)
			c.LineTo(3, 6)
			c.CubicTo(3, 8, 5, 8, 5, 10)
			c.LineTo(5, 14)

			c.MoveTo(5, 2)
			c.LineTo(5, 6)
			c.CubicTo(5, 8, 3, 8, 3, 10)
			c.LineTo(3, 14)
		}

	case 2:
		{
			c.MoveTo(3, 2)
			c.LineTo(3, 4)
			c.CubicTo(3, 6, 5, 6, 5, 8)
			c.CubicTo(5, 10, 3, 10, 3, 12)
			c.LineTo(3, 14)

			c.MoveTo(5, 2)
			c.LineTo(5, 4)
			c.CubicTo(5, 6, 3, 6, 3, 8)
			c.CubicTo(3, 10, 5, 10, 5, 12)
			c.LineTo(5, 14)
		}

	case 3:
		{
			c.MoveTo(3, 2)
			c.CubicTo(3, 4, 5, 4, 5, 6)
			c.CubicTo(5, 8, 3, 8, 3, 10)
			c.CubicTo(3, 12, 5, 12, 5, 14)

			c.MoveTo(5, 2)
			c.CubicTo(5, 4, 3, 4, 3, 6)
			c.CubicTo(3, 8, 5, 8, 5, 10)
			c.CubicTo(5, 12, 3, 12, 3, 14)
		}
	}
	//----------------------------------------------------------------------

	x1 = x1 - 1
	x2 = x2 + 1

	{
		var y1 float64 = 3
		var dy float64 = 1

		bs := parseBools("000-000")

		switch val {
		case 0, 1:
			bs = parseBools("000-000")

		// case 2, 3:
		// 	bs = parseBools("00000-1-00000")
		case 2:
			bs = parseBools("00000-0-00000")
		case 3:
			bs = parseBools("00000-0-00000")

		//---------------------------------------------
		// case 4, 5:
		// 	bs = parseBools("01000-0-00010")
		case 4:
			bs = parseBools("01000-0-00010")
		case 5:
			bs = parseBools("01000-0-00010")

		//---------------------------------------------
		case 6:
			bs = parseBools("01000-0-00010")
		case 7:
			bs = parseBools("01000-0-00010")

		//---------------------------------------------
		// case 8:
		// 	bs = parseBools("01010-0-01010")
		// case 9:
		// 	bs = parseBools("01010-0-01010")
		// case 10:
		// 	bs = parseBools("01010-0-01010")
		// case 11:
		// 	bs = parseBools("01010-0-01010")

		case 8, 9, 10, 11:
			bs = parseBools("01010-0-01010")
			//bs = parseBools("10100-0-00101")

		//---------------------------------------------
		// case 8, 9:
		// 	bs = parseBools("10100-0-00101")
		// case 10, 11:
		// 	bs = parseBools("10100-1-00101")

		//---------------------------------------------
		// case 12, 13, 14, 15:
		// 	bs = parseBools("11100-0-00111")
		case 12, 13, 14, 15:
			//bs = parseBools("01110-0-01110")
			bs = parseBools("11100-0-00111")
			//bs = parseBools("01010-0-01010")
		}

		for i, b := range bs {
			if b {
				y := y1
				y += dy * float64(i)

				c.MoveTo(x1, y)
				c.LineTo(x2, y)
			}
		}
	}
	//----------------------------------------------------------------------

	c.Stroke()

	if negative || val == 0 {

		if (val % 2) != 0 {
			c.DrawCircle(2.5, 8, radius)
			c.DrawCircle(5.5, 8, radius)
		} else {
			c.DrawCircle(4, 8, radius)
		}
		c.Fill()

		// if true {
		// 	c.DrawCircle(2.5, 8, radius)
		// 	c.DrawCircle(5.5, 8, radius)
		// 	c.Fill()
		// } else {
		// 	c.DrawCircle(4, 6, radius)
		// 	c.DrawCircle(4, 10, radius)
		// 	c.Fill()
		// }
	}

	if false {
		c.Push()

		// (5, 12, 4, 12)

		c.SetRGB(1, 0, 0)
		radius := 0.25
		c.DrawCircle(5, 10, radius)
		c.DrawCircle(5, 11, radius)
		c.DrawCircle(4, 12, radius)
		c.Fill()

		c.Pop()
	}
}
