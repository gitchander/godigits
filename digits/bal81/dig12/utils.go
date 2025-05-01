package dig12

import (
	"github.com/gitchander/godigits/utils/digits"
)

func clamp(x float64, min, max float64) float64 {
	if min > max {
		return 0
	}
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func calcTrits(x int, n int) []int {
	const (
		min = -1 // digit min
		max = +1 // digit max
	)
	trits := make([]int, n)
	var trit int
	for i := range trits {
		x, trit = digits.RestDigit(x, min, max)
		trits[i] = trit
	}
	return trits
}
