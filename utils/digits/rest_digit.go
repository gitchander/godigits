package digits

import (
	"fmt"
	"math"
)

type RestDigiter interface {
	RestDigit(x int) (rest, digit int)
}

func NewRestDigiter(min, max int) (RestDigiter, error) {
	rd, err := NewDigiter2(min, max)
	return rd, err
}

func MustNewRestDigiter(min, max int) RestDigiter {
	rd, err := NewRestDigiter(min, max)
	if err != nil {
		panic(err)
	}
	return rd
}

//------------------------------------------------------------------------------

func checkMinMax(min, max int) error {
	if min > 0 {
		return fmt.Errorf("Invalid interval (%d,%d): min > 0", min, max)
	}
	if max < 0 {
		return fmt.Errorf("Invalid interval (%d,%d): max < 0", min, max)
	}
	return nil
}

// RestDigit
// a <= b

// value: ....................... | a ... b | .......................
// rest:  ... |   -2    |   -1    |    0    |    1    |    2    | ...
// digit: ... | a ... b | a ... b | a ... b | a ... b | a ... b | ...

// min <= max
// base = max - min + 1
// x = rest * base + digit

//------------------------------------------------------------------------------

type restDigiter1 struct {
	min, max int
	base     int
}

var _ RestDigiter = &restDigiter1{}

func newRestDigiter1(min, max int) (*restDigiter1, error) {

	err := checkMinMax(min, max)
	if err != nil {
		return nil, err
	}

	rd := &restDigiter1{
		min: min,
		max: max,

		base: max - min + 1,
	}
	return rd, nil
}

func (p *restDigiter1) RestDigit(x int) (rest, digit int) {

	var (
		min = p.min
		max = p.max

		base = p.base
	)

	var q, r int

	if x < min {
		q, r = quoRem(x-max, base)
		r += max
	} else {
		q, r = quoRem(x-min, base)
		r += min
	}

	rest = q
	digit = r

	return
}

func calcRestDigit1_mod(x int, min, max int) (rest, digit int) {

	base := max - min + 1

	var q, r int

	switch {
	case x < min:
		dx := min - 1
		q, r = quoRem(x-dx, base)
		q--
		r += dx + base
	case x > max:
		dx := max + 1
		q, r = quoRem(x-dx, base)
		q++
		r += dx - base
	default:
		q = 0
		r = x
	}

	rest = q
	digit = r

	return
}

func calcRestDigit2(x int, min, max int) (rest, digit int) {

	base := max - min + 1

	rest, digit = quoRem(x, base)

	for digit < min {
		if rest == math.MinInt {
			panic("overflow min")
		}
		rest--
		digit += base
	}
	for digit > max {
		if rest == math.MaxInt {
			panic("overflow max")
		}
		rest++
		digit -= base
	}

	return
}

//------------------------------------------------------------------------------

type restDigiter3 struct {
	min, max int
	base     int
}

var _ RestDigiter = &restDigiter3{}

func newRestDigiter3(min, max int) (*restDigiter3, error) {

	err := checkMinMax(min, max)
	if err != nil {
		return nil, err
	}

	rd := &restDigiter3{
		min: min,
		max: max,

		base: max - min + 1,
	}
	return rd, nil
}

func (p *restDigiter3) RestDigit(x int) (rest, digit int) {

	var (
		min = p.min
		max = p.max

		base = p.base
	)

	rest, digit = quoRem(x, base)

	if digit < min {
		k := ceilDiv(min-digit, base)
		rest -= k
		digit += base * k
	}
	if digit > max {
		k := ceilDiv(digit-max, base)
		rest += k
		digit -= base * k
	}

	return
}
