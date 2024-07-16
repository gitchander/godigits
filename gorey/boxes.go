package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type HBox struct {
	Objects []Object
}

var _ Object = HBox{}

func (HBox) IsObject() {}

func (v HBox) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	if len(v.Objects) == 0 {
		return
	}

	var (
		margin = geom.MakeFrame1(0.05)

		osep = 0.03
	)

	var (
		w = minFloat64(r.Dx(), r.Dy())

		mr = r.Shrink(margin.MulScalar(w))
	)

	var (
		sw = osep * w

		dx          = (mr.Dx() + sw) / float64(len(v.Objects))
		objectWidth = dx - sw
	)

	x := mr.Min.X

	y1 := mr.Min.Y
	y2 := mr.Max.Y

	for _, object := range v.Objects {

		cr := geom.Rectangle2f{
			Min: geom.Pt2f(x, y1),
			Max: geom.Pt2f(x+objectWidth, y2),
		}

		if object != nil {
			object.Draw(c, cr, level+1)
		}

		x += dx
	}
}

//------------------------------------------------------------------------------
// Couple

// +-------+---+
// |       | B |
// +-------+---+
// |       |   |
// |   A   |   |
// |       |   |
// +-------+---+
// |       |   |
// |<--a-->|   |
//
//	-->| b |<--

// type DiagonalBox struct {
// 	Ratio float64
// 	A, B  Object
// }

// var _ Object = DiagonalBox{}

// func (DiagonalBox) IsObject() {}

// func (v DiagonalBox) Draw(c *cairo.Canvas) {

// 	size := Pt2f(1, 1)

// 	w := minFloat64(size.X, size.Y)

// 	var (
// 		a = w * (v.Ratio / (1 + v.Ratio))
// 		b = w * (1 / (1 + v.Ratio))
// 	)

// 	var (
// 		xa, xb = 0.0, a
// 		ya, yb = b, 0.0
// 	)

// 	drawAreas := false

// 	if drawAreas {
// 		c.SetSourceRGBA(1, 0, 0, 0.5)
// 		drawContentArea(c, xa, ya, a)
// 	}
// 	drawContent(c, v.A, xa, ya, a)

// 	if drawAreas {
// 		c.SetSourceRGBA(0, 0, 1, 0.5)
// 		drawContentArea(c, xb, yb, b)
// 	}
// 	drawContent(c, v.B, xb, yb, b)
// }

// ------------------------------------------------------------------------------
