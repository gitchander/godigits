package dgdr

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
)

type DigitDrawerB interface {
	DrawDigit(c *gg.Context, b geom.Bounds, digit int)
}

func DrawMatrixDDB(c *gg.Context, d DigitDrawerB, nX, nY int,
	digitSize image.Point, digits []int) {

	var (
		dsX = digitSize.X
		dsY = digitSize.Y
	)

	for y := 0; y < nY; y++ {
		for x := 0; x < nX; x++ {
			b := geom.MakeBounds(
				float64((x+0)*dsX), float64((y+0)*dsY),
				float64((x+1)*dsX), float64((y+1)*dsY),
			)
			if true {
				if ((x + y) % 2) == 0 {
					c.SetRGB(0.7, 0.9, 1.0)
				} else {
					c.SetRGB(1, 1, 1)
				}
				drawBounds(c, b)
				c.Fill()
			}
			if len(digits) > 0 {
				digit := digits[0]
				digits = digits[1:]

				c.SetRGB(0, 0, 0)
				c.DrawString(fmt.Sprintf("%d", digit), b.Min.X, b.Min.Y+c.FontHeight())

				c.SetRGB(0, 0, 0)
				d.DrawDigit(c, b, digit)
			}
		}
	}
}

func drawBounds(c *gg.Context, b geom.Bounds) {
	c.DrawRectangle(b.Min.X, b.Min.Y, b.Max.X, b.Max.Y)
}
