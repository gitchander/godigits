package random

// In interval:
// [min..max)
// [min..max-1]
func RandIntIn(r *Rand, min, max int) int {
	return min + r.Intn(max-min)
}

func RandIntsIn(r *Rand, n int, min, max int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = RandIntIn(r, min, max)
	}
	return a
}
