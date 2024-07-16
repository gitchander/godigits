package main

import (
	"log"
	"path/filepath"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/digits/bal81/dig11"
	"github.com/gitchander/godigits/utils"
)

func main() {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	checkError(drawTestDigit(dirName))
	checkError(makeImages(dirName))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func drawTestDigit(dirName string) error {

	dgdr.DrawGreedEnable = true
	var dd dgdr.DigitDrawer

	dd = dig11.Digit11_v1{}
	//dd = Digit11_v2{}
	//dd = Digit11_v3{}
	//dd = Digit11_v4{}
	//dd = Digit11_v5{}
	//dd = Digit11_v6{}

	filename := filepath.Join(dirName, "digit11.png")
	return dgdr.MakeDigitImage(filename, dd, 512, -31)
}

func makeImages(dirName string) error {

	dgdr.DrawGreedEnable = false

	var (
		//digitHeight float64 = 90
		digitHeight float64 = 120
	)

	samples := []sample{
		{
			dirName:     dirName,
			dd:          dig11.Digit11_v1{},
			fileSuffix:  "_v1",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          dig11.Digit11_v2{},
			fileSuffix:  "_v2",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          dig11.Digit11_v3{},
			fileSuffix:  "_v3",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          dig11.Digit11_v4{},
			fileSuffix:  "_v4",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          dig11.Digit11_v5{},
			fileSuffix:  "_v5",
			digitHeight: digitHeight,
		},
		{
			dirName:     dirName,
			dd:          dig11.Digit11_v6{},
			fileSuffix:  "_v6",
			digitHeight: digitHeight,
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
