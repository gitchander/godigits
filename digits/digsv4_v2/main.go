package main

import (
	"log"
	"path/filepath"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
)

func main() {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	ds := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4}

	var dd dgdr.DigitDrawer

	//dd = Digit1{}
	dd = Digit2{}
	//dd = Digit3{}

	filename := filepath.Join(dirName, "digits.png")
	err := dgdr.MakeDigitsImage(filename, dd, 128, ds)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//------------------------------------------------------------------------------

type Digit1 struct{}

var _ dgdr.DigitDrawer = Digit1{}

func (Digit1) Width(height float64) (width float64) {
	return height / 2
}

func (d Digit1) DrawDigit(c *gg.Context, x, y float64, height float64, digit int) {

	var (
		nx = 4
		ny = 8

		width = d.Width(height)

		dx = width / float64(nx)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.3 * dx

		//circleRadius = 0.60
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	// Draw anchor
	{
		var x1, x2 float64 = 2, 3.5

		var (
			//yl1, yl2, yl3 float64 = 0.5, 6, 7.5
			yl1, yl2, yl3 float64 = 1, 5.5, 7
		)

		c.MoveTo(x1, yl1)
		c.LineTo(x1, yl2)
		c.QuadraticTo(x1, yl3, x2, yl3)
		c.Stroke()
	}

	// Draw digit
	{
		var (
			x1       float64 = 2
			xd1, xd2 float64 = 0.0, 1.5
			//xd1, xd2 float64 = -1.0, 1.0

			yd1, ydd float64 = 2.0, 0.75
		)

		// var negative bool
		if digit < 0 {
			digit = -digit
			xd2 = -xd2
			xd1 = -xd1
		}

		if digit > 0 {
			c.MoveTo(x1+xd2, yd1+0*ydd)
			c.LineTo(x1+xd1, yd1+1*ydd)
		}
		if digit > 1 {
			c.LineTo(x1+xd2, yd1+2*ydd)
		}
		if digit > 2 {
			c.LineTo(x1+xd1, yd1+3*ydd)
		}
		if digit > 3 {
			c.LineTo(x1+xd2, yd1+4*ydd)
		}

		c.Stroke()
	}
}

//------------------------------------------------------------------------------

type Digit2 struct{}

var _ dgdr.DigitDrawer = Digit2{}

func (Digit2) Width(height float64) (width float64) {
	return height / 2
}

func (d Digit2) DrawDigit(c *gg.Context, x, y float64, height float64, digit int) {

	var (
		nx = 4
		ny = 8

		width = d.Width(height)

		dx = width / float64(nx)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.3 * dx

		circleRadius = 0.3
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	// Draw anchor
	{
		// var x1, x2 float64 = 2, 3.5

		// var (
		// 	//yl1, yl2, yl3 float64 = 0.5, 6, 7.5
		// 	yl1, yl2, yl3 float64 = 1, 5.5, 7
		// )

		c.MoveTo(1, 1.5)
		c.QuadraticTo(3, 0, 3, 2)
		c.LineTo(3, 7)

		c.Stroke()
	}

	// Draw digit
	{
		var (
			x1 float64 = 1

			//-----------------------------------------
			//xd1, xd2 float64 = 0.5, 2.0
			//xd1, xd2 float64 = -1.0, 1.0
			//-----------------------------------------
			xd1, xd2 float64 = 1.0, 2.0
			yd1, ydd float64 = 3.0, 0.75
			//-----------------------------------------
			// xd1, xd2 float64 = 0.5, 2.0
			// yd1, ydd float64 = 3.0, 0.5
			//-----------------------------------------
		)

		var negative bool
		if digit < 0 {
			digit = -digit
			negative = true
		}

		if digit > 0 {
			c.MoveTo(x1+xd2, yd1+0*ydd)
			c.LineTo(x1+xd1, yd1+1*ydd)
		}
		if digit > 1 {
			c.LineTo(x1+xd2, yd1+2*ydd)
		}
		if digit > 2 {
			c.LineTo(x1+xd1, yd1+3*ydd)
		}
		if digit > 3 {
			c.LineTo(x1+xd2, yd1+4*ydd)
		}

		c.Stroke()

		if negative {
			c.DrawCircle(2, 2, circleRadius)
			c.Fill()
		}
	}
}

//------------------------------------------------------------------------------

type Digit3 struct{}

var _ dgdr.DigitDrawer = Digit3{}

func (Digit3) Width(height float64) (width float64) {
	return height / 2
}

func (d Digit3) DrawDigit(c *gg.Context, x, y float64, height float64, digit int) {

	var (
		nx = 8
		ny = 16

		width = d.Width(height)

		dx = width / float64(nx)
		dy = dx

		circleRadius = 0.5

		greedWidth = 0.02 * dx
		lineWidth  = circleRadius * dx
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	// Draw anchor
	{
		c.MoveTo(1, 4)

		if true {
			c.QuadraticTo(5, 1, 5, 4)
		} else {
			c.CubicTo(3, 2, 5, 2, 5, 4)
		}

		c.LineTo(5, 14)

		c.Stroke()
	}

	// Draw digit
	{
		var (
			x1 float64 = 5
			//xd1, xd2 float64 = -0.75, 0.75
			//xd1, xd2 float64 = -1.0, 0.0
			xd1, xd2 float64 = -1, 1

			yd1, ydd float64 = 7, (4.0 / 3.0) // 0.666
			//yd1, ydd float64 = 3.0, 1.0
			//yd1, ydd float64 = 3.5, 2.0 / 3.0
		)

		var negative bool
		if digit < 0 {
			digit = -digit
			negative = true
		}

		for i := 0; i < digit; i++ {
			y := yd1 + float64(i)*ydd
			c.MoveTo(x1+xd1, y)
			c.LineTo(x1+xd2, y)
		}
		c.Stroke()

		if negative {
			c.DrawCircle(3.75, 4, circleRadius)
			c.Fill()
		}
	}
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
