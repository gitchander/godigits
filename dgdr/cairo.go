package dgdr

import (
	"image"
	"image/color"

	"github.com/gitchander/cairo"
	"github.com/gitchander/godigits/utils/colorf"
)

type CairoDigitDrawer interface {
	// ratio = digitHeight / digitWidth
	// digitWidth = digitHeight / ratio
	Ratio() float64

	DrawDigit(c *cairo.Canvas, x, y float64, digitHeight float64, digit int)
}

func CairoMakeDigitImage(filename string, dd CairoDigitDrawer, digitHeight float64, digit int) error {

	var (
		digitWidth = digitHeight / dd.Ratio()

		surfaceSize = image.Point{
			X: ceilInt(digitWidth),
			Y: ceilInt(digitHeight),
		}
	)

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, surfaceSize.X, surfaceSize.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer c.Destroy()

	// fill background
	if true {
		c.SetSourceRGB(1, 1, 1)
		c.Rectangle(0, 0, float64(surfaceSize.X), float64(surfaceSize.Y))
		c.Fill()
	}

	colorf.CairoSetSourceColor(c, color.Black)

	dd.DrawDigit(c, 0, 0, digitHeight, digit)

	return surface.WriteToPNG(filename)
}

func CairoMakeDigitsImageMatrix(filename string, dd CairoDigitDrawer,
	xn, yn int, digitHeight float64, digits []int) error {

	var (
		digitWidth = digitHeight / dd.Ratio()

		surfaceSize = image.Point{
			X: ceilInt(digitWidth * float64(xn)),
			Y: ceilInt(digitHeight * float64(yn)),
		}
	)

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, surfaceSize.X, surfaceSize.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer c.Destroy()

	fontSize := digitHeight * 0.1
	// err := setFont(c, fontSize)
	// if err != nil {
	// 	return err
	// }

	// fill background
	if true {
		c.SetSourceRGB(1, 1, 1)
		c.Rectangle(0, 0, float64(surfaceSize.X), float64(surfaceSize.Y))
		c.Fill()
	}

	colorf.CairoSetSourceColor(c, color.Black)

	ddt := digitDecText{
		enable: true,
		x:      0,
		y:      fontSize,
	}

	var (
		dx = digitWidth
		dy = digitHeight
	)

	cairoDrawMatrix(c, dd, xn, yn, 0, 0, dx, dy, ddt, digits)

	return surface.WriteToPNG(filename)
}

func cairoDrawMatrix(c *cairo.Canvas, dd CairoDigitDrawer, xn, yn int,
	x0, y0 float64, dx, dy float64, ddt digitDecText, digits []int) {

	for y := 0; y < yn; y++ {
		for x := 0; x < xn; x++ {
			var (
				x1 = x0 + float64(x)*dx
				y1 = y0 + float64(y)*dy
			)
			if true {
				if ((x + y) % 2) == 0 {
					c.SetSourceRGB(1, 1, 1)
				} else {
					c.SetSourceRGB(0.7, 0.9, 1)
				}
				c.Rectangle(x1, y1, dx, dy)
				c.Fill()
				colorf.CairoSetSourceColor(c, color.Black)
			}
			if len(digits) > 0 {
				digit := digits[0]
				digits = digits[1:]

				if ddt.enable {
					c.MoveTo(x1+ddt.x, y1+ddt.y)
					c.ShowText(formatInt(digit))
				}

				dd.DrawDigit(c, x1, y1, dy, digit)
			}
		}
	}
}
