package gorey

import (
	"fmt"
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
	ivl "github.com/gitchander/godigits/utils/interval"
)

// https://en.wikipedia.org/wiki/Lemniscate_of_Bernoulli

// Lemniscate of Bernoulli
func lemniscate(radius, theta float64) (float64, bool) {
	x := square(radius) * math.Cos(2*theta)
	if x < 0 {
		return 0, false
	}
	return math.Sqrt(x), true
}

func drawLemniscateV1(c *cairo.Canvas, center geom.Point2f, radius float64) {
	n := 120
	angle := 0.0
	deltaAngle := tau / float64(n)
	for i := 0; i <= n; i++ {
		r, ok := lemniscate(radius, angle)
		if ok {
			p := geom.PolarToCartesian(geom.ShPolar(r, angle))
			p = center.Add(p)
			c.LineTo(p.X, p.Y)
		}
		angle += deltaAngle
	}
	c.Stroke()
}

func drawLemniscateV2(c *cairo.Canvas, center geom.Point2f, radius float64) {

	const piDiv4 = math.Pi / 4

	var (
		ps  = []float64{0, math.Pi}
		ivs = make([]ivl.IntervalFloat, len(ps))
	)
	for i, p := range ps {
		ivs[i] = ivl.IntervalFloat{
			Min: p - piDiv4,
			Max: p + piDiv4,
		}
	}

	n := 81

	for _, iv := range ivs {
		vr := ivl.NewIntervalValuer(iv, n)
		for i := 0; i < n; i++ {
			theta := vr.IndexToValue(i)
			r, _ := lemniscate(radius, theta)
			p := geom.PolarToCartesian(geom.ShPolar(r, theta))
			p = p.Add(center)

			if i == 0 {
				c.MoveTo(p.X, p.Y)
			} else {
				c.LineTo(p.X, p.Y)
			}
		}
	}
	c.Stroke()
}

func drawLemniscateV3(c *cairo.Canvas, center geom.Point2f, radius float64) {

	c.SetLineCap(cairo.LINE_CAP_ROUND)

	const angleMax = (math.Pi / 4)

	deltaAngle := angleMax * 0.1
	deltaAngleMin := 0.001
	deltaAngleFactor := 0.95

	// deltaAngle := angleMax * 0.2
	// deltaAngleMin := 0.001
	// deltaAngleFactor := 0.85

	// deltaAngle := angleMax * 0.15
	// deltaAngleMin := 0.001
	// deltaAngleFactor := 0.88

	angle := 0.0

	var as []geom.Point2f
	for {
		stop := false
		if angle > angleMax {
			angle = angleMax
			stop = true
		}

		r, _ := lemniscate(radius, angle)

		p := geom.PolarToCartesian(geom.ShPolar(r, angle))
		as = append(as, p)

		angle += deltaAngle
		deltaAngle *= deltaAngleFactor
		if deltaAngle < deltaAngleMin {
			deltaAngle = deltaAngleMin
		}

		if deltaAngle < 1e-13 {
			break
		}
		if stop {
			break
		}
	}

	if false {
		fmt.Println(len(as))
	}

	for i, a := range as {
		p := geom.Point2f{
			X: center.X + a.X,
			Y: center.Y + a.Y,
		}
		if i == 0 {
			c.MoveTo(p.X, p.Y)
		} else {
			c.LineTo(p.X, p.Y)
		}
	}

	for i, a := range as {
		p := geom.Point2f{
			X: center.X - a.X,
			Y: center.Y + a.Y,
		}
		if i == 0 {
			c.MoveTo(p.X, p.Y)
		} else {
			c.LineTo(p.X, p.Y)
		}
	}

	for i, a := range as {
		p := geom.Point2f{
			X: center.X + a.X,
			Y: center.Y - a.Y,
		}
		if i == 0 {
			c.MoveTo(p.X, p.Y)
		} else {
			c.LineTo(p.X, p.Y)
		}
	}

	for i, a := range as {
		p := geom.Point2f{
			X: center.X - a.X,
			Y: center.Y - a.Y,
		}
		if i == 0 {
			c.MoveTo(p.X, p.Y)
		} else {
			c.LineTo(p.X, p.Y)
		}
	}

	c.Stroke()

	if true {
		c.SetSourceRGB(1, 0, 0)
		for _, a := range as {
			p := center.Add(a)
			c.Arc(p.X, p.Y, 0.005, 0, tau)
			c.Fill()
		}
	}
}

func drawLemniscate(c *cairo.Canvas, r geom.Rectangle2f) {

	const lineWidthRel = 5.0

	var (
		vm = vmin(r)

		lineWidthAbs = lineWidthRel * vm
	)

	r1 := r.Shrink(geom.MakeFrame1(lineWidthAbs / 2))

	const aspectRatio = 1.0 / (2.0 * math.Sqrt2)

	r2 := subRectByAspectRatio(r1, aspectRatio)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(3))
		cairoRectangle(c, r1)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(4))
		cairoRectangle(c, r2)
		c.Fill()
	}

	center := r2.Center()

	c.SetSourceRGB(0, 0, 0)
	c.SetLineWidth(lineWidthAbs)

	radius := r2.Dy() * math.Sqrt2

	drawLemniscateV2(c, center, radius)
}
