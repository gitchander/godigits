package base13

import (
	"github.com/gitchander/godigits/utils/digits"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcDigits(v int, ds []int) []int {
	const (
		min = -1
		max = +1
	)
	rd := digits.MustNewRestDigiter(min, max)
	digits.CalcDigits(rd, v, ds)
	return ds
}
