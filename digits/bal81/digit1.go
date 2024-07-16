package bal81

import (
	"github.com/fogleman/gg"
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/dgdr"
)

type Digit1 struct{}

var _ dgdr.DigitDrawer = Digit1{}

func (Digit1) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit1) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

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

	nodes := []nodeDrawers{
		{
			positive: dgdr.MakeLine(2, 2, 3, 3),
			negative: dgdr.MakeLine(2, 2, 1, 3),
		},
		{
			positive: dgdr.MakeLine(2, 4, 3, 3),
			negative: dgdr.MakeLine(2, 4, 1, 3),
		},
		{
			positive: dgdr.MakeLine(2, 4, 3, 5),
			negative: dgdr.MakeLine(2, 4, 1, 5),
		},
		{
			positive: dgdr.MakeLine(2, 6, 3, 5),
			negative: dgdr.MakeLine(2, 6, 1, 5),
		},
	}

	bs := CalcDigitsBal3(digit)

	n := minInt(len(nodes), len(bs))
	for i, node := range nodes[:n] {
		switch bs[i] {
		case 1:
			node.positive.DrawGG(c)
		case -1:
			node.negative.DrawGG(c)
		}
	}

	c.Stroke()
}

//------------------------------------------------------------------------------

type CairoDigit1 struct{}

var _ dgdr.CairoDigitDrawer = CairoDigit1{}

func (CairoDigit1) Ratio() float64 {
	return 2
}

func (d CairoDigit1) DrawDigit(c *cairo.Canvas, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 4
		ny = 8

		w = digitHeight / float64(ny)

		greedWidth = 0.02
		lineWidth  = 0.4
	)

	c.Save()
	defer c.Restore()

	c.Translate(x, y)
	c.Scale(w, w)

	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.SetLineJoin(cairo.LINE_JOIN_ROUND)

	dgdr.DrawGreedCairo(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	c.MoveTo(2, 2)
	c.LineTo(2, 6)

	nodes := []nodeDrawers{
		{
			positive: dgdr.MakeLine(2, 2, 3, 3),
			negative: dgdr.MakeLine(2, 2, 1, 3),
		},
		{
			positive: dgdr.MakeLine(2, 4, 3, 3),
			negative: dgdr.MakeLine(2, 4, 1, 3),
		},
		{
			positive: dgdr.MakeLine(2, 4, 3, 5),
			negative: dgdr.MakeLine(2, 4, 1, 5),
		},
		{
			positive: dgdr.MakeLine(2, 6, 3, 5),
			negative: dgdr.MakeLine(2, 6, 1, 5),
		},
	}

	bs := CalcDigitsBal3(digit)

	n := minInt(len(nodes), len(bs))
	for i, node := range nodes[:n] {
		switch bs[i] {
		case 1:
			node.positive.DrawCairo(c)
		case -1:
			node.negative.DrawCairo(c)
		}
	}

	c.Stroke()
}
