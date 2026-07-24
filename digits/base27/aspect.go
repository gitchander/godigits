package base27

import (
	"math"

	"github.com/gitchander/godigits/geom"
)

var DigitAspectRatio = geom.MakeAspectRatioF(1, 2)

func CalcSizeY(sizeX int) (sizeY int) {
	var (
		dx = float64(sizeX)
		dy = DigitAspectRatio.CalcDy(dx)
	)
	sizeY = int(math.Ceil(dy))
	return sizeY
}
