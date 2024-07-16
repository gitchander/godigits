package main

import (
	"log"
	"path/filepath"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/digits/bal81"
	"github.com/gitchander/godigits/utils"
)

func main() {
	//drawDigit11()
	//return

	dgdr.DrawGreedEnable = false

	checkError(testGG())
	checkError(testCairo())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testGG() error {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	var (
		digitHeight float64 = 90
		//digitHeight float64 = 120
	)

	samples := []sample{
		{
			dirName:     dirName,
			dd:          bal81.Digit1{},
			fileSuffix:  "_v1_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit2{},
			fileSuffix:  "_v2_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          Digit3{},
			fileSuffix:  "_v3_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          Digit4{},
			fileSuffix:  "_v4_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          Digit5{},
			fileSuffix:  "_v5_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          Digit6{},
			fileSuffix:  "_v6_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          Digit7{},
			fileSuffix:  "_v7_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          Digit8{},
			fileSuffix:  "_v8_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          Digit9{},
			fileSuffix:  "_v9_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          Digit10{},
			fileSuffix:  "_v10_gg",
			digitHeight: 120,
		},
	}

	for _, sample := range samples {
		err := makeSample(sample)
		if err != nil {
			return err
		}
	}
	return nil
}

func testCairo() error {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	samples := []sampleCairo{
		{
			dirName:     dirName,
			dd:          bal81.CairoDigit1{},
			fileSuffix:  "_v1_cairo",
			digitHeight: 120,
		},
	}

	for _, sample := range samples {
		ds := utils.MakeInts(-40, 41, 1)
		filename := filepath.Join(sample.dirName, ("digits" + sample.fileSuffix + ".png"))
		err := dgdr.CairoMakeDigitsImageMatrix(filename, sample.dd, 9, 9, sample.digitHeight, ds)
		if err != nil {
			return err
		}
	}
	return nil
}

type sampleCairo struct {
	dirName     string
	dd          dgdr.CairoDigitDrawer
	fileSuffix  string
	digitHeight float64
}
