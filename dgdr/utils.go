package dgdr

import (
	"math"
	"strconv"
)

func ceilInt(x float64) int {
	return int(math.Ceil(x))
}

func formatInt(a int) string {
	return strconv.Itoa(a)
}
