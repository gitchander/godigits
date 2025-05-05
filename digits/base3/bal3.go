package base3

import (
	"fmt"
	"strings"

	"github.com/gitchander/godigits/utils/digits"
)

const (
	// charNegative = 'T'
	// charZero     = '0'
	// charPositive = '1'

	// charNegative = 'N'
	// charZero     = '0'
	// charPositive = '1'

	// charNegative = 'N'
	// charZero     = '0'
	// charPositive = 'P'

	charNegative = 'N'
	charZero     = 'Z'
	charPositive = 'P'
)

func FormatDigits(ds []int) (string, error) {
	var b strings.Builder
	for i, d := range ds {
		var char byte
		switch d {
		case -1:
			char = charNegative
		case 0:
			char = charZero
		case +1:
			char = charPositive
		default:
			return "", fmt.Errorf("Invalid digit [index:%d] %d", i, d)
		}
		b.WriteByte(char)
	}
	return b.String(), nil
}

func DigitsToInt(ds []int, rest int) int {
	const (
		min = -1
		max = +1
	)
	return digits.DigitsToInt(min, max, ds, rest)
}
