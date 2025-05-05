package gorey

import (
	"fmt"
	"strings"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type AspectRatio struct {
	X int
	Y int
}

func MakeAspectRatio(x, y int) AspectRatio {
	g := gcd(x, y)
	if g != 0 {
		x /= g
		y /= g
	}
	return AspectRatio{
		X: x,
		Y: y,
	}
}

func (a AspectRatio) Normalize() AspectRatio {
	return MakeAspectRatio(a.X, a.Y)
}

func (a AspectRatio) CalcDx(dy float64) (dx float64, ok bool) {
	if a.Y == 0 {
		return 0, false
	}
	dx = dy * float64(a.X) / float64(a.Y)
	return dx, true
}

func (a AspectRatio) CalcDy(dx float64) (dy float64, ok bool) {
	if a.X == 0 {
		return 0, false
	}
	dy = dx * float64(a.Y) / float64(a.X)
	return dy, true
}

func (a AspectRatio) String() string {
	return fmt.Sprintf("%d:%d", a.X, a.Y)
}

func ParseAspectRatio(s string) (AspectRatio, error) {
	var zero AspectRatio
	vs := strings.Split(s, ":")
	if len(vs) != 2 {
		return zero, errParseAspectRatio(s)
	}
	x, err := parseInt(vs[0])
	if err != nil {
		return zero, errParseAspectRatio(s)
	}
	y, err := parseInt(vs[1])
	if err != nil {
		return zero, errParseAspectRatio(s)
	}
	return MakeAspectRatio(x, y), nil
}

func MustParseAspectRatio(s string) AspectRatio {
	a, err := ParseAspectRatio(s)
	if err != nil {
		panic(err)
	}
	return a
}

func errParseAspectRatio(s string) error {
	return fmt.Errorf("parse aspect ratio: invalid string (%s)", s)
}

// ------------------------------------------------------------------------------
// Factorer
type Aspector struct {
	AspectRatio float64 // dy / dx

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

func subRectByAspectRatio(b Bounds, aspectRatio float64) Bounds {

	var (
		dx = b.Dx()
		dy = b.Dy()
	)

	var (
		dx1 = dy / aspectRatio
		dy1 = dx * aspectRatio
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
		AspectRatio: 1,
		Content:     o,
	}
}
