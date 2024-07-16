package main

import (
	"log"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils"
)

func main() {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	var dd dgdr.DigitDrawer

	dd = Digit1{}
	//dd = DigitSasha{}

	ds := serialInts(16)
	filename := filepath.Join(dirName, "digits.png")
	err := dgdr.MakeDigitsImageMatrix(filename, dd, 4, 4, 200, ds)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

func randInts(n int, min, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := make([]int, n)
	for i := range a {
		a[i] = min + r.Intn(max-min)
	}
	return a
}
