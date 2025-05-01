package bal81

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
)

type Digit3 struct{}

var _ dgdr.DigitDrawer = Digit3{}

func (Digit3) Width(height float64) (width float64) {
	width = height
	return
}

func (d Digit3) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 4
		ny = 4

		dx = digitHeight / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.2 * dx
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	//center := Pt2f(2, 2)

	var (
		x1, x2, x3 float64 = 1, 2, 3
		y1, y2, y3 float64 = 1, 2, 3

		// x1, x2, x3 float64 = 0.5, 2, 3.5
		// y1, y2, y3 float64 = 0.5, 2, 3.5
	)

	c.MoveTo(x1, y1)
	c.LineTo(x3, y3)

	c.MoveTo(x1, y3)
	c.LineTo(x3, y1)

	c.Stroke()

	//--------------------------------------------------------------------------

	// nodes := []Node{
	// 	{
	// 		Positive: []Point2f{Pt2f(x3, y1), Pt2f(x3, y2)},
	// 		Negative: []Point2f{Pt2f(x3, y1), Pt2f(x2, y1)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(x3, y3), Pt2f(x2, y3)},
	// 		Negative: []Point2f{Pt2f(x3, y3), Pt2f(x3, y2)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(x1, y3), Pt2f(x1, y2)},
	// 		Negative: []Point2f{Pt2f(x1, y3), Pt2f(x2, y3)},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(x1, y1), Pt2f(x2, y1)},
	// 		Negative: []Point2f{Pt2f(x1, y1), Pt2f(x1, y2)},
	// 	},
	// }

	// nodes := []NodeDrawers{
	// 	{
	// 		Positive: MakeLine(x2, y1, x3, y1),
	// 		Negative: MakeLine(x2, y1, x1, y1),
	// 	},
	// 	{
	// 		Positive: MakeLine(x3, y2, x3, y3),
	// 		Negative: MakeLine(x3, y2, x3, y1),
	// 	},
	// 	{
	// 		Positive: MakeLine(x2, y3, x1, y3),
	// 		Negative: MakeLine(x2, y3, x3, y3),
	// 	},
	// 	{
	// 		Positive: MakeLine(x1, y2, x1, y1),
	// 		Negative: MakeLine(x1, y2, x1, y3),
	// 	},
	// }

	// nodes := []Node{
	// 	{
	// 		Positive: []Point2f{Pt2f(x2, y1), Pt2f(middle(x2, x3), middle(y1, y2))},
	// 		Negative: []Point2f{Pt2f(x2, y1), Pt2f(middle(x1, x2), middle(y1, y2))},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(x3, y2), Pt2f(middle(x2, x3), middle(y2, y3))},
	// 		Negative: []Point2f{Pt2f(x3, y2), Pt2f(middle(x2, x3), middle(y1, y2))},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(x2, y3), Pt2f(middle(x1, x2), middle(y2, y3))},
	// 		Negative: []Point2f{Pt2f(x2, y3), Pt2f(middle(x2, x3), middle(y2, y3))},
	// 	},
	// 	{
	// 		Positive: []Point2f{Pt2f(x1, y2), Pt2f(middle(x1, x2), middle(y1, y2))},
	// 		Negative: []Point2f{Pt2f(x1, y2), Pt2f(middle(x1, x2), middle(y2, y3))},
	// 	},
	// }

	nodes := []Node{
		{
			Positive: []geom.Point2f{geom.Pt2f(x2, y1), geom.Pt2f(x3, y1)},
			Negative: []geom.Point2f{geom.Pt2f(x2, y3), geom.Pt2f(x1, y3)},
		},
		{
			Positive: []geom.Point2f{geom.Pt2f(x3, y2), geom.Pt2f(x3, y1)},
			Negative: []geom.Point2f{geom.Pt2f(x1, y2), geom.Pt2f(x1, y3)},
		},
		{
			Positive: []geom.Point2f{geom.Pt2f(x3, y2), geom.Pt2f(x3, y3)},
			Negative: []geom.Point2f{geom.Pt2f(x1, y2), geom.Pt2f(x1, y1)},
		},
		{
			Positive: []geom.Point2f{geom.Pt2f(x2, y3), geom.Pt2f(x3, y3)},
			Negative: []geom.Point2f{geom.Pt2f(x2, y1), geom.Pt2f(x1, y1)},
		},
	}

	//--------------------------------------------------------------------------

	bs := CalcDigitsBal3(digit)

	n := minInt(len(nodes), len(bs))
	for i := 0; i < n; i++ {

		var (
			node = nodes[i]
			b    = bs[i]
		)

		switch b {
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

		// switch b {
		// case 1:
		// 	node.Positive.DrawGG(c)
		// case -1:
		// 	node.Negative.DrawGG(c)
		// }
	}

	c.Stroke()
}
