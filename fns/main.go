package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
	"github.com/gitchander/godigits/utils/random"
)

// The Forgotten Number System - Numberphile
// https://www.youtube.com/watch?v=9p55Qgt7Ciw

func main() {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	var dd dgdr.DigitDrawer

	dd = Digit1{}

	{
		ds := []int{1410, 4173, 5750, 1368, 6666}
		//ds = randInts(100, 0, 10000)
		for _, d := range ds {
			filename := filepath.Join(dirName, fmt.Sprintf("digits_%04d.png", d))
			err := dgdr.MakeDigitImage(filename, dd, 128, d)
			checkError(err)
		}
	}

	{
		ds := serialInts(10)
		//ds := randInts(10, 0, 10000)
		filename := filepath.Join(dirName, "digits.png")
		err := dgdr.MakeDigitsImage(filename, dd, 128, ds)
		checkError(err)
		fmt.Println(ds)
	}

	{
		ds := randInts(100, 0, 10000)
		filename := filepath.Join(dirName, "digits_matrix.png")
		err := dgdr.MakeDigitsImageMatrix(filename, dd, 10, 6, 128, ds)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Digit1 struct{}

func (Digit1) Width(height float64) (width float64) {
	width = height / 2
	return width
}

func (d Digit1) DrawDigit(c *gg.Context, x0, y0 float64, height float64, number int) {

	var (
		nx = 4
		ny = 8

		dx = height / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.30 * dx
	)

	c.Push()
	defer c.Pop()

	c.Translate(x0, y0)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var (
		x1 = 2.0
		y1 = 2.0

		//y2 = 5.0
		y2 = 6.0

		xxd = 1.0
		yyd = 1.0
	)

	c.MoveTo(x1, y1)
	c.LineTo(x1, y2)
	c.Stroke()

	type coordDelta struct {
		coord gg.Point
		delta gg.Point
	}

	cds := []coordDelta{
		// Units
		{
			coord: gg.Point{X: x1, Y: y1},
			delta: gg.Point{X: xxd, Y: yyd},
		},
		// Tens
		{
			coord: gg.Point{X: x1, Y: y1},
			delta: gg.Point{X: -xxd, Y: yyd},
		},
		// Hundreds
		{
			coord: gg.Point{X: x1, Y: y2},
			delta: gg.Point{X: xxd, Y: -yyd},
		},
		// Thousands
		{
			coord: gg.Point{X: x1, Y: y2},
			delta: gg.Point{X: -xxd, Y: -yyd},
		},
	}

	digits := getDigits(number, len(cds))

	for i, digit := range digits {

		var (
			cd    = cds[i]
			coord = cd.coord
			delta = cd.delta
		)

		switch digit {
		case 1:
			c.MoveTo(coord.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y)
		case 2:
			c.MoveTo(coord.X, coord.Y+delta.Y)
			c.LineTo(coord.X+delta.X, coord.Y+delta.Y)
		case 3:
			c.MoveTo(coord.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y+delta.Y)
		case 4:
			c.MoveTo(coord.X, coord.Y+delta.Y)
			c.LineTo(coord.X+delta.X, coord.Y)
		case 5:
			c.MoveTo(coord.X, coord.Y+delta.Y)
			c.LineTo(coord.X+delta.X, coord.Y)
			c.LineTo(coord.X, coord.Y)
		case 6:
			c.MoveTo(coord.X+delta.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y+delta.Y)
		case 7:
			c.MoveTo(coord.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y+delta.Y)
		case 8:
			c.MoveTo(coord.X+delta.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y+delta.Y)
			c.LineTo(coord.X, coord.Y+delta.Y)
		case 9:
			c.MoveTo(coord.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y)
			c.LineTo(coord.X+delta.X, coord.Y+delta.Y)
			c.LineTo(coord.X, coord.Y+delta.Y)
		}
	}
	c.Stroke()
}

func getDigits(v int, n int) []int {
	ds := make([]int, n)
	for i := range ds {
		v, ds[i] = quoRem(v, 10)
	}
	return ds
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

func randInts(n int, min, max int) []int {
	r := random.NewRandNow()
	return random.RandInts(r, n, min, max)
}
