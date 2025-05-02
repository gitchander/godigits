package dig11

import (
	"math"

	"github.com/gitchander/godigits/utils/digits"
)

func calcDigits(v int, ds []int) []int {
	const (
		min = -1
		max = +1
	)
	rd := digits.MustNewRestDigiter(min, max)
	digits.CalcDigits(rd, v, ds)
	return ds
}

func trimLast0(digits []int) []int {
	k := 0
	for i, digit := range digits {
		if digit != 0 {
			k = i
		}
	}
	return digits[:k+1]
}

const (
	tau = 2 * math.Pi
)

func DegToRad(deg float64) float64 {
	return deg * tau / 360
}

func RadToDeg(rad float64) float64 {
	return rad * 360 / tau
}
