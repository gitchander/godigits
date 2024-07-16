package main

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit7 struct{}

var _ dgdr.DigitDrawer = Digit7{}

func (Digit7) Width(height float64) (width float64) {
	width = height
	return
}

func (d Digit7) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 4
		ny = 4

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.25 * w
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineCap(gg.LineCapRound)
	c.SetLineJoin(gg.LineJoinRound)
	c.SetLineWidth(lineWidth)

	var (
		x1, x2, x3 float64 = 1, 2, 3
		y1, y2, y3 float64 = 1, 2, 3
	)

	if true {
		var (
			dx = 0.0
			dy = 0.0

			// dx = 0.25
			// dy = 0.25

			// dx = 0.3
			// dy = 0.3
		)

		c.MoveTo(x2, y1-dy)
		c.LineTo(x2, y3+dy)

		c.MoveTo(x1-dx, y2)
		c.LineTo(x3+dx, y2)

		c.Stroke()
	}

	// var (
	// 	// dx = 0.5
	// 	// dy = 0.5

	// 	dx = 0.0
	// 	dy = 0.0

	// 	nodes = []NodeDrawers{
	// 		{
	// 			Positive: MakeLine(x2, y1, x3, y2),
	// 			Negative: MakeLine(x2, y2, x3-dx, y1+dy),
	// 		},
	// 		{
	// 			Positive: MakeLine(x3, y2, x2, y3),
	// 			Negative: MakeLine(x2, y2, x3-dx, y3-dy),
	// 		},
	// 		{
	// 			Positive: MakeLine(x2, y3, x1, y2),
	// 			Negative: MakeLine(x2, y2, x1+dx, y3-dy),
	// 		},
	// 		{
	// 			Positive: MakeLine(x1, y2, x2, y1),
	// 			Negative: MakeLine(x2, y2, x1+dx, y1+dy),
	// 		},
	// 	}
	// )

	// var (
	// 	dx = invSqrt2
	// 	dy = invSqrt2

	// 	nodes = []NodeDrawers{
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, -tau/4, 0),
	// 			Negative: MakeLine(x2, y2, x2+dx, y2-dy),
	// 		},
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, 0, tau/4),
	// 			Negative: MakeLine(x2, y2, x2+dx, y2+dy),
	// 		},
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, tau/4, tau/2),
	// 			Negative: MakeLine(x2, y2, x2-dx, y2+dy),
	// 		},
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, -tau/2, -tau/4),
	// 			Negative: MakeLine(x2, y2, x2-dx, y2-dy),
	// 		},
	// 	}
	// )

	var (
		dx     = 0.5
		dy     = 0.5
		radius = 0.2

		nodes = []NodeDrawers{
			{
				Positive: dgdr.MakeArc(x2, y2, 1, -tau/4, 0),
				Negative: dgdr.MakeCircle(x2+dx, y2-dy, radius),
			},
			{
				Positive: dgdr.MakeArc(x2, y2, 1, 0, tau/4),
				Negative: dgdr.MakeCircle(x2+dx, y2+dy, radius),
			},
			{
				Positive: dgdr.MakeArc(x2, y2, 1, tau/4, tau/2),
				Negative: dgdr.MakeCircle(x2-dx, y2+dy, radius),
			},
			{
				Positive: dgdr.MakeArc(x2, y2, 1, -tau/2, -tau/4),
				Negative: dgdr.MakeCircle(x2-dx, y2-dy, radius),
			},
		}
	)

	// var (
	// 	nodes = []NodeDrawers{
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, -tau/4, 0),
	// 			Negative: MakeArc(x3, y1, 1, tau/4, tau/2),
	// 		},
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, 0, tau/4),
	// 			Negative: MakeArc(x3, y3, 1, -tau/2, -tau/4),
	// 		},
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, tau/4, tau/2),
	// 			Negative: MakeArc(x1, y3, 1, -tau/4, 0),
	// 		},
	// 		{
	// 			Positive: MakeArc(x2, y2, 1, -tau/2, -tau/4),
	// 			Negative: MakeArc(x1, y1, 1, 0, tau/4),
	// 		},
	// 	}
	// )

	bs := make([]int, len(nodes))
	calcDigits(digit, bs)

	n := minInt(len(nodes), len(bs))
	for i, node := range nodes[:n] {
		switch bs[i] {
		case 1:
			node.Positive.DrawGG(c)
		case -1:
			node.Negative.DrawGG(c)
		}
	}

	c.Stroke()
}
