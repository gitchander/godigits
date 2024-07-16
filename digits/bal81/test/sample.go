package main

import (
	"path/filepath"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
)

type sample struct {
	dirName     string
	dd          dgdr.DigitDrawer
	fileSuffix  string
	digitHeight float64
}

func makeSample(se sample) error {
	ds := utils.MakeInts(-40, 41, 1)
	filename := filepath.Join(se.dirName, ("digits" + se.fileSuffix + ".png"))
	return dgdr.MakeDigitsImageMatrix(filename, se.dd, 9, 9, se.digitHeight, ds)
}
