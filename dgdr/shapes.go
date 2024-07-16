package dgdr

import (
	"math"

	"github.com/fogleman/gg"
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/utils/gocairo"
)

const tau = 2 * math.Pi

//------------------------------------------------------------------------------

type Line struct {
	A, B geom.Point2f
}

var _ DrawerCairoGG = Line{}

func (l Line) DrawGG(c *gg.Context) {
	c.MoveTo(l.A.X, l.A.Y)
	c.LineTo(l.B.X, l.B.Y)
	c.Stroke()
}

func (l Line) DrawCairo(c *cairo.Canvas) {
	c.MoveTo(l.A.X, l.A.Y)
	c.LineTo(l.B.X, l.B.Y)
	c.Stroke()
}

func MakeLine(x1, y1 float64, x2, y2 float64) Line {
	return Line{
		A: geom.Pt2f(x1, y1),
		B: geom.Pt2f(x2, y2),
	}
}

//------------------------------------------------------------------------------

type Circle struct {
	Center geom.Point2f
	Radius float64
}

var _ DrawerCairoGG = Circle{}

func (x Circle) DrawGG(c *gg.Context) {
	c.DrawCircle(x.Center.X, x.Center.Y, x.Radius)
	c.Fill()
}

func (x Circle) DrawCairo(c *cairo.Canvas) {
	gocairo.CairoCircle(c, x.Center, x.Radius)
	c.Fill()
}

func MakeCircle(x, y float64, r float64) Circle {
	return Circle{
		Center: geom.Pt2f(x, y),
		Radius: r,
	}
}

//------------------------------------------------------------------------------

type Arc struct {
	Center geom.Point2f
	Radius float64
	Angle1 float64
	Angle2 float64
}

var _ DrawerCairoGG = Arc{}

func (v Arc) DrawGG(c *gg.Context) {
	c.DrawArc(v.Center.X, v.Center.Y, v.Radius, v.Angle1, v.Angle2)
	c.Stroke()
}

func (v Arc) DrawCairo(c *cairo.Canvas) {
	c.Arc(v.Center.X, v.Center.Y, v.Radius, v.Angle1, v.Angle2)
	c.Stroke()
}

func MakeArc(x, y float64, r float64, angle1, angle2 float64) Arc {
	return Arc{
		Center: geom.Pt2f(x, y),
		Radius: r,
		Angle1: angle1,
		Angle2: angle2,
	}
}
