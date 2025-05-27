package base16

func parseBools(s string) []bool {
	var (
		rs = []rune(s)
		bs = make([]bool, 0, len(rs))
	)
	for _, r := range rs {
		switch r {
		case '0':
			bs = append(bs, false)
		case '1':
			bs = append(bs, true)
		case '-', '_': // skip
		default:
			panic("invalid bits")
		}
	}
	return bs
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
