package bal81

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit4 struct{}

var _ dgdr.DigitDrawer = Digit4{}

func (Digit4) Width(height float64) (width float64) {
	width = height
	return
}

func (d Digit4) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 8
		ny = 8

		dx = digitHeight / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.4 * dx
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	//center := Pt2f(2, 2)

	var (
		x1, x2, x3 float64 = 2, 4, 6
		y1, y2, y3 float64 = 2, 4, 6
	)

	bs := CalcDigitsBal3(digit)

	if true {
		c.MoveTo(x1, y1)
		c.LineTo(x3, y3)

		c.MoveTo(x1, y3)
		c.LineTo(x3, y1)
	} else {
		c.MoveTo(middle(x1, x2), middle(y1, y2))
		c.LineTo(middle(x2, x3), middle(y2, y3))

		c.MoveTo(middle(x1, x2), middle(y2, y3))
		c.LineTo(middle(x2, x3), middle(y1, y2))
	}

	c.Stroke()

	//--------------------------------------------------------------------------

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(5, 3, 4, 2),
	// 		Negative: linePoints(3, 3, 4, 2),
	// 	},
	// 	{
	// 		Positive: linePoints(5, 5, 6, 4),
	// 		Negative: linePoints(5, 3, 6, 4),
	// 	},
	// 	{
	// 		Positive: linePoints(3, 5, 4, 6),
	// 		Negative: linePoints(5, 5, 4, 6),
	// 	},
	// 	{
	// 		Positive: linePoints(3, 3, 2, 4),
	// 		Negative: linePoints(3, 5, 2, 4),
	// 	},
	// }

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(6, 2, 4, 2),
	// 		Negative: linePoints(2, 6, 4, 6),
	// 	},
	// 	{
	// 		Positive: linePoints(6, 2, 6, 4),
	// 		Negative: linePoints(2, 6, 2, 4),
	// 	},
	// 	{
	// 		Positive: linePoints(6, 6, 6, 4),
	// 		Negative: linePoints(2, 2, 2, 4),
	// 	},
	// 	{
	// 		Positive: linePoints(6, 6, 4, 6),
	// 		Negative: linePoints(2, 2, 4, 2),
	// 	},
	// }

	nodes := []Node{
		{
			Positive: linePoints(5, 3, 4, 2),
			Negative: linePoints(3, 5, 4, 6),
		},
		{
			Positive: linePoints(5, 3, 6, 4),
			Negative: linePoints(3, 5, 2, 4),
		},
		{
			Positive: linePoints(5, 5, 6, 4),
			Negative: linePoints(3, 3, 2, 4),
		},
		{
			Positive: linePoints(5, 5, 4, 6),
			Negative: linePoints(3, 3, 4, 2),
		},
	}

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(6, 2, 4, 2),
	// 		Negative: linePoints(5, 3, 4, 2),
	// 	},
	// 	{
	// 		Positive: linePoints(6, 2, 6, 4),
	// 		Negative: linePoints(5, 3, 6, 4),
	// 	},
	// 	{
	// 		Positive: linePoints(6, 6, 6, 4),
	// 		Negative: linePoints(5, 5, 6, 4),
	// 	},
	// 	{
	// 		Positive: linePoints(6, 6, 4, 6),
	// 		Negative: linePoints(5, 5, 4, 6),
	// 	},
	// }

	//--------------------------------------------------------------------------

	n := minInt(len(nodes), len(bs))
	for i, node := range nodes[:n] {
		switch bs[i] {
		case 1:
			{
				var (
					a1 = node.Positive[0]
					a2 = node.Positive[1]
				)
				c.MoveTo(a1.X, a1.Y)
				c.LineTo(a2.X, a2.Y)
			}
		case -1:
			{
				var (
					a1 = node.Negative[0]
					a2 = node.Negative[1]
				)
				c.MoveTo(a1.X, a1.Y)
				c.LineTo(a2.X, a2.Y)
			}
		}
	}

	c.Stroke()
}
