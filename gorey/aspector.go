package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

//------------------------------------------------------------------------------

// Factorer
type Aspector struct {
	AspectRatio geom.AspectRatioF

	Content Object
}

var _ Object = Aspector{}

func (Aspector) IsObject() {}

func (v Aspector) Draw(c *cairo.Canvas, b Bounds, level int) {

	if v.Content == nil {
		return
	}

	cr := subRectByAspectRatio(b, v.AspectRatio)

	v.Content.Draw(c, cr, level+1)
}

func subRectByAspectRatio(b Bounds, aspectRatio geom.AspectRatioF) Bounds {

	var (
		dx = b.Dx()
		dy = b.Dy()
	)

	var (
		dx1 = aspectRatio.CalcDx(dy)
		dy1 = aspectRatio.CalcDy(dx)
	)

	if true {
		if dx1 < dx {
			dx = dx1
		} else {
			dy = dy1
		}
	} else { // or
		if dy1 < dy {
			dy = dy1
		} else {
			dx = dx1
		}
	}

	center := b.Center()
	f := geom.MakeFrame2(dx, dy).DivScalar(2)
	r1 := geom.Point2fToBounds(center).Grow(f)

	return r1
}

// ------------------------------------------------------------------------------
func SquareAspector(o Object) Object {
	return Aspector{
		AspectRatio: geom.MakeAspectRatioF(1, 1),
		Content:     o,
	}
}
