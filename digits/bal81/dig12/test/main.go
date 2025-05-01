package main

import (
	"fmt"
	"image"
	"log"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gitchander/godigits/digits/bal81/dig12"
	"github.com/gitchander/godigits/geom"
)

func main() {
	makeDigit2()
}

func makeDigit2() {

	var (
		sizeX = 60
		sizeY = sizeX * 2
	)
	digitSize := image.Pt(sizeX, sizeY)

	// dA := 30.0
	// d := dig12.Digit1{
	// 	A: dA,
	// 	B: dA * 0.2,
	// 	C: dA * 0.2,
	// }.DigitDrawer()

	d := dig12.Digit2{}

	//digits := []int{40, 40}
	//digits := []int{-3, -2, -1, 0, 1, 2, 3}
	//digits := serialInts(7)
	//digits := intervalInts(-6, 6)
	digits := intervalInts(-40, 40)

	var (
		nX = 9
		nY = 9
	)
	c := gg.NewContext(digitSize.X*nX, digitSize.Y*nY)

	if true {
		c.SetRGB(1, 1, 1)
		c.Clear()
	}

	fontSize := float64(digitSize.Y) * 0.1
	err := setFont(c, fontSize)
	checkError(err)

	drawMatrix(c, d, nX, nY, digitSize, digits)

	err = c.SavePNG("digit2_matrix.png")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

func intervalInts(min, max int) []int {
	n := max - min + 1
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = min + i
	}
	return a
}

func drawBounds(c *gg.Context, b geom.Bounds) {
	c.DrawRectangle(b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
}

func drawMatrix(c *gg.Context, d dig12.DigitDrawer, nX, nY int,
	digitSize image.Point, digits []int) {
	for y := 0; y < nY; y++ {
		for x := 0; x < nX; x++ {
			b := geom.MakeBounds(
				float64((x+0)*digitSize.X), float64((y+0)*digitSize.Y),
				float64((x+1)*digitSize.X), float64((y+1)*digitSize.Y),
			)
			if true {
				if ((x + y) % 2) == 0 {
					c.SetRGB(0.7, 0.9, 1)
				} else {
					c.SetRGB(1, 1, 1)
				}
				drawBounds(c, b)
				c.Fill()
			}
			if len(digits) > 0 {
				digit := digits[0]
				digits = digits[1:]

				c.SetRGB(0, 0, 0)
				c.DrawString(fmt.Sprintf("%d", digit), b.Min.X, b.Min.Y+c.FontHeight())

				c.SetRGB(0, 0, 0)
				d.DrawDigit(c, b, digit)
			}
		}
	}
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
