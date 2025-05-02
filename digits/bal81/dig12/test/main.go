package main

import (
	"image"
	"log"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/digits/bal81/dig12"
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

	var (
		//d = dig12.MakeDigit1_p1()
		d = dig12.Digit2{}
	)

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

	dgdr.DrawMatrixDDB(c, d, nX, nY, digitSize, digits)

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

func setFont(c *gg.Context, fontSize float64) error {

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	c.SetFontFace(face)

	return nil
}
