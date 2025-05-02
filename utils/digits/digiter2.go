package digits

import (
	"fmt"
)

// number:              ... |a...b| ...
// -------|-----|-----|-----|-----|-----|-----|-----|-----
// rest:  | -3  | -2  | -1  |  0  |  1  |  2  |  3  | ...
// -------|-----|-----|-----|-----|-----|-----|-----|-----
// digit: |a...b|a...b|a...b|a...b|a...b|a...b|a...b| ...

type Digiter2 struct {
	min, max int
	base     int // max - min + 1
}

var _ RestDigiter = &Digiter2{}

func NewDigiter2(min, max int) (*Digiter2, error) {

	// err:= checkMinMax(min, max)
	// if err != nil {
	// 	return nil, err
	// }

	if min > 0 {
		return nil, fmt.Errorf("Invalid interval (%d,%d): min > 0", min, max)
	}
	if max < 0 {
		return nil, fmt.Errorf("Invalid interval (%d,%d): max < 0", min, max)
	}

	d := &Digiter2{
		min: min,
		max: max,

		base: max - min + 1,
	}
	return d, nil
}

func MustNewDigiter2(min, max int) *Digiter2 {
	d, err := NewDigiter2(min, max)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *Digiter2) RestDigit(x int) (rest, digit int) {

	var (
		min = d.min
		max = d.max

		base = d.base // max - min + 1
	)

	if x < min {
		rest, digit := quoRem((x - (min - 1)), base)
		rest--
		digit += max
		return rest, digit
	}

	if x > max {
		rest, digit := quoRem((x - (max + 1)), base)
		rest++
		digit += min
		return rest, digit
	}

	return 0, x
}
