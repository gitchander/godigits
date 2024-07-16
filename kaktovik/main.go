package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
)

func main() {
	printKaktovikNumerals()
	drawDigits()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Kaktovik numerals
func printKaktovikNumerals() {

	// '\U0001D2C0' - 0
	// '\U0001D2CF' - 15
	// '\U0001D2D0' - 16
	// '\U0001D2D3' - 19

	r := '\U0001D2C1'
	fmt.Printf("(%#U)\n", r)
	fmt.Printf("(%#U)\n", 'Ъ')

	var min uint32 = 0x0001D2C0

	for i := 0; i < 20; i++ {
		r := rune(min + uint32(i))
		fmt.Printf("digit %2d: (%#U)\n", i, r)
	}
}

type sample struct {
	dirName     string
	dd          dgdr.DigitDrawer
	fileSuffix  string
	digitHeight float64
}

func drawDigits() {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	samples := []sample{
		{
			dirName:     dirName,
			dd:          Digit1{},
			fileSuffix:  "_d1",
			digitHeight: 128,
		},
	}

	for _, sample := range samples {
		err := makeSample(sample)
		checkError(err)
	}
}

func makeSample(se sample) error {

	ds := utils.MakeInts(0, 20, 1)
	filename := filepath.Join(se.dirName, ("digits" + se.fileSuffix + ".png"))
	err := dgdr.MakeDigitsImageMatrix(filename, se.dd, 5, 4, se.digitHeight, ds)
	if err != nil {
		return err
	}

	return nil
}
