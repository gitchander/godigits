package main

import (
	"fmt"
	"math"

	"github.com/gitchander/godigits/digits/base3"
)

func main() {

	var xs []int

	k := 15

	for i := 0; i < k; i++ {
		xs = append(xs, math.MinInt+i)
	}
	xs = append(xs, serialInts(-11, 12)...)
	for i := 0; i < k; i++ {
		xs = append(xs, math.MaxInt-k+1+i)
	}

	d := base3.Base3Bal()

	for _, x := range xs {
		rest, digit := d.RestDigit(x)
		fmt.Printf("%4d %4d %4d\n", x, rest, digit)
	}
}

func serialInts(a, b int) []int {
	n := b - a
	if n < 0 {
		n = 0
	}
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = a + i
	}
	return xs
}
