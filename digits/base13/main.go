package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
	"github.com/gitchander/godigits/utils/digits"
	"github.com/gitchander/godigits/utils/random"
)

func main() {
	makeDigits()
	makeNumbers()
	makeDigitsRandom()
}

func makeDigits() {
	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	var dd dgdr.DigitDrawer

	dd = Digit1{}

	se := sample{
		dirName:     dirName,
		dd:          dd,
		fileSuffix:  "_d4",
		digitHeight: 160,
	}

	err := makeSample(se)
	checkError(err)
}

func makeNumbers() {
	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)
	dd := Digit2{}
	var ds []int
	var (
		trits = 4
		w     = pow(3, trits)

		min = -(w - 1) / 2
		max = +(w - 1) / 2
	)
	ds = appendInts(ds, min, max+1)
	for i, d := range ds {
		filename := filepath.Join(dirName, fmt.Sprintf("digits_%04d.png", i))
		ds := numberToDigits(d)
		err := dgdr.MakeDigitsImage(filename, dd, 256, ds)
		checkError(err)
	}
}

func pow(x, n int) int {
	y := 1
	for i := 0; i < n; i++ {
		y *= x
	}
	return y
}

func makeDigitsRandom() {
	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	dd := Digit2{}
	//dd := Digit6{}

	r := random.NewRandNow()
	ds := make([]int, 12)
	for i := range ds {
		ds[i] = random.RandIntMinMax(r, -40, 41)
	}
	//intsSerial(ds, -40, 1)

	var (
		//digitHeight float64 = 22
		//digitHeight float64 = 64
		digitHeight float64 = 128
	)

	filename := filepath.Join(dirName, "digits_n_random.png")
	err := dgdr.MakeDigitsImage(filename, dd, digitHeight, ds)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type sample struct {
	dirName     string
	dd          dgdr.DigitDrawer
	fileSuffix  string
	digitHeight float64
}

func makeSample(se sample) error {

	ds := utils.MakeInts(-13, 14, 1)
	filename := filepath.Join(se.dirName, ("digits" + se.fileSuffix + ".png"))
	err := dgdr.MakeDigitsImageMatrix(filename, se.dd, 9, 3, se.digitHeight, ds)
	if err != nil {
		return err
	}

	return nil
}

func calcDigits(v int, ds []int) []int {
	const (
		min = -1
		max = +1
	)
	digits.CalcDigits(v, min, max, ds)
	return ds
}

func numberToDigits(x int) []int {
	const (
		min = -40
		max = 40
	)
	ds, _ := digits.CalcDigitsN(x, min, max, 10)
	return ds
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func appendInts(p []int, min, max int) []int {
	for x := min; x < max; x++ {
		p = append(p, x)
	}
	return p
}

func intsSerial(as []int, min int, step int) {
	for i := range as {
		as[i] = min + i*step
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
