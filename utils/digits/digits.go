package digits

import (
	"fmt"

	"github.com/gitchander/godigits/utils/overflows"
)

func CalcDigits(rd RestDigiter, v int, ds []int) (rest int) {
	rest = v
	var digit int
	for i := range ds {
		rest, digit = rd.RestDigit(rest)
		ds[i] = digit
	}
	return rest
}

func CalcDigitsN(rd RestDigiter, v int, n int) (ds []int, rest int) {
	ds = make([]int, 0, n)
	rest = v
	var digit int
	for i := 0; i < n; i++ {
		if (rest == 0) && (len(ds) > 0) {
			break
		}
		rest, digit = rd.RestDigit(rest)
		ds = append(ds, digit)
	}
	return ds, rest
}

//------------------------------------------------------------------------------

func DigitsToInt(min, max int, digits []int, rest int) int {
	return digitsToIntV1(min, max, digits, rest)
	//return digitsToIntV2(min, max, digits, rest)
}

func digitsToIntV1(min, max int, digits []int, rest int) int {
	base := max - min + 1
	v := rest
	for i := len(digits) - 1; i >= 0; i-- {
		v = (v * base) + digits[i]
	}
	return v
}

func digitsToIntV2(min, max int, digits []int, rest int) int {
	base := max - min + 1
	v := rest
	p := 1 // base^0
	for i := 0; i < len(digits); i++ {
		v += digits[i] * p // digit * base^i
		p *= base
	}
	return v
}

func checkDigit(min, max int, digit int) error {
	if (min <= digit) && (digit <= max) {
		return nil
	}
	return fmt.Errorf("invalid digit %d, want interval [%d .. %d]", digit, min, max)
}

func digitsToIntV3(min, max int, digits []int, rest int) (int, error) {
	base := max - min + 1
	v := rest
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		err := checkDigit(min, max, digit)
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
