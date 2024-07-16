package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/utils/digits"
)

const (
	tau      = 2 * math.Pi
	invSqrt2 = 1.0 / math.Sqrt2
)

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func RadToDeg(rad float64) float64 {
	return rad * 180 / math.Pi
}

type Node struct {
	Positive []geom.Point2f
	Negative []geom.Point2f
}

type NodeDrawers struct {
	Positive dgdr.DrawerGG
	Negative dgdr.DrawerGG
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randInts(n int, min, max int) []int {
	r := newRandNow()
	a := make([]int, n)
	for i := range a {
		a[i] = min + r.Intn(max-min)
	}
	return a
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func linePoints(x1, y1 float64, x2, y2 float64) []geom.Point2f {
	return []Point2f{
		Pt2f(x1, y1),
		Pt2f(x2, y2),
	}
}

//------------------------------------------------------------------------------

func drawPolyLineGG(c *gg.Context, ps []geom.Point2f) {
	n := len(ps)
	if n > 0 {
		p := ps[0]
		c.MoveTo(p.X, p.Y)
	}
	for i := 1; i < n; i++ {
		p := ps[i]
		c.LineTo(p.X, p.Y)
	}
}

func drawPolyLineCairo(c *cairo.Canvas, ps []geom.Point2f) {
	n := len(ps)
	if n > 0 {
		p := ps[0]
		c.MoveTo(p.X, p.Y)
	}
	for i := 1; i < n; i++ {
		p := ps[i]
		c.LineTo(p.X, p.Y)
	}
}

func calcDigits(v int, ds []int) []int {
	const (
		min = -1
		max = +1
	)
	digits.CalcDigits(v, min, max, ds)
	return ds
}

//------------------------------------------------------------------------------

type Point2f = geom.Point2f

func Pt2f(x, y float64) Point2f {
	return Point2f{
		X: x,
		Y: y,
	}
}

//------------------------------------------------------------------------------
