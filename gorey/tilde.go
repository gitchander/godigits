package gorey

import (
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Tilde struct {
	Content Object
}

var _ Object = Tilde{}

func (Tilde) IsObject() {}

func (v Tilde) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	const lineWidth = 0.05

	c.SetSourceRGB(0, 0, 0)
	c.SetLineWidth(lineWidth)

	//c.SetLineJoin(cairo.LINE_JOIN_BEVEL)
	c.SetLineJoin(cairo.LINE_JOIN_MITER)
	//c.SetLineJoin(cairo.LINE_JOIN_ROUND)

	m := 0.04
	d := 0.1
	k := d / math.Sqrt2

	y1 := m
	xm := (1-2*m)/2 + m

	c.CurveTo(xm-2*k, y1+k, xm-k, y1+2*k, xm, y1+k)
	c.CurveTo(xm, y1+k, xm+k, y1, xm+2*k, y1+k)

	c.Stroke()

	// Content
	// {
	// 	var (
	// 		width = 0.75
	// 		x1    = (1 - width) / 2
	// 		y1    = 0.2
	// 	)

	// 	// c.SetSourceRGBA(1, 0, 0, 0.5)
	// 	// drawContentArea(c, x1, y1, width)
	// }

	cr := geom.Rectangle2f{}

	if v.Content != nil {
		v.Content.Draw(c, cr, level+1)
	}
}
