package base3

import (
	"github.com/gitchander/godigits/utils/digits"
)

func Base3() digits.RestDigiter {
	const (
		min = 0
		max = 2
	)
	return digits.MustNewRestDigiter(min, max)
}

func Base3Bal() digits.RestDigiter {
	const (
		min = -1
		max = +1
	)
	return digits.MustNewRestDigiter(min, max)
}
