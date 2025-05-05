package dgdr

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

//------------------------------------------------------------------------------

// Width - ширина (dx)
// Height - висота (dy)

type DigitDrawer interface {
	// Digit Width
	Width(digitHeight float64) (digitWidth float64)

	DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int)
}

func MakeDigitImage(filename string, dd DigitDrawer, digitHeight float64, digit int) error {

	var (
		dx = dd.Width(digitHeight)
		dy = digitHeight

		cw = ceilInt(dx)
		ch = ceilInt(dy)
	)

	c := gg.NewContext(cw, ch)

	// fill background
	c.SetColor(color.White)
	c.Clear()

	c.SetColor(color.Black)

	dd.DrawDigit(c, 0, 0, dy, digit)

	return c.SavePNG(filename)
}

func MakeDigitsImage(filename string, dd DigitDrawer, digitHeight float64, digits []int) error {

	var (
		dx = dd.Width(digitHeight)
		dy = digitHeight

		cw = ceilInt(dx * float64(len(digits)))
		ch = ceilInt(dy)
	)

	c := gg.NewContext(cw, ch)

	var (
		bc = color.White
		fc = color.Black

		// bc = color.Black
		// fc = color.White
	)

	// fill background
	c.SetColor(bc)
	c.Clear()

	c.SetColor(fc)

	drawDigits(c, dd, 0, 0, dx, dy, digits)

	return c.SavePNG(filename)
}

func MakeDigitsImageMatrix(filename string, dd DigitDrawer, xn, yn int, digitHeight float64, digits []int) error {

	var (
		dx = dd.Width(digitHeight)
		dy = digitHeight

		cw = ceilInt(dx * float64(xn))
		ch = ceilInt(dy * float64(yn))
	)

	c := gg.NewContext(cw, ch)

	fontSize := digitHeight * 0.08
	err := SetFontSizeGG(c, fontSize)
	if err != nil {
		return err
	}

	// fill background
	c.SetColor(color.White)
	c.Clear()

	c.SetColor(color.Black)

	c.SetLineCap(gg.LineCapRound)
	//c.SetLineCap(gg.LineCapButt)
	//c.SetLineCap(gg.LineCapSquare)

	c.SetLineJoin(gg.LineJoinRound)
	//c.SetLineJoin(gg.LineJoinBevel)

	ddt := digitDecText{
		enable: true,
		x:      0,
		y:      fontSize,
	}

	drawMatrix(c, dd, xn, yn, 0, 0, dx, dy, ddt, digits)

	return c.SavePNG(filename)
}

type digitDecText struct {
	enable bool
	x, y   float64
}

func SetFontSizeGG(c *gg.Context, fontSize float64) error {

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	c.SetFontFace(face)

	return nil
}

func drawDigits(c *gg.Context, dd DigitDrawer, x0, y0 float64, dx, dy float64, digits []int) {
	for _, digit := range digits {
		dd.DrawDigit(c, x0, y0, dy, digit)
		x0 += dx
	}
}

func drawMatrix(c *gg.Context, dd DigitDrawer, xn, yn int, x0, y0 float64,
	dx, dy float64, ddt digitDecText, digits []int) {

	for y := 0; y < yn; y++ {
		for x := 0; x < xn; x++ {
			var (
				x1 = x0 + float64(x)*dx
				y1 = y0 + float64(y)*dy
			)
			if true {
				if ((x + y) % 2) == 0 {
					c.SetRGB(1, 1, 1)
				} else {
					c.SetRGB(0.7, 0.9, 1)
				}
				c.DrawRectangle(x1, y1, dx, dy)
				c.Fill()
				c.SetColor(color.Black)
			}
			if len(digits) > 0 {
				digit := digits[0]
				digits = digits[1:]

				if ddt.enable {
					c.DrawString(formatInt(digit), x1+ddt.x, y1+ddt.y)
				}

				dd.DrawDigit(c, x1, y1, dy, digit)
			}
		}
	}
}

//------------------------------------------------------------------------------
