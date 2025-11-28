package base27

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/utils/digits"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcTrits(v int, ds []int) {
	const (
		min = -1
		max = +1
	)
	rd := digits.MustNewRestDigiter(min, max)
	digits.CalcDigits(rd, v, ds)
}

func calcTritsBase27(v int) []int {
	trits := make([]int, 3)
	calcTrits(v, trits)
	return trits
}

//------------------------------------------------------------------------------

type Coord = geom.Point2f

type geomDrawer struct {
	c *gg.Context
}

func newGeomDrawer(c *gg.Context) *geomDrawer {
	return &geomDrawer{c: c}
}

func (d *geomDrawer) DrawLine(a, b Coord) {
	d.c.MoveTo(a.X, a.Y)
	d.c.LineTo(b.X, b.Y)
}

func (d *geomDrawer) DrawLineTo(b Coord) {
	d.c.LineTo(b.X, b.Y)
}

func (d *geomDrawer) DrawCircle(center Coord, radius float64) {
	d.c.DrawCircle(center.X, center.Y, radius)
}
