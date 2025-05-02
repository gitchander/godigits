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
	rd := digits.MustNewRestDigiter(min, max)
	trits := make([]int, n)
	var trit int
	for i := range trits {
		x, trit = rd.RestDigit(x)
		trits[i] = trit
	}
	return trits
}
