package digits

import (
	"fmt"

	"github.com/gitchander/godigits/utils/overflows"
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
	v := d.digitsToIntV1(digits, rest)
	return v, nil
	//return d.digitsToIntV2(digits, rest)
}

func (d *Digiter1) digitsToIntV1(ds []int, rest int) int {
	v := rest
	for i := len(ds) - 1; i >= 0; i-- {
		v = (v * d.base) + ds[i]
	}
	return v
}

func (d *Digiter1) digitsToIntV2(digits []int, rest int) (int, error) {
	base := d.base
	v := rest
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		err := d.checkDigit(digit)
		if err != nil {
			return 0, err
		}

		//fmt.Println(v)

		// v = (v * d.base) + digit
		// v = (v * d.base) + digit + (k*d.base - k*d.base)
		// v = (v + k)*d.base - k*d.base + digit

		k := 0
		switch {
		case v < 0:
			k = +base
		case v > 0:
			k = -base
		}

		vb, ok := overflows.MulInt((v + k), base)
		if !ok {
			fmt.Println("digits:", digits, len(digits))
			return 0, fmt.Errorf("mul overflow: %d * %d", v, k)
		}
		s1, ok := overflows.AddInt(-(k * base), digit)
		if !ok {
			return 0, fmt.Errorf("add overflow: s1")
		}
		s2, ok := overflows.AddInt(vb, s1)
		if !ok {
			return 0, fmt.Errorf("add overflow: s2")
		}
		v = s2
	}
	return v, nil
}
