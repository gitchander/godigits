package digits

func CalcDigits(v int, min, max int, ds []int) (rest int) {
	var digit int
	for i := range ds {
		v, digit = RestDigit(v, min, max)
		ds[i] = digit
	}
	rest = v
	return rest
}

func CalcDigitsN(v int, min, max int, n int) (ds []int, rest int) {
	var digit int
	ds = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if (v == 0) && (len(ds) > 0) {
			break
		}
		v, digit = RestDigit(v, min, max)
		ds = append(ds, digit)
	}
	rest = v
	return ds, rest
}
