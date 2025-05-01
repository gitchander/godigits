package bal81

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type Digit9 struct{}

var _ dgdr.DigitDrawer = Digit9{}

func (Digit9) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit9) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 4
		ny = 8

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.4 * w
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineCap(gg.LineCapRound)
	c.SetLineJoin(gg.LineJoinRound)
	c.SetLineWidth(lineWidth)

	c.MoveTo(2, 2)
	c.LineTo(2, 6)

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(2, 2, 3, 3),
	// 		Negative: linePoints(2, 6, 1, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 3, 3),
	// 		Negative: linePoints(2, 4, 1, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 3, 5),
	// 		Negative: linePoints(2, 4, 1, 3),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 6, 3, 5),
	// 		Negative: linePoints(2, 2, 1, 3),
	// 	},
	// }

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(1, 3, 2, 2),
	// 		Negative: linePoints(1, 5, 2, 6),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 2, 3, 3),
	// 		Negative: linePoints(2, 6, 3, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 1, 3),
	// 		Negative: linePoints(2, 4, 1, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 3, 3),
	// 		Negative: linePoints(2, 4, 3, 5),
	// 	},
	// }

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(2, 2, 1, 3),
	// 		Negative: linePoints(2, 6, 3, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 2, 3, 3),
	// 		Negative: linePoints(2, 6, 1, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 1, 3),
	// 		Negative: linePoints(2, 4, 3, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 3, 3),
	// 		Negative: linePoints(2, 4, 1, 5),
	// 	},
	// }

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(2, 2, 1, 3),
	// 		Negative: linePoints(2, 6, 3, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 1, 3),
	// 		Negative: linePoints(2, 4, 3, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 2, 3, 3),
	// 		Negative: linePoints(2, 6, 1, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 3, 3),
	// 		Negative: linePoints(2, 4, 1, 5),
	// 	},
	// }

	// nodes := []Node{
	// 	{
	// 		Positive: linePoints(2, 2, 1, 3),
	// 		Negative: linePoints(2, 6, 3, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 1, 3),
	// 		Negative: linePoints(2, 4, 3, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 4, 3, 3),
	// 		Negative: linePoints(2, 4, 1, 5),
	// 	},
	// 	{
	// 		Positive: linePoints(2, 2, 3, 3),
	// 		Negative: linePoints(2, 6, 1, 5),
	// 	},
	// }

	nodes := []Node{
		{
			Positive: linePoints(2, 2, 3, 3),
			Negative: linePoints(2, 6, 1, 5),
		},
		{
			Positive: linePoints(2, 4, 3, 3),
			Negative: linePoints(2, 4, 1, 5),
		},
		{
			Positive: linePoints(2, 4, 1, 3),
			Negative: linePoints(2, 4, 3, 5),
		},
		{
			Positive: linePoints(2, 2, 1, 3),
			Negative: linePoints(2, 6, 3, 5),
		},
	}

	bs := make([]int, len(nodes))
	calcDigits(digit, bs)

	n := minInt(len(nodes), len(bs))
	for i, node := range nodes[:n] {
		switch bs[i] {
		case 1:
			drawPolyLineGG(c, node.Positive)
		case -1:
			drawPolyLineGG(c, node.Negative)
		}
	}

	c.Stroke()
}
