package hexagone

import (
	"github.com/gitchander/cairo"
	"github.com/gitchander/godigits/geom"
)

func CairoDrawPolygon(c *cairo.Canvas, ps []geom.Point2f) {
	n := len(ps)
	if n > 0 {
		p := ps[0]
		c.MoveTo(p.X, p.Y)
	}
	for i := 1; i < n; i++ {
		p := ps[i]
		c.LineTo(p.X, p.Y)
	}
	if n > 2 {
		p := ps[0]
		c.LineTo(p.X, p.Y)
	}
}

func CairoDrawSegment(c *cairo.Canvas, a, b geom.Point2f, lineWidth float64, angle float64) {

	ps := SegmentPolygone(a, b, lineWidth, angle)
	CairoDrawPolygon(c, ps)

	c.Fill()
}
