package gocairo

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

const tau = 2.0 * math.Pi

func CairoRectangle(c *cairo.Canvas, b geom.Bounds) {
	c.Rectangle(b.Min.X, b.Min.Y, b.Dx(), b.Dy())
}

func CairoCircle(c *cairo.Canvas, center geom.Point2f, radius float64) {
	c.Arc(center.X, center.Y, radius, 0, tau)
}
