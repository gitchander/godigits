package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gitchander/godigits/utils"
)

func main() {
	var (
		//size float64 = 32
		//size float64 = 64
		//size float64 = 128
		size float64 = 256
	)

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	var (
		trits = 3

		w   = pow(3, trits)
		min = -(w - 1) / 2
		max = +(w - 1) / 2
	)
	fmt.Println(w, min, max)

	var ds []int
	ds = appendInts(ds, min, max+1)

	for i, d := range ds {
		filename := filepath.Join(dirName, fmt.Sprintf("dig_%04d.png", i))
		err := MakeNumberImage(filename, d, size)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func appendInts(p []int, min, max int) []int {
	for x := min; x < max; x++ {
		p = append(p, x)
	}
	return p
}

func pow(x, n int) int {
	y := 1
	for i := 0; i < n; i++ {
		y *= x
	}
	return y
}
