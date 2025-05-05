package digits

import (
	"fmt"
)

type Digiter1 struct {
	min, max int
	base     int

	rd RestDigiter
}

var _ RestDigiter = &Digiter1{}

func NewDigiter1(min, max int) (*Digiter1, error) {

	if min > max {
		return nil, fmt.Errorf("Invalid interval (%d,%d)", min, max)
	}
	rd, err := NewRestDigiter(min, max)
	if err != nil {
		return nil, err
	}

	d := &Digiter1{
		min: min,
		max: max,

		base: (max - min + 1),

		rd: rd,
	}
	return d, nil
}

func MustNewDigiter1(min, max int) *Digiter1 {
	d, err := NewDigiter1(min, max)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *Digiter1) checkDigit(digit int) error {
	if (d.min <= digit) && (digit <= d.max) {
		return nil
	}
	return fmt.Errorf("invalid digit %d, want interval [%d .. %d]", digit, d.min, d.max)
}

func (d *Digiter1) Base() int {
	return d.base
}

func (d *Digiter1) RestDigit(x int) (rest, digit int) {
	return d.rd.RestDigit(x)
}

// dl - digit interval
func (d *Digiter1) IntToDigits(v int, ds []int) (rest int) {
	var digit int
	for i := range ds {
		v, digit = d.RestDigit(v)
		ds[i] = digit
	}
	rest = v
	return rest
}

func (d *Digiter1) IntToDigitsN(v int, n int) (ds []int, rest int) {
	var digit int
	ds = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if (v == 0) && (len(ds) > 0) {
			break
		}
		v, digit = d.RestDigit(v)
		ds = append(ds, digit)
	}
	rest = v
	return ds, rest
}

func (d *Digiter1) DigitsToInt(digits []int, rest int) (int, error) {
	v := digitsToIntV1(d.min, d.max, digits, rest)
	return v, nil
}
