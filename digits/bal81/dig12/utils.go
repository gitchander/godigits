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

func calcTrits(v int, trits []int) {
	const (
		min = -1 // digit min
		max = +1 // digit max
	)
	rd := digits.MustNewRestDigiter(min, max)
	for i := range trits {
		v, trits[i] = rd.RestDigit(v)
	}
}

func calcTritsBal81(v int) []int {
	trits := make([]int, 4)
	calcTrits(v, trits)
	return trits
}
