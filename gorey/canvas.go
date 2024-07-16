package gorey

import (
	"image"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type ICanvas struct {
	c *cairo.Canvas

	p0       geom.Point2f
	cellSize geom.Point2f
}

func NewICanvas(c *cairo.Canvas, r geom.Rectangle2f, nx, ny int) *ICanvas {
	return &ICanvas{
		c:  c,
		p0: r.Min,
		cellSize: geom.Point2f{
			X: r.Dx() / float64(nx),
			Y: r.Dy() / float64(ny),
		},
	}
}

func (d *ICanvas) pt(x, y int) geom.Point2f {
	return geom.Point2f{
		X: d.p0.X + (d.cellSize.X * float64(x)),
		Y: d.p0.Y + (d.cellSize.Y * float64(y)),
	}
}

func (d *ICanvas) MoveTo(xi, yi int) {
	p := d.pt(xi, yi)
	d.c.MoveTo(p.X, p.Y)
}

func (d *ICanvas) LineTo(xi, yi int) {
	p := d.pt(xi, yi)
	d.c.LineTo(p.X, p.Y)
}

func (d *ICanvas) Circle(x, y int, radius float64) {
	p := d.pt(x, y)
	cairoCircle(d.c, p, radius)
}

func (d *ICanvas) BesierCubic(ps [4]image.Point) {
	as := make([]geom.Point2f, len(ps))
	for i, p := range ps {
		as[i] = d.pt(p.X, p.Y)
	}
	BesierCubic(d.c, as[0], as[1], as[2], as[3])
}

func (d *ICanvas) BesierQuad(ps [3]image.Point) {
	as := make([]geom.Point2f, len(ps))
	for i, p := range ps {
		as[i] = d.pt(p.X, p.Y)
	}
	BesierQuad(d.c, as[0], as[1], as[2])
}
