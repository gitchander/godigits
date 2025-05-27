package main

import (
	"fmt"
	"image"
	"log"
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/digits/base3"
	"github.com/gitchander/godigits/utils/digits"
	"github.com/gitchander/godigits/utils/random"
)

func main() {
	// testRestDigit()
	// testDigits()

	checkError(makeDigitMatrix("base3bal_d1.png", base3.DigitDrawer1{}))
	checkError(makeDigitMatrix("base3bal_d2.png", base3.DigitDrawer2{}))
	checkError(makeDigitMatrix("base3bal_d3.png", base3.DigitDrawer3{}))

	checkError(makeRandDigits("base3bal_d1_rand.png", base3.DigitDrawer1{}))
	checkError(makeRandDigits("base3bal_d2_rand.png", base3.DigitDrawer2{}))
	checkError(makeRandDigits("base3bal_d3_rand.png", base3.DigitDrawer3{}))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testRestDigit() {
	var (
		xs = makeSomeInts()
		rd = base3.Base3Bal()
	)
	for _, x := range xs {
		rest, digit := rd.RestDigit(x)
		fmt.Printf("%4d %4d %4d\n", x, rest, digit)
	}
}

func makeSomeInts() []int {
	var xs []int
	k := 15
	for i := 0; i < k; i++ {
		xs = append(xs, math.MinInt+i)
	}
	p := 41
	xs = append(xs, serialInts((-p+0), (+p+1))...)
	for i := 0; i < k; i++ {
		xs = append(xs, math.MaxInt-k+1+i)
	}
	return xs
}

func serialInts(a, b int) []int {
	n := b - a
	if n < 0 {
		n = 0
	}
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = a + i
	}
	return xs
}

func intervalInts(min, max int) []int {
	n := max - min
	if n < 0 {
		n = 0
	}
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = min + i
	}
	return a
}

func testDigits() {
	rd := base3.Base3Bal()

	//xs := serialInts(-1000, +1000)
	xs := makeSomeInts()

	for _, x := range xs {
		ds, rest := digits.CalcDigitsN(rd, x, 100)

		fds, err := base3.FormatDigits(ds)
		if err != nil {
			panic(err)
		}

		y := base3.DigitsToInt(ds, rest)
		if x != y {
			err := fmt.Errorf("%d != %d", x, y)
			panic(err)
		}

		fmt.Printf("%d: {%s %d}: %d\n", x, fds, rest, y)
	}
}

func makeRandDigits(filename string, d dgdr.DigitDrawerB) error {

	var (
		sizeX = 32
		sizeY = base3.CalcSizeY(sizeX)
	)

	digitSize := image.Pt(sizeX, sizeY)

	//digits := intervalInts((-1 + 0), (+1 + 1))
	r := random.NextRand()
	digits := make([]int, 10)
	for i := range digits {
		digits[i] = random.RandIntIn(r, -1, +1+1)
	}

	var (
		nX = len(digits)
		nY = 1

		cX = digitSize.X * nX
		cY = digitSize.Y * nY
	)
	c := gg.NewContext(cX, cY)

	if true {
		c.SetRGB(1, 1, 1)
		c.Clear()
	}

	fontSize := float64(digitSize.Y) * 0.1
	err := dgdr.SetFontSizeGG(c, fontSize)
	if err != nil {
		return err
	}

	var (
		dXf = float64(digitSize.X)
		dYf = float64(digitSize.Y)
	)
	dgdr.DrawDigitsDDB(c, d, dXf, dYf, digits)

	return c.SavePNG(filename)
}

func makeDigitMatrix(filename string, d dgdr.DigitDrawerB) error {

	var (
		sizeX = 50
		sizeY = base3.CalcSizeY(sizeX)
	)

	digitSize := image.Pt(sizeX, sizeY)

	digits := intervalInts((-1 + 0), (+1 + 1))

	var (
		nX = 3
		nY = 1

		cX = digitSize.X * nX
		cY = digitSize.Y * nY
	)
	c := gg.NewContext(cX, cY)

	if true {
		c.SetRGB(1, 1, 1)
		c.Clear()
	}

	fontSize := float64(digitSize.Y) * 0.1
	err := dgdr.SetFontSizeGG(c, fontSize)
	if err != nil {
		return err
	}

	dgdr.DrawMatrixDDB(c, d, nX, nY, digitSize, digits)

	return c.SavePNG(filename)
}
