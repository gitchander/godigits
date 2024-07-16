package utils

func SerialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

// min <= value < max
func MakeInts(min, max, step int) []int {
	var (
		n  = (max - min) / step
		as = make([]int, 0, n)
	)
	for a := min; a < max; a += step {
		as = append(as, a)
	}
	return as
}

func MakeIntsByInterval(min, max int) []int {
	var (
		n  = max - min
		as = make([]int, n)
	)
	for i := 0; i < n; i++ {
		as[i] = min + i
	}
	return as
}
