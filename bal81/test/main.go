package main

import (
	"image"
	"log"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/bal81"
	"github.com/gitchander/godigits/geom"
)

func main() {

	nA := 20
	dA := float64(nA)

	digitSize := image.Pt(nA*2, nA*4)

	d := bal81.Digit1{
		A: dA,
		B: dA * 0.2,
		C: dA * 0.2,
	}.DigitDrawer()

	//digits := []int{-4, -3, -2, 2, 3, 4, 40}
	//digits := serialInts(7)
	digits := intervalInts(-6, 6)

	c := gg.NewContext(digitSize.X*len(digits), digitSize.Y)

	if true {
		c.SetRGB(1, 1, 1)
		c.Clear()
	}

	c.SetRGB(0, 0, 0)

	b := geom.MakeBounds(0, 0, float64(digitSize.X), float64(digitSize.Y))
	bsh := geom.MakePoint(float64(digitSize.X), 0)

	for _, digit := range digits {
		d.DrawDigit(c, b, digit)
		b = b.Add(bsh)
	}

	err := c.SavePNG("result.png")
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
