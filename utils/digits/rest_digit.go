package digits

import (
	"math"
)

type RestDigiter interface {
	RestDigit(x int) (rest, digit int)
}

func NewRestDigiter(min, max int) (RestDigiter, error) {
	rd, err := NewRestDigiter2(min, max)
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

	if x < min {
		rest, digit = quoRem(x-max, base)
		digit += max
	} else {
		rest, digit = quoRem(x-min, base)
		digit += min
	}

	return
}

func calcRestDigit1_mod(x int, min, max int) (rest, digit int) {

	base := max - min + 1

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
