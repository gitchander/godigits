package dig12

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawer interface {
	DrawDigit(c *gg.Context, b geom.Bounds, digit int)
}
