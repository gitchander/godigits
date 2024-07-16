package bal81

import (
	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/utils/digits"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type nodeDrawers struct {
	positive dgdr.DrawerCairoGG
	negative dgdr.DrawerCairoGG
}

func CalcDigitsBal3(v int) []int {
	const (
		min = -1 // digit min
		max = +1 // digit max
	)
	ds := make([]int, 4)
	digits.CalcDigits(v, min, max, ds)
	return ds
}

// func calcDigits(v int, ds []int) []int {
// 	const (
// 		min = -1
// 		max = +1
// 	)
// 	digits.CalcDigits(v, min, max, ds)
// 	return ds
// }
