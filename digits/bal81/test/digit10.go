package main

import (
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit10 struct{}

var _ dgdr.DigitDrawer = Digit10{}

func (Digit10) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit10) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 6
		ny = 12

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.4 * w

		radius = 0.5
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	{
		c.SetLineCap(gg.LineCapButt) // !
		c.SetLineJoin(gg.LineJoinBevel)

		c.MoveTo(3, 2)
		c.LineTo(3, 10)

		c.MoveTo(1, 4)
		c.LineTo(5, 4)

		c.MoveTo(1, 6)
		c.LineTo(5, 6)

		c.MoveTo(1, 8)
		c.LineTo(5, 8)

		c.Stroke()
	}

	// round frame
	if true {
		const (
			pi     = math.Pi
			piDiv2 = pi / 2
		)

		c.DrawArc(4, 3, 1, -piDiv2, 0)
		c.LineTo(5, 9)
		c.DrawArc(4, 9, 1, 0, piDiv2)
		c.LineTo(2, 10)
		c.DrawArc(2, 9, 1, piDiv2, pi)
		c.LineTo(1, 3)
		c.DrawArc(2, 3, 1, -pi, -piDiv2)
		c.LineTo(4, 2)

		c.SetLineCap(gg.LineCapRound)
		c.SetLineJoin(gg.LineJoinRound)

		c.Stroke()
	}

	nodes := []NodeDrawers{
		{
			Positive: dgdr.MakeCircle(4, 3, radius),
			Negative: dgdr.MakeCircle(2, 3, radius),
		},
		{
			Positive: dgdr.MakeCircle(4, 5, radius),
			Negative: dgdr.MakeCircle(2, 5, radius),
		},
		{
			Positive: dgdr.MakeCircle(4, 7, radius),
			Negative: dgdr.MakeCircle(2, 7, radius),
		},
		{
			Positive: dgdr.MakeCircle(4, 9, radius),
			Negative: dgdr.MakeCircle(2, 9, radius),
		},
	}

	// nodes := []NodeDrawers{
	// 	{
	// 		Positive: MakeLine(3, 2, 5, 4),
	// 		Negative: MakeLine(3, 2, 1, 4),
	// 	},
	// 	{
	// 		Positive: MakeLine(3, 6, 5, 4),
	// 		Negative: MakeLine(3, 6, 1, 4),
	// 	},
	// 	{
	// 		Positive: MakeLine(3, 6, 5, 8),
	// 		Negative: MakeLine(3, 6, 1, 8),
	// 	},
	// 	{
	// 		Positive: MakeLine(3, 10, 5, 8),
	// 		Negative: MakeLine(3, 10, 1, 8),
	// 	},
	// }

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
}
