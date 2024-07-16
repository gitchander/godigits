package main

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/digits/bal81"
)

type Digit6 struct{}

var _ dgdr.DigitDrawer = Digit6{}

func (Digit6) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit6) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 4
		ny = 8

		dx = digitHeight / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.35 * dx
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	if true {
		c.MoveTo(2, 2)
		c.LineTo(2, 6)

		c.MoveTo(1, 2)
		c.LineTo(3, 2)

		c.MoveTo(1, 6)
		c.LineTo(3, 6)
	} else {
		c.MoveTo(2, 1)
		c.LineTo(2, 7)

		// c.MoveTo(2, 2)
		// c.LineTo(2, 6)
	}

	//--------------------------------------------------------------------------

	// nodes = []Node{
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 3), Pt2f(3, 2)},
	// 		Negative: []Point2f{Pt2f(2, 3), Pt2f(1, 2)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 3), Pt2f(3, 4)},
	// 		Negative: []Point2f{Pt2f(2, 3), Pt2f(1, 4)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 5), Pt2f(3, 4)},
	// 		Negative: []Point2f{Pt2f(2, 5), Pt2f(1, 4)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 5), Pt2f(3, 6)},
	// 		Negative: []Point2f{Pt2f(2, 5), Pt2f(1, 6)},
	// 	},
	// }

	// nodes = []Node{
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 3), Pt2f(3, 2)},
	// 		Negative: []Point2f{Pt2f(2, 3), Pt2f(1, 2)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 4), Pt2f(3, 3)},
	// 		Negative: []Point2f{Pt2f(2, 4), Pt2f(1, 3)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 4), Pt2f(3, 5)},
	// 		Negative: []Point2f{Pt2f(2, 4), Pt2f(1, 5)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 5), Pt2f(3, 6)},
	// 		Negative: []Point2f{Pt2f(2, 5), Pt2f(1, 6)},
	// 	},
	// }

	// nodes = []Node{
	// 	{
	// 		Positive: []Point2f{Pt2f(3, 2), Pt2f(3, 3)},
	// 		Negative: []Point2f{Pt2f(1, 2), Pt2f(1, 3)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 4), Pt2f(3, 3)},
	// 		Negative: []Point2f{Pt2f(2, 4), Pt2f(1, 3)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(2, 4), Pt2f(3, 5)},
	// 		Negative: []Point2f{Pt2f(2, 4), Pt2f(1, 5)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(3, 6), Pt2f(3, 5)},
	// 		Negative: []Point2f{Pt2f(1, 6), Pt2f(1, 5)},
	// 	},
	// }

	// nodes := []Node{
	// 	{
	// 		Positive: []Point2f{Pt2f(3, 3), Pt2f(3, 2)},
	// 		Negative: []Point2f{Pt2f(1, 3), Pt2f(1, 2)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(3, 3), Pt2f(2, 3)},
	// 		Negative: []Point2f{Pt2f(1, 3), Pt2f(2, 3)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(3, 5), Pt2f(2, 5)},
	// 		Negative: []Point2f{Pt2f(1, 5), Pt2f(2, 5)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(3, 5), Pt2f(3, 6)},
	// 		Negative: []Point2f{Pt2f(1, 5), Pt2f(1, 6)},
	// 	},
	// }

	nodes := []Node{
		{
			Positive: []Point2f{Pt2f(3, 2), Pt2f(3, 3)},
			Negative: []Point2f{Pt2f(1, 2), Pt2f(1, 3)},
		},
		{
			Positive: []Point2f{Pt2f(2, 4), Pt2f(3, 3)},
			Negative: []Point2f{Pt2f(2, 4), Pt2f(1, 3)},
		},
		{
			Positive: []Point2f{Pt2f(2, 4), Pt2f(3, 5)},
			Negative: []Point2f{Pt2f(2, 4), Pt2f(1, 5)},
		},
		{
			Positive: []Point2f{Pt2f(3, 6), Pt2f(3, 5)},
			Negative: []Point2f{Pt2f(1, 6), Pt2f(1, 5)},
		},
	}

	//--------------------------------------------------------------------------

	bs := bal81.CalcDigitsBal3(digit)

	n := minInt(len(nodes), len(bs))

	for i := 0; i < n; i++ {
		n := nodes[i]
		if bs[i] == 1 {
			var (
				a1 = n.Positive[0]
				a2 = n.Positive[1]
			)
			c.MoveTo(a1.X, a1.Y)
			c.LineTo(a2.X, a2.Y)
		} else if bs[i] == -1 {
			var (
				a1 = n.Negative[0]
				a2 = n.Negative[1]
			)
			c.MoveTo(a1.X, a1.Y)
			c.LineTo(a2.X, a2.Y)
		}
	}

	c.Stroke()
}
