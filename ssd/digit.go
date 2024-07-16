package main

import (
	"image/color"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
)

type DigitSSD struct{}

var _ dgdr.DigitDrawer = DigitSSD{}

func (DigitSSD) Width(height float64) (width float64) {
	width = height / 2
	return width
}

func (d DigitSSD) DrawDigit(c *gg.Context, x0, y0 float64, digitHeight float64, digit int) {
	// todo
}

func digitBounds(size float64) gg.Point {
	return gg.Point{
		X: size / 2,
		Y: size,
	}
}

func makeImage() {
	var (
		//size float64 = 16
		//size float64 = 32
		//size float64 = 64
		//size float64 = 128
		//size float64 = 256

		size float64 = 84
	)

	dirName := "images"
	err := os.Mkdir(dirName, os.ModePerm)
	if !(os.IsExist(err)) {
		checkError(err)
	}

	if false {
		var (
			//ds = serialInts(25)
			ds = generateDigits()
		)

		filename := filepath.Join(dirName, "digits.png")
		err = makeDigitsImage(filename, ds, size)
		checkError(err)
	}

	{
		const (
			min = -19
			max = +19
		)
		ds := utils.MakeInts(min, (max + 1), 1)

		filename := filepath.Join(dirName, "digits_matrix.png")
		err = makeDigitsImageMatrix(filename, ds, 13, 3, 128)
		checkError(err)
	}
}

func makeDigitsImage(filename string, digits []uint8, size float64) error {

	var (
		digitSize     = digitBounds(size)
		width, height = ceilInt(digitSize.X * float64(len(digits))), ceilInt(digitSize.Y)
	)

	c := gg.NewContext(width, height)

	// fill background
	c.SetColor(color.White)
	c.Clear()

	c.SetColor(color.Black)

	var x0, y0 float64 = 0, 0
	for _, digit := range digits {
		err := drawSSD(c, x0, y0, size, digit)
		if err != nil {
			return err
		}
		x0 += digitSize.X
	}

	return c.SavePNG(filename)
}

func makeDigitsImageMatrix(filename string, digits []int, xn, yn int, size float64) error {

	var (
		digitSize = digitBounds(size)

		width  = ceilInt(digitSize.X * float64(xn))
		height = ceilInt(digitSize.Y * float64(yn))
	)

	c := gg.NewContext(width, height)

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 16})
	c.SetFontFace(face)

	// fill background
	// c.SetColor(color.White)
	// c.Clear()

	c.SetColor(color.Black)

start:
	for y := 0; y < yn; y++ {
		for x := 0; x < xn; x++ {
			if len(digits) == 0 {
				break start
			}
			digit := digits[0]
			digits = digits[1:]

			var (
				x0 = float64(x) * digitSize.X
				y0 = float64(y) * digitSize.Y
			)

			if true {
				if ((x + y) % 2) == 0 {
					c.SetRGB(1, 1, 1)
				} else {
					c.SetRGB(0.7, 0.9, 1)
				}
				c.DrawRectangle(x0, y0, digitSize.X, digitSize.Y)
				c.Fill()
				c.SetColor(color.Black)
			}

			c.DrawString(formatInt(digit), x0+8, y0+20)

			digitBits := digitsMap[digit]

			err := drawSSD(c, x0, y0, size, digitBits)
			if err != nil {
				return err
			}
		}
	}

	return c.SavePNG(filename)
}
