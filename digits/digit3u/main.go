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

	samples := []Sample{
		{
			filename:  "digits_v3_4x4_0p.png",
			dd:        Digit1{},
			valueMin:  0,
			valueMax:  16,
			valueStep: 1,

			matrixXn:    4,
			matrixYn:    4,
			digitHeight: 256,
		},
	}

	for _, sample := range samples {

		ds := utils.MakeInts(sample.valueMin, sample.valueMax, sample.valueStep)
		filename := filepath.Join(dirName, sample.filename)

		err := dgdr.MakeDigitsImageMatrix(filename, sample.dd,
			sample.matrixXn, sample.matrixYn, sample.digitHeight, ds)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Sample struct {
	filename string

	dd dgdr.DigitDrawer

	valueMin  int
	valueMax  int
	valueStep int

	matrixXn int
	matrixYn int

	digitHeight float64
}

func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

type Digit1 struct{}

func (Digit1) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit1) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 8
		ny = 16

		dx = digitHeight / float64(ny)
		dy = dx

		relativeSize = 0.5

		greedWidth = 0.025 * dx
		lineWidth  = relativeSize * dx

		radius = relativeSize
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	//----------------------------------------------------------------------

	var (
		// x1, x2 float64 = 2, 6
		x1, x2 float64 = 3, 5

		y1, y2 float64 = 3, 13
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

	vert := 0
	if (val % 2) != 0 {
		vert = 3
	}

	switch vert {

	case 0:
		{
			c.MoveTo(x1, y1)
			c.LineTo(x1, y2)

			c.MoveTo(x2, y1)
			c.LineTo(x2, y2)
		}

	case 1:
		{
			c.MoveTo(x1, y1)
			c.LineTo(x2, y2)

			c.MoveTo(x2, y1)
			c.LineTo(x1, y2)
		}

	case 2:
		{
			c.MoveTo(x1, y1)
			c.CubicTo(x1, 8, x2, 8, x2, y2)

			c.MoveTo(x2, y1)
			c.CubicTo(x2, 8, x1, 8, x1, y2)
		}

	case 3:
		{
			{
				c.MoveTo(x1, y1)
				c.LineTo(x1, 6)
				c.CubicTo(x1, 8, x2, 8, x2, 10)

				c.MoveTo(x2, 10)
				c.LineTo(x2, y2)
			}
			//----------------------------------------------
			{
				c.MoveTo(x2, y1)
				c.LineTo(x2, 6)
				c.CubicTo(x2, 8, x1, 8, x1, 10)

				c.MoveTo(x1, 10)
				c.LineTo(x1, y2)
			}
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
		case 2, 3:
			bs = parseBools("00000-1-00000")

		//---------------------------------------------
		case 4, 5:
			bs = parseBools("01000-0-00010")
		case 6, 7:
			bs = parseBools("01000-1-00010")

		// case 4, 5:
		// 	bs = parseBools("00100-0-00100")
		// case 6, 7:
		// 	bs = parseBools("00100-1-00100")

		//---------------------------------------------
		case 8, 9:
			bs = parseBools("01100-0-00110")
		case 10, 11:
			bs = parseBools("01100-1-00110")

		// case 8, 9:
		// 	bs = parseBools("01010-0-01010")
		// case 10, 11:
		// 	bs = parseBools("01010-1-01010")

		//---------------------------------------------

		case 12, 13:
			bs = parseBools("01110-0-01110")
		case 14, 15:
			bs = parseBools("01110-1-01110")
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

	if negative {
		if false {
			c.DrawCircle(x1, 8, radius)
			c.DrawCircle(x2, 8, radius)
			c.Fill()
		} else {
			// c.DrawCircle(4, 6, radius)
			// c.DrawCircle(4, 10, radius)

			c.DrawCircle(4, 6.5, radius)
			c.DrawCircle(4, 9.5, radius)
		}
		c.Fill()
	}
}

func parseBools(s string) []bool {
	var (
		rs = []rune(s)
		bs = make([]bool, 0, len(rs))
	)
	for _, r := range rs {
		switch r {
		case '0':
			bs = append(bs, false)
		case '1':
			bs = append(bs, true)
		case '-', '_': // skip
		default:
			panic("invalid bits")
		}
	}
	return bs
}
