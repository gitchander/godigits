package bal81

import (
	"github.com/gitchander/godigits/numbers"
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
	b := numbers.Bal3
	trits := make([]int, n)
	for i := range trits {
		x, trits[i] = numbers.RestDigit(b, x)
	}
	return trits
}
