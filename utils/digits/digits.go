package digits

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
	rest = v
	var digit int
	ds = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if (rest == 0) && (len(ds) > 0) {
			break
		}
		rest, digit = rd.RestDigit(rest)
		ds = append(ds, digit)
	}
	return ds, rest
}
