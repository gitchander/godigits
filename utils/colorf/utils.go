package colorf

func not(b bool) bool {
	return !b
}

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/Nibble
func byteToNibbles(b byte) (hi, lo byte) {
	hi = b >> 4
	lo = b & 0xf
	return
}

func nibblesToByte(hi, lo byte) (b byte) {
	b |= hi << 4
	b |= lo & 0xf
	return
}

//------------------------------------------------------------------------------

func clampFloat64(a float64, min, max float64) float64 {
	if a < min {
		a = min
	}
	if a > max {
		a = max
	}
	return a
}
