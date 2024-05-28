package numbers

// number:          ...(a-2),(a-1)|a...(b-1)|b,(b+1),(b+2)...
// rest:   ...|   -2    |   -1    |    0    |    1    |    2    |...
// digit   ...|a...(b-1)|a...(b-1)|a...(b-1)|a...(b-1)|a...(b-1)|...

func RestDigit(b Base, x int) (rest, digit int) {
	return restDigitV1(b, x)
}

func restDigitV1(b Base, x int) (rest, digit int) {

	var (
		min = b.Min
		max = b.Max - 1
		bw  = b.Width()
	)

	if x < min {
		rest, digit = quoRem(x-max, bw)
		digit += max
	} else {
		rest, digit = quoRem(x-min, bw)
		digit += min
	}

	return
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
