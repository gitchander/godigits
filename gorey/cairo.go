package gorey

import (
	"image/color"
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/utils/colorf"
)

const tau = 2.0 * math.Pi

func cairoRectangle(c *cairo.Canvas, r Bounds) {
	c.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

func cairoCircle(c *cairo.Canvas, center geom.Point2f, radius float64) {
	c.Arc(center.X, center.Y, radius, 0, tau)
}

func cairoSetSourceColor(cs *cairo.Canvas, cr color.Color) {
	colorf.CairoSetSourceColor(cs, cr)
}
