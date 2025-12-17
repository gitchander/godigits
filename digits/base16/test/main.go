package main

import (
	"log"
	"path/filepath"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/digits/base16"
	"github.com/gitchander/godigits/utils"
)

func main() {
	checkError(run())
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

func run() error {
	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	samples := []Sample{
		{
			filename:  "base16_d1.png",
			dd:        base16.Digit1{},
			valueMin:  0,
			valueMax:  16,
			valueStep: 1,

			matrixXn:    4,
			matrixYn:    4,
			digitHeight: 160,
		},
		{
			filename:  "base16_d1_02.png",
			dd:        base16.Digit1{},
			valueMin:  -17,
			valueMax:  18,
			valueStep: 1,

			matrixXn:    7,
			matrixYn:    5,
			digitHeight: 160,
		},
		{
			filename:  "base16_d1_03.png",
			dd:        base16.Digit1{},
			valueMin:  -22,
			valueMax:  23,
			valueStep: 1,

			matrixXn:    9,
			matrixYn:    5,
			digitHeight: 160,
		},
		{
			filename:  "base16_d1_04.png",
			dd:        base16.Digit1{},
			valueMin:  -16,
			valueMax:  17,
			valueStep: 1,

			matrixXn:    11,
			matrixYn:    3,
			digitHeight: 160,
		},
		{
			filename:  "base16_d2.png",
			dd:        base16.Digit2{},
			valueMin:  0,
			valueMax:  16,
			valueStep: 1,

			matrixXn:    4,
			matrixYn:    4,
			digitHeight: 160,
		},
		{
			filename:  "base16_d3.png",
			dd:        base16.Digit3{},
			valueMin:  0,
			valueMax:  16,
			valueStep: 1,

			matrixXn:    4,
			matrixYn:    4,
			digitHeight: 160,
		},
		{
			filename:  "base16_d4.png",
			dd:        base16.Digit4{},
			valueMin:  0,
			valueMax:  16,
			valueStep: 1,

			matrixXn:    4,
			matrixYn:    4,
			digitHeight: 160,
		},
		{
			filename:  "base16_d5.png",
			dd:        base16.Digit5{},
			valueMin:  0,
			valueMax:  16,
			valueStep: 1,

			matrixXn:    4,
			matrixYn:    4,
			digitHeight: 160,
		},
	}

	for _, sample := range samples {

		ds := utils.MakeInts(sample.valueMin, sample.valueMax, sample.valueStep)
		filename := filepath.Join(dirName, sample.filename)

		err := dgdr.MakeDigitsImageMatrix(filename, sample.dd,
			sample.matrixXn, sample.matrixYn, sample.digitHeight, ds)
		checkError(err)
	}
	return nil
}
