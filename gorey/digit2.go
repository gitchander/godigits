package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/hexagone"
	"github.com/gitchander/godigits/utils/gobits"
)

type DigitV2 struct {
	Value int
}

var _ Object = DigitV2{}

func (DigitV2) IsObject() {}

func (d DigitV2) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	const (
		//lineWidthRel = 5
		lineWidthRel = 7
	)

	var (
		marginRel = geom.MakeFrame1(3)
	)

	var (
		vm = vmin(r)

		marginAbs = marginRel.MulScalar(vm)
		lineWidth = lineWidthRel * vm
	)

	var (
		mr     = r.Shrink(marginAbs)
		mrw    = minFloat64(mr.Dx(), mr.Dy())
		center = mr.Center()
		radius = (mrw - lineWidth) / 2
	)

	var (
		drawAreas = DrawLevelArea
		//drawAreas = false
	)
	if drawAreas {
		cairoSetSourceColor(c, levelToColor(0))
		cairoRectangle(c, r)
		c.Fill()
		//---------------------------------------------
		cairoSetSourceColor(c, levelToColor(1))
		cairoRectangle(c, mr)
		c.Fill()
		//---------------------------------------------
		cairoSetSourceColor(c, levelToColor(2))
		cairoCircle(c, center, (radius + lineWidth/2))
		c.Fill()
	}

	ps := hexagone.HexagoneVertexes(center, radius)

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(lineWidth)
	c.SetLineCap(cairo.LINE_CAP_ROUND)

	var drawLine func(a, b geom.Point2f)

	switch 2 {
	case 0:
		drawLine = func(a, b geom.Point2f) {
			c.MoveTo(a.X, a.Y)
			c.LineTo(b.X, b.Y)
		}

	case 1:
		drawLine = func(a, b geom.Point2f) {
			const (
				//koef = 0.14
				koef = 0.2
			)
			var (
				a1 = geom.PtLerp(a, b, koef)
				b1 = geom.PtLerp(b, a, koef)
			)
			c.MoveTo(a1.X, a1.Y)
			c.LineTo(b1.X, b1.Y)
		}

	case 2:
		drawLine = func(a, b geom.Point2f) {
			a, b = hexagone.ShrinkPoints(a, b, 0.02)
			hexagone.CairoDrawSegment(c, a, b, lineWidth, tau/6)
		}
	}

	//v := uint16(d.Value)
	v := digitV2Segments(d.Value % 40)
	//v := digitV2Segments_v2(d.Value)

	// Rays:
	for i, p := range ps {
		if gobits.GetBit(v, i) == 1 {
			drawLine(center, p)
		}
	}

	// Edges:
	for i, p := range ps {
		if gobits.GetBit(v, 6+i) == 1 {
			next := ps[(i+1)%len(ps)]
			drawLine(p, next)
		}
	}

	c.Stroke()
}

func digitV2Segments(v int) uint16 {

	useLower := true

	switch v {
	case 0:
		//return 0b_111111_000000
		return 0b_111111_010010
	case 1:
		return 0b_100000_001001
	case 2:
		return 0b_101101_010010
	case 3:
		return 0b_101101_000110
	case 4:
		return 0b_010000_011001
	case 5:
		return 0b_101101_100100
	case 6:
		return 0b_111101_010100
	case 7:
		return 0b_100001_001010
	case 8:
		return 0b_101101_110110
	case 9:
		return 0b_101111_100010
	case 10: // A
		return 0b_110011_100010
	case 11: // B
		if useLower {
			return 0b_000100_001101 // lower
		} else {
			return 0b_000101_001111
		}
	case 12: // C
		return 0b_111101_000000
	case 13: // D
		if useLower {
			return 0b_001000_011001 // lower
		} else {
			return 0b_000111_001001
		}
	case 14: // E
		if useLower {
			return 0b_011000_110000
		} else {
			return 0b_101000_110000
		}
	case 15: // F
		return 0b_000001_001101
	case 16: // G
		if useLower {
			return 0b_101000_101001 // lower
		} else {
			return 0b_111101_000100
		}
	case 17: // H
		return 0b_000010_001011
	case 18: // I
		return 0b_000000_001001
	case 19: // J
		return 0b_001000_001001
	case 20: // K
		return 0b_000000_001111
	case 21: // L
		return 0b_000100_001001
	case 22: // M
		return 0b_010010_100010
	case 23: // N
		return 0b_010010_010010
	case 24: // O
		return 0b_111111_000000
	case 25: // P
		return 0b_000001_001011
	case 26: // Q
		return 0b_100000_101001
	case 27: // R
		return 0b_000001_001111
	case 28: // S
		return 0b_000100_000110
	case 29: // T
		return 0b_100001_001001
	case 30: // U
		return 0b_011110_000000

	default:
		return 0xfff
	}
}

func digitV2Segments_v2(v int) uint16 {
	switch v {
	case 0:
		return 0b_000000_001001
	case 1:
		return 0b_000001_001001
	case 2:
		return 0b_000001_001011
	case 3:
		return 0b_000001_101011
	case 4:
		return 0b_000000_001101
	case 5:
		return 0b_000001_001101
	case 6:
		return 0b_000001_001111
	case 7:
		return 0b_000001_101111
	case 8:
		return 0b_000100_001101
	case 9:
		return 0b_000101_001101
	case 10:
		return 0b_000101_001111
	case 11:
		return 0b_000101_101111
	case 12:
		return 0b_001100_001101
	case 13:
		return 0b_001101_001101
	case 14:
		return 0b_001101_001111
	case 15:
		return 0b_001101_101111
	default:
		return 0x0
	}
}
