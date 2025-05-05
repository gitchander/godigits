package digits

import (
	"fmt"
)

// check (min <= 0) and (0 <= max)
func checkMinMax(min, max int) error {
	if min > 0 {
		return fmt.Errorf("Invalid interval (%d,%d): min > 0", min, max)
	}
	if max < 0 {
		return fmt.Errorf("Invalid interval (%d,%d): max < 0", min, max)
	}
	return nil
}

//------------------------------------------------------------------------------

// number:              ... |a...b| ...
// -------|-----|-----|-----|-----|-----|-----|-----|-----
// rest:  | -3  | -2  | -1  |  0  |  1  |  2  |  3  | ...
// -------|-----|-----|-----|-----|-----|-----|-----|-----
// digit: |a...b|a...b|a...b|a...b|a...b|a...b|a...b| ...

type RestDigiter2 struct {
	min, max int
	base     int // max - min + 1
}

var _ RestDigiter = &RestDigiter2{}

func NewRestDigiter2(min, max int) (*RestDigiter2, error) {

	err := checkMinMax(min, max)
	if err != nil {
		return nil, err
	}

	d := &RestDigiter2{
		min: min,
		max: max,

		base: max - min + 1,
	}
	return d, nil
}

func MustNewRestDigiter2(min, max int) *RestDigiter2 {
	d, err := NewRestDigiter2(min, max)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *RestDigiter2) RestDigit(x int) (rest, digit int) {

	var (
		min = d.min
		max = d.max

		base = d.base
	)

	switch {
	case x < min:
		rest, digit = quoRem((x - (min - 1)), base)
		rest--
		digit += max
	case x > max:
		rest, digit = quoRem((x - (max + 1)), base)
		rest++
		digit += min
	default:
		rest = 0
		digit = x
	}

	return
}
