package base27

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/utils/digits"
)

type NodeDrawers struct {
	Positive dgdr.DrawerGG
	Negative dgdr.DrawerGG
}

type Digit1 struct{}

var _ dgdr.DigitDrawer = Digit1{}

func (Digit1) Width(height float64) (width float64) {
	return height / 2
}

func (d Digit1) DrawDigit(c *gg.Context, x, y float64, height float64, digit int) {

	var (
		nx = 8
		ny = 16

		width = d.Width(height)

		dx = width / float64(nx)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.5 * dx

		circleRadius = 0.5
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var (
		x0 float64 = 5
		y0 float64 = 4
	)

	// Draw anchor
	{
		if true {
			c.MoveTo(x0-4, y0)
			c.QuadraticTo(x0, y0-3, x0, y0)
		} else {
			c.MoveTo(x0, y0)
		}
		c.LineTo(x0, y0+10)

		c.Stroke()
	}

	// Draw digit
	{
		type node struct {
			positive geom.Point2f
			negative geom.Point2f
		}

		var (
			xnd float64 = 1

			yn0 float64 = y0 + 2
			ynd float64 = 3
		)

		nodes := []node{
			{
				negative: geom.Pt2f(x0-xnd, yn0+0*ynd),
				positive: geom.Pt2f(x0+xnd, yn0+0*ynd),
			},
			{
				negative: geom.Pt2f(x0-xnd, yn0+1*ynd),
				positive: geom.Pt2f(x0+xnd, yn0+1*ynd),
			},
			{
				negative: geom.Pt2f(x0-xnd, yn0+2*ynd),
				positive: geom.Pt2f(x0+xnd, yn0+2*ynd),
			},
		}

		rd := digits.MustNewRestDigiter(-1, +1)

		bs := make([]int, len(nodes))
		digits.CalcDigits(rd, digit, bs)

		var (
			//xd float64 = 0.5
			//xd float64 = 1.25
			//xd float64 = 1.5
			//xd float64 = 1.75
			xd float64 = 2

			yd1, ydd float64 = y0 + 3.5, 3
		)

		if digit < 0 {
			digit = -digit
		}

		k := len(bs)
		if false {
			for (k > 0) && (bs[k-1] == 0) {
				k--
			}
		}
		for i := 0; i < k-1; i++ {
			y := yd1 + float64(i)*ydd
			c.MoveTo(x0-xd, y)
			c.LineTo(x0+xd, y)
		}
		c.Stroke()

		for i, t := range bs {
			switch t {
			case -1:
				v := nodes[i].negative
				c.DrawCircle(v.X, v.Y, circleRadius)
			case 1:
				v := nodes[i].positive
				c.DrawCircle(v.X, v.Y, circleRadius)
			}
			c.Fill()
		}
	}
}

//------------------------------------------------------------------------------

type Digit2 struct{}

var _ dgdr.DigitDrawer = Digit2{}

func (Digit2) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit2) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 6
		ny = 12

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.75 * w
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineCap(gg.LineCapRound)
	c.SetLineJoin(gg.LineJoinRound)
	c.SetLineWidth(lineWidth)

	var (
		//------------------------------
		x0 float64 = 1
		y0 float64 = 2

		dx float64 = 2
		dy float64 = 2
		//------------------------------
		// x0 float64 = 1.5
		// y0 float64 = 3

		// dx float64 = 1.5
		// dy float64 = 1.5
		//------------------------------
		// x0 float64 = 0
		// y0 float64 = 0

		// dx float64 = 3
		// dy float64 = 3
		//------------------------------
	)

	var (
		x1 = x0 + 1*dx
		x2 = x0 + 2*dx

		y1 = y0 + 1*dy
		y2 = y0 + 2*dy
		y3 = y0 + 3*dy
		y4 = y0 + 4*dy
	)

	c.MoveTo(x1, y0)
	c.LineTo(x1, y4)

	nodes := []NodeDrawers{
		{
			Positive: dgdr.MakeLine(x1, y0, x2, y1),
			Negative: dgdr.MakeLine(x1, y0, x0, y1),
		},
		{
			Positive: dgdr.MakeLine(x1, y2, x2, y1),
			Negative: dgdr.MakeLine(x1, y2, x0, y1),
		},
		{
			Positive: dgdr.MakeLine(x1, y2, x2, y3),
			Negative: dgdr.MakeLine(x1, y2, x0, y3),
		},
		{
			Positive: dgdr.MakeLine(x1, y4, x2, y3),
			Negative: dgdr.MakeLine(x1, y4, x0, y3),
		},
	}

	trits := calcTritsBal27(digit)

	n := minInt(len(nodes), len(trits))
	for i, node := range nodes[:n] {
		switch trits[i] {
		case 1:
			node.Positive.DrawGG(c)
		case -1:
			node.Negative.DrawGG(c)
		}
	}

	c.Stroke()
}

//------------------------------------------------------------------------------

type Digit6 struct{}

var _ dgdr.DigitDrawer = Digit6{}

func (Digit6) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit6) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 6
		ny = 12

		dx = digitHeight / float64(ny)
		dy = dx

		greedWidth = 0.02 * dx
		lineWidth  = 0.75 * dx
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(dx, dy)

	//dgdr.DrawGreedEnable = true
	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var (
		// x0 float64 = 1
		// y0 float64 = 2

		// xd float64 = 1.5
		// yd float64 = 1.5

		x0 float64 = 1
		y0 float64 = 2

		xd float64 = 2
		yd float64 = 2
	)

	var (
		x1 = x0 + 1*xd
		x2 = x0 + 2*xd

		y1 = y0 + 1*yd
		y2 = y0 + 2*yd
		y3 = y0 + 3*yd
		y4 = y0 + 4*yd
	)

	if true {
		c.MoveTo(x1, y0)
		c.LineTo(x1, y4)

		c.MoveTo(x0, y0)
		c.LineTo(x2, y0)

		c.MoveTo(x0, y4)
		c.LineTo(x2, y4)
	}

	nodes := []NodeDrawers{
		{
			Positive: dgdr.MakeLine(x2, y0, x2, y1),
			Negative: dgdr.MakeLine(x0, y0, x0, y1),
		},
		{
			Positive: dgdr.MakeLine(x1, y2, x2, y1),
			Negative: dgdr.MakeLine(x1, y2, x0, y1),
		},
		{
			Positive: dgdr.MakeLine(x1, y2, x2, y3),
			Negative: dgdr.MakeLine(x1, y2, x0, y3),
		},
		{
			Positive: dgdr.MakeLine(x2, y4, x2, y3),
			Negative: dgdr.MakeLine(x0, y4, x0, y3),
		},
	}

	trits := calcTritsBal27(digit)

	n := minInt(len(nodes), len(trits))
	for i, node := range nodes[:n] {
		switch trits[i] {
		case 1:
			node.Positive.DrawGG(c)
		case -1:
			node.Negative.DrawGG(c)
		}
	}

	c.Stroke()
}
