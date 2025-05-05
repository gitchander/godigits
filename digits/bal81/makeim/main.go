package main

import (
	"log"
	"path/filepath"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/digits/bal81"
	"github.com/gitchander/godigits/utils"
)

func main() {
	dgdr.DrawGreedEnable = false

	checkError(testGG())
	checkError(testCairo())
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
	ds := utils.MakeInts(-40, 41, 1)
	filename := filepath.Join(se.dirName, ("bal81" + se.fileSuffix + ".png"))
	return dgdr.MakeDigitsImageMatrix(filename, se.dd, 9, 9, se.digitHeight, ds)
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
			fileSuffix:  "_d1_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit2{},
			fileSuffix:  "_d2_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit3{},
			fileSuffix:  "_d3_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit4{},
			fileSuffix:  "_d4_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit5{},
			fileSuffix:  "_d5_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit6{},
			fileSuffix:  "_d6_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit7{},
			fileSuffix:  "_d7_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit8{},
			fileSuffix:  "_d8_gg",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit9{},
			fileSuffix:  "_d9_gg",
			digitHeight: 120,
		},
		{
			dirName:     dirName,
			dd:          bal81.Digit10{},
			fileSuffix:  "_d10_gg",
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
			fileSuffix:  "_d1_cairo",
			digitHeight: 120,
		},
	}

	for _, sample := range samples {
		ds := utils.MakeInts(-40, 41, 1)
		filename := filepath.Join(sample.dirName, ("bal81" + sample.fileSuffix + ".png"))
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
