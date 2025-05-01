package bal81

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
)

type (
	Digit8 = Digit8Size4x4
	//Digit8 = Digit8Size6x6
)

//------------------------------------------------------------------------------

type Digit8Size4x4 struct{}

var _ dgdr.DigitDrawer = Digit8Size4x4{}

func (Digit8Size4x4) Width(height float64) (width float64) {
	width = height
	return
}

func (d Digit8Size4x4) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

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

	{
		c.MoveTo(2, 1)
		c.LineTo(2, 3)

		c.MoveTo(1, 2)
		c.LineTo(3, 2)
		c.Stroke()
	}

	c.DrawCircle(2.5, 1.5, 0.2)
	c.Fill()

	//--------------------------------------------------------------------------

	nodes := []NodeDrawers{
		{
			Positive: dgdr.MakeLine(3, 2, 3, 1),
			Negative: dgdr.MakeLine(1, 2, 1, 3),
		},
		{
			Positive: dgdr.MakeLine(2, 1, 3, 1),
			Negative: dgdr.MakeLine(2, 3, 1, 3),
		},
		{
			Positive: dgdr.MakeLine(2, 1, 1, 1),
			Negative: dgdr.MakeLine(2, 3, 3, 3),
		},
		{
			Positive: dgdr.MakeLine(1, 2, 1, 1),
			Negative: dgdr.MakeLine(3, 2, 3, 3),
		},
	}

	//--------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

type Digit8Size6x6 struct{}

var _ dgdr.DigitDrawer = Digit8Size6x6{}

func (Digit8Size6x6) Width(height float64) (width float64) {
	width = height
	return
}

func (d Digit8Size6x6) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 6
		ny = 6

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.33 * w
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineCap(gg.LineCapRound)
	c.SetLineJoin(gg.LineJoinRound)
	c.SetLineWidth(lineWidth)

	{
		c.MoveTo(3, 1)
		c.LineTo(3, 5)

		c.MoveTo(1, 3)
		c.LineTo(5, 3)
	}

	// {
	// 	c.MoveTo(3, 2)
	// 	c.LineTo(3, 4)

	// 	c.MoveTo(2, 3)
	// 	c.LineTo(4, 3)
	// }

	c.Stroke()

	c.DrawCircle(4, 2, 0.25)
	c.Fill()

	//--------------------------------------------------------------------------

	nodes := []NodeDrawers{
		{
			Positive: dgdr.MakeLine(3, 1, 5, 1),
			Negative: dgdr.MakeLine(3, 5, 1, 5),
		},
		{
			Positive: dgdr.MakeLine(5, 3, 5, 1),
			Negative: dgdr.MakeLine(1, 3, 1, 5),
		},
		{
			Positive: dgdr.MakeLine(5, 3, 5, 5),
			Negative: dgdr.MakeLine(1, 3, 1, 1),
		},
		{
			Positive: dgdr.MakeLine(3, 5, 5, 5),
			Negative: dgdr.MakeLine(3, 1, 1, 1),
		},
	}

	// nodes := []NodeDrawers{
	// 	{
	// 		Positive: MakeLine(3, 2, 4, 2),
	// 		Negative: MakeLine(3, 4, 2, 4),
	// 	},
	// 	{
	// 		Positive: MakeLine(4, 3, 4, 2),
	// 		Negative: MakeLine(2, 3, 2, 4),
	// 	},
	// 	{
	// 		Positive: MakeLine(4, 3, 4, 4),
	// 		Negative: MakeLine(2, 3, 2, 2),
	// 	},
	// 	{
	// 		Positive: MakeLine(3, 4, 4, 4),
	// 		Negative: MakeLine(3, 2, 2, 2),
	// 	},
	// }

	// nodes := []NodeDrawers{
	// 	{
	// 		Positive: MakeLine(3, 1, 4, 2),
	// 		Negative: MakeLine(3, 5, 2, 4),
	// 	},
	// 	{
	// 		Positive: MakeLine(5, 3, 4, 2),
	// 		Negative: MakeLine(1, 3, 2, 4),
	// 	},
	// 	{
	// 		Positive: MakeLine(5, 3, 4, 4),
	// 		Negative: MakeLine(1, 3, 2, 2),
	// 	},
	// 	{
	// 		Positive: MakeLine(3, 5, 4, 4),
	// 		Negative: MakeLine(3, 1, 2, 2),
	// 	},
	// }

	//--------------------------------------------------------------------------

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
