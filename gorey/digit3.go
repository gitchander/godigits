package gorey

import (
	"image"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
)

type DigitV3 struct {
	Value int
}

var _ Object = DigitV3{}

func (DigitV3) IsObject() {}

func (d DigitV3) Draw(c *cairo.Canvas, r Bounds, level int) {
	//d.draw1(c, r, level)
	// d.draw2(c, r, level)
	d.draw3(c, r, level)
}

func (d DigitV3) draw1(c *cairo.Canvas, r Bounds, level int) {

	const (
		lineWidthRel = 0.5
	)

	var (
		marginRel = geom.MakeFrame1(0.05)
		//marginRel = MakeFrame4(0.02, 0.34, 0.16, 0.08)
	)

	var (
		w = minFloat64(r.Dx(), r.Dy())

		marginAbs = marginRel.MulScalar(w)
	)

	var (
		nx = 8
		ny = 16

		greedWidth = 1.0
	)

	dSize := image.Pt(nx, ny)
	aspectRatio := float64(dSize.Y) / float64(dSize.X)

	var (
		r1 = r.Shrink(marginAbs)
		r2 = subRectByAspectRatio(r1, aspectRatio)

		cellSize = r2.Dx() / float64(dSize.X)
		//cellSize = r2.Dy() / float64(dSize.Y)

		wLineWidth = cellSize * lineWidthRel
		radius     = wLineWidth
	)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, r1)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(2))
		cairoRectangle(c, r2)
		c.Fill()
	}

	dc := NewICanvas(c, r2, dSize.X, dSize.Y)

	dgdr.DrawGreedCairo(c, nx, ny, greedWidth)

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(wLineWidth)

	//----------------------------------------------------------------------

	var (
		// x1 = 2
		// x2 = 6

		x1 = 3
		x2 = 5

		y1 = 1
		y2 = 15
	)

	var (
		val      = d.Value
		negative = false
	)

	//val = ((val + 13) % 27) - 13

	const maxValue = 13
	val = ((val + maxValue) % (2*maxValue + 1)) - maxValue

	if val < 0 {
		val = -val
		negative = true
	}

	vert := 0
	if (val % 2) != 0 {
		vert = 3
	}

	switch vert {

	case 0:
		{
			dc.MoveTo(x1, y1)
			dc.LineTo(x1, y2)

			dc.MoveTo(x2, y1)
			dc.LineTo(x2, y2)
		}

	case 1:
		{
			dc.MoveTo(x1, y1)
			dc.LineTo(x2, y2)

			dc.MoveTo(x2, y1)
			dc.LineTo(x1, y2)
		}

	case 2:
		{
			dc.MoveTo(x1, 1)
			dc.BesierCubic(
				[4]image.Point{
					image.Pt(x1, 1),
					image.Pt(x1, 8),
					image.Pt(x2, 8),
					image.Pt(x2, 15),
				},
			)

			dc.MoveTo(x2, 1)
			dc.BesierCubic(
				[4]image.Point{
					image.Pt(x2, 1),
					image.Pt(x2, 8),
					image.Pt(x1, 8),
					image.Pt(x1, 15),
				},
			)
		}

	case 3:
		{
			{
				dc.MoveTo(x1, 1)
				dc.LineTo(x1, 5)

				dc.BesierCubic(
					[4]image.Point{
						image.Pt(x1, 5),
						image.Pt(x1, 8),
						image.Pt(x2, 8),
						image.Pt(x2, 11),
					},
				)

				dc.MoveTo(x2, 11)
				dc.LineTo(x2, 15)
			}
			//----------------------------------------------
			{
				dc.MoveTo(x2, 1)
				dc.LineTo(x2, 5)

				dc.BesierCubic(
					[4]image.Point{
						image.Pt(x2, 5),
						image.Pt(x2, 8),
						image.Pt(x1, 8),
						image.Pt(x1, 11),
					},
				)

				dc.MoveTo(x1, 11)
				dc.LineTo(x1, 15)
			}
		}
	}
	//----------------------------------------------------------------------

	x1 = x1 - 2
	x2 = x2 + 2

	{
		y1 := 2

		// dy := 1
		// bs := boolsFromBits("1110000000111")

		dy := 2
		bs := parseBools("0000000")

		switch val {
		case 0, 1:
			bs = parseBools("0000000")
		case 2, 3:
			bs = parseBools("0100000")
		case 4, 5:
			bs = parseBools("0100010")
		case 6, 7:
			bs = parseBools("0110010")
		case 8, 9:
			bs = parseBools("0110110")
		case 10, 11:
			bs = parseBools("1110110")
		case 12, 13:
			bs = parseBools("1110111")
		}

		// switch val {
		// case 0:
		// 	bs = boolsFromBits("0000000")
		// case 1:
		// 	bs = boolsFromBits("0000000")
		// case 2:
		// 	bs = boolsFromBits("0100000")
		// case 3:
		// 	bs = boolsFromBits("0100000")
		// case 4:
		// 	bs = boolsFromBits("0100010")
		// case 5:
		// 	bs = boolsFromBits("0100010")
		// case 6:
		// 	bs = boolsFromBits("0110010")
		// case 7:
		// 	bs = boolsFromBits("0110010")
		// case 8:
		// 	bs = boolsFromBits("0110110")
		// case 9:
		// 	bs = boolsFromBits("0110110")
		// case 10:
		// 	bs = boolsFromBits("1110110")
		// case 11:
		// 	bs = boolsFromBits("1110110")
		// case 12:
		// 	bs = boolsFromBits("1110111")
		// case 13:
		// 	bs = boolsFromBits("1110111")
		// }

		for i, b := range bs {
			if b {
				y := y1 + dy*i
				dc.MoveTo(x1, y)
				dc.LineTo(x2, y)
			}
		}
	}
	//----------------------------------------------------------------------

	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.Stroke()

	if negative {

		x1 -= 0
		x2 += 0

		dc.Circle(x1, 8, radius)
		dc.Circle(x2, 8, radius)
		c.Fill()
	}
}

//------------------------------------------------------------------------------

func (d DigitV3) draw2(c *cairo.Canvas, r Bounds, level int) {

	const (
		lineWidthRel = 0.5
	)

	var (
		marginRel = geom.MakeFrame1(0.0)
		//marginRel = MakeFrame4(0.02, 0.34, 0.16, 0.08)
	)

	var (
		w = minFloat64(r.Dx(), r.Dy())

		marginAbs = marginRel.MulScalar(w)
	)

	var (
		nx = 8
		ny = 16

		greedWidth = 1.0
	)

	dSize := image.Pt(nx, ny)
	aspectRatio := float64(dSize.Y) / float64(dSize.X)

	var (
		r1 = r.Shrink(marginAbs)
		r2 = subRectByAspectRatio(r1, aspectRatio)

		cellSize = r2.Dx() / float64(dSize.X)
		//cellSize = r2.Dy() / float64(dSize.Y)

		wLineWidth = cellSize * lineWidthRel
		radius     = wLineWidth
	)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, r1)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(2))
		cairoRectangle(c, r2)
		c.Fill()
	}

	dc := NewICanvas(c, r2, dSize.X, dSize.Y)

	dgdr.DrawGreedCairo(c, nx, ny, greedWidth)

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(wLineWidth)

	//----------------------------------------------------------------------

	var (
		// x1 = 2
		// x2 = 6

		x1 = 3
		x2 = 5

		y1 = 2
		y2 = 14
	)

	var (
		val      = d.Value
		negative = false
	)

	const maxValue = 13
	val = ((val + maxValue) % (2*maxValue + 1)) - maxValue

	if val < 0 {
		val = -val
		negative = true
	}

	vert := 0
	if (val % 2) != 0 {
		vert = 3
	}

	switch vert {

	case 0:
		{
			dc.MoveTo(x1, y1)
			dc.LineTo(x1, y2)

			dc.MoveTo(x2, y1)
			dc.LineTo(x2, y2)
		}

	case 1:
		{
			dc.MoveTo(x1, y1)
			dc.LineTo(x2, y2)

			dc.MoveTo(x2, y1)
			dc.LineTo(x1, y2)
		}

	case 2:
		{
			dc.MoveTo(x1, y1)
			dc.BesierCubic(
				[4]image.Point{
					image.Pt(x1, y1),
					image.Pt(x1, 8),
					image.Pt(x2, 8),
					image.Pt(x2, y2),
				},
			)

			dc.MoveTo(x2, y1)
			dc.BesierCubic(
				[4]image.Point{
					image.Pt(x2, y1),
					image.Pt(x2, 8),
					image.Pt(x1, 8),
					image.Pt(x1, y2),
				},
			)
		}

	case 3:
		{
			{
				dc.MoveTo(x1, y1)
				dc.LineTo(x1, 5)

				dc.BesierCubic(
					[4]image.Point{
						image.Pt(x1, 5),
						image.Pt(x1, 8),
						image.Pt(x2, 8),
						image.Pt(x2, 11),
					},
				)

				dc.MoveTo(x2, 11)
				dc.LineTo(x2, y2)
			}
			//----------------------------------------------
			{
				dc.MoveTo(x2, y1)
				dc.LineTo(x2, 5)

				dc.BesierCubic(
					[4]image.Point{
						image.Pt(x2, 5),
						image.Pt(x2, 8),
						image.Pt(x1, 8),
						image.Pt(x1, 11),
					},
				)

				dc.MoveTo(x1, 11)
				dc.LineTo(x1, y2)
			}
		}
	}
	//----------------------------------------------------------------------

	x1 = x1 - 1
	x2 = x2 + 1

	{
		y1 := 3
		y2 := 8

		dy := 1
		bs := parseBools("000-000")

		// switch val {
		// case 0, 1:
		// 	bs = parseBools("000-000")
		// case 2, 3:
		// 	bs = parseBools("010-000")
		// case 4, 5:
		// 	bs = parseBools("101-000")
		// case 6, 7:
		// 	bs = parseBools("111-000")
		// case 8, 9:
		// 	bs = parseBools("111-010")
		// case 10, 11:
		// 	bs = parseBools("111-101")
		// case 12, 13:
		// 	bs = parseBools("111-111")
		// }

		switch val {
		case 0, 1:
			bs = parseBools("000-000")
		case 2, 3:
			bs = parseBools("010-000")
		case 4, 5:
			bs = parseBools("010-010")
		case 6, 7:
			bs = parseBools("011-010")
		case 8, 9:
			bs = parseBools("011-110")
		case 10, 11:
			bs = parseBools("111-110")
		case 12, 13:
			bs = parseBools("111-111")
		}

		for i, b := range bs {
			if b {
				y := y1
				if i > 2 {
					y = y2
				}

				y += dy * i

				dc.MoveTo(x1, y)
				dc.LineTo(x2, y)
			}
		}
	}
	//----------------------------------------------------------------------

	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.Stroke()

	if negative {

		x1 -= 0
		x2 += 0

		dc.Circle(x1, 8, radius)
		dc.Circle(x2, 8, radius)
		c.Fill()
	}
}

//------------------------------------------------------------------------------

func (d DigitV3) draw3(c *cairo.Canvas, r Bounds, level int) {

	const (
		lineWidthRel = 0.5
		//lineWidthRel = 0.45
	)

	var (
		marginRel = geom.MakeFrame1(0.0)
		//marginRel = MakeFrame4(0.02, 0.34, 0.16, 0.08)
	)

	var (
		w = minFloat64(r.Dx(), r.Dy())

		marginAbs = marginRel.MulScalar(w)
	)

	var (
		nx = 8
		ny = 16

		greedWidth = 1.0
	)

	dSize := image.Pt(nx, ny)
	aspectRatio := float64(dSize.Y) / float64(dSize.X)

	var (
		r1 = r.Shrink(marginAbs)
		r2 = subRectByAspectRatio(r1, aspectRatio)

		cellSize = r2.Dx() / float64(dSize.X)
		//cellSize = r2.Dy() / float64(dSize.Y)

		wLineWidth = cellSize * lineWidthRel
		radius     = wLineWidth
	)

	if DrawLevelArea {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, r1)
		c.Fill()

		cairoSetSourceColor(c, levelToColor(2))
		cairoRectangle(c, r2)
		c.Fill()
	}

	dc := NewICanvas(c, r2, dSize.X, dSize.Y)

	dgdr.DrawGreedCairo(c, nx, ny, greedWidth)

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(wLineWidth)

	//----------------------------------------------------------------------

	var (
		// x1 = 2
		// x2 = 6

		x1 = 3
		x2 = 5

		y1 = 2
		y2 = 14
	)

	var (
		val      = d.Value
		negative = false
	)

	const maxValue = 15
	val = ((val + maxValue) % (2*maxValue + 1)) - maxValue

	if val < 0 {
		val = -val
		negative = true
	}

	vert := 0
	if (val % 2) != 0 {
		vert = 3
	}

	switch vert {

	case 0:
		{
			dc.MoveTo(x1, y1)
			dc.LineTo(x1, y2)

			dc.MoveTo(x2, y1)
			dc.LineTo(x2, y2)
		}

	case 1:
		{
			dc.MoveTo(x1, y1)
			dc.LineTo(x2, y2)

			dc.MoveTo(x2, y1)
			dc.LineTo(x1, y2)
		}

	case 2:
		{
			dc.MoveTo(x1, y1)
			dc.BesierCubic(
				[4]image.Point{
					image.Pt(x1, y1),
					image.Pt(x1, 8),
					image.Pt(x2, 8),
					image.Pt(x2, y2),
				},
			)

			dc.MoveTo(x2, y1)
			dc.BesierCubic(
				[4]image.Point{
					image.Pt(x2, y1),
					image.Pt(x2, 8),
					image.Pt(x1, 8),
					image.Pt(x1, y2),
				},
			)
		}

	case 3:
		{
			{
				dc.MoveTo(x1, y1)
				dc.LineTo(x1, 5)

				dc.BesierCubic(
					[4]image.Point{
						image.Pt(x1, 5),
						image.Pt(x1, 8),
						image.Pt(x2, 8),
						image.Pt(x2, 11),
					},
				)

				dc.MoveTo(x2, 11)
				dc.LineTo(x2, y2)
			}
			//----------------------------------------------
			{
				dc.MoveTo(x2, y1)
				dc.LineTo(x2, 5)

				dc.BesierCubic(
					[4]image.Point{
						image.Pt(x2, 5),
						image.Pt(x2, 8),
						image.Pt(x1, 8),
						image.Pt(x1, 11),
					},
				)

				dc.MoveTo(x1, 11)
				dc.LineTo(x1, y2)
			}
		}
	}
	//----------------------------------------------------------------------

	x1 = x1 - 1
	x2 = x2 + 1

	{
		y1 := 3
		//y2 := 8

		dy := 1
		bs := parseBools("000-000")

		switch val {
		case 0, 1:
			bs = parseBools("000-000")
		case 2, 3:
			bs = parseBools("00000-1-00000")
		case 4, 5:
			bs = parseBools("01000-0-00010")
		case 6, 7:
			bs = parseBools("01000-1-00010")

		//---------------------------------------------
		case 8, 9:
			bs = parseBools("01100-0-00110")
		case 10, 11:
			bs = parseBools("01100-1-00110")
		//---------------------------------------------
		// case 8, 9:
		// 	bs = parseBools("10100-0-00101")
		// case 10, 11:
		// 	bs = parseBools("10100-1-00101")
		//---------------------------------------------

		case 12, 13:
			bs = parseBools("11100-0-00111")
		case 14, 15:
			bs = parseBools("11100-1-00111")
		}

		for i, b := range bs {
			if b {
				y := y1
				y += dy * i

				dc.MoveTo(x1, y)
				dc.LineTo(x2, y)
			}
		}
	}
	//----------------------------------------------------------------------

	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.Stroke()

	if negative {

		//----------------------------------------------------------------------
		// dc.Circle(x1, 8, radius)
		// dc.Circle(x2, 8, radius)
		// c.Fill()
		//----------------------------------------------------------------------
		dc.Circle(4, 6, radius)
		dc.Circle(4, 10, radius)
		c.Fill()
		//----------------------------------------------------------------------
		// dc.Circle(4, 8, radius)
		// cairoSetSourceColor(c, levelToColor(2))
		// c.FillPreserve()
		// c.SetSourceColor(Black)
		// c.Stroke()
		//----------------------------------------------------------------------
	}
}

func parseBools(s string) []bool {
	var (
		rs = []rune(s)
		bs = make([]bool, 0, len(rs))
	)
	for _, r := range rs {
		switch r {
		case '0':
			bs = append(bs, false)
		case '1':
			bs = append(bs, true)
		case '-', '_': // skip
		default:
			panic("invalid bits")
		}
	}
	return bs
}
