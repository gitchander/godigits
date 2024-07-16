package colorf

import (
	"image/color"

	"github.com/gitchander/cairo"
)

func CairoSetSourceColor(dc *cairo.Canvas, cr color.Color) {

	if cf, ok := cr.(Colorf); ok {
		dc.SetSourceRGBA(cf.R, cf.G, cf.B, cf.A)
		return
	}

	uc := color.NRGBA64Model.Convert(cr).(color.NRGBA64)

	var (
		fR = float64(uc.R) / colorComponentMax
		fG = float64(uc.G) / colorComponentMax
		fB = float64(uc.B) / colorComponentMax
		fA = float64(uc.A) / colorComponentMax
	)

	if uc.A == colorComponentMax {
		dc.SetSourceRGB(fR, fG, fB)
	} else {
		dc.SetSourceRGBA(fR, fG, fB, fA)
	}
}
