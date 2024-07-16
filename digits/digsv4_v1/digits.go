package main

import (
	"image/color"
	"math"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils/digits"
)

// MeasureString
func digitsBounds(size float64, numberOfDigits int) (width, height float64) {

	var (
		nx int = 2 * (numberOfDigits + 1)
		ny int = 8

		dx float64 = size / float64(ny)
		//dy float64 = size / float64(ny)
	)

	width = float64(nx) * dx
	//height = float64(ny) * dy
	height = size

	return
}

// dh - digit Height

func drawDigits(c *gg.Context, x0, y0 float64, size float64, digits []int) error {

	numberOfDigits := len(digits)

	var (
		nx int = 2 * (numberOfDigits + 1)
		ny int = 8

		dx float64 = size / float64(ny)
		dy float64 = size / float64(ny)
	)

	var (
		greedWidth = 0.02 * minFloat64(dx, dy)
		lineWidth  = 0.3 * minFloat64(dx, dy)
	)

	c.NewSubPath()
	c.Translate(x0, y0)
	c.Scale(dx, dy)

	dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	// Draw digits
	{
		c.SetLineWidth(lineWidth)
		for i, digit := range digits {

			var negative bool
			if digit < 0 {
				negative = true
				digit = -digit
			}

			var (
				p0 = gg.Point{
					X: 0.5 + float64(i)*2,
					Y: 2,
				}
				p1 = gg.Point{
					X: p0.X + 2,
					Y: 0.5,
				}
				p2 = gg.Point{
					X: p0.X + 2,
					Y: 2,
				}
			)

			c.MoveTo(p0.X, p0.Y)
			c.QuadraticTo(p1.X, p1.Y, p2.X, p2.Y)
			c.LineTo(p2.X, 7)
			c.Stroke()

			// c.DrawCircle(p0.X, p0.Y, 0.2)
			// c.DrawCircle(p1.X, p1.Y, 0.2)
			// c.DrawCircle(p2.X, p2.Y, 0.2)
			// c.SetRGB(0, 0.6, 0)
			// c.Fill()

			y0 := p2.Y + 1.5
			dy := 0.7

			ddy := 0.25

			for j := 0; j < digit; j++ {
				y01 := y0 + dy*float64(j)
				c.MoveTo(p1.X-0.5, y01+ddy)
				c.LineTo(p1.X+0.5, y01-ddy)
			}
			c.Stroke()

			if negative {
				c.DrawCircle(p0.X+1.3, p0.Y, 0.3)
				c.Fill()
			}
		}
	}

	c.ClosePath()

	return nil
}

func minFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

//------------------------------------------------------------------------------

func numberToDigits(x int) []int {
	const (
		min = -4
		max = 4
	)
	ds, _ := digits.CalcDigitsN(x, min, max, 10)
	return ds
}

func MakeNumberImage(filename string, number int, size float64) error {

	digits := numberToDigits(number)

	var (
		fW, fH        = digitsBounds(size, len(digits))
		width, height = ceilInt(fW), ceilInt(fH)
	)

	c := gg.NewContext(width, height)

	// fill background
	c.SetColor(color.White)
	c.Clear()

	c.SetColor(color.Black)
	//c.SetRGB(0, 0, 1)

	var (
		digitHeight = size
		fontSize    = digitHeight * 0.08
	)
	setFont(c, fontSize)
	c.DrawString(strconv.Itoa(number), 0, fontSize)

	var x0, y0 float64 = 0, 0
	err := drawDigits(c, x0, y0, size, digits)
	if err != nil {
		return err
	}

	return c.SavePNG(filename)
}

func ceilInt(x float64) int {
	return int(math.Ceil(x))
}

func setFont(c *gg.Context, fontSize float64) error {

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	c.SetFontFace(face)

	return nil
}
