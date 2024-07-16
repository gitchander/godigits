package gocairo

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

const tau = 2.0 * math.Pi

func CairoRectangle(c *cairo.Canvas, r geom.Rectangle2f) {
	c.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

func CairoCircle(c *cairo.Canvas, center geom.Point2f, radius float64) {
	c.Arc(center.X, center.Y, radius, 0, tau)
}
