package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/hexagone"
	"github.com/gitchander/godigits/utils/gobits"
)

// ------------------------------------------------------------------------------
// Seven-segment display
// https://en.wikipedia.org/wiki/Seven-segment_display
// ------------------------------------------------------------------------------

//     +-----------+
//   x       a       x
// +   +-----------+   +
// |   |           |   |
// |   |           |   |
// | f |           | b |
// |   |           |   |
// |   |           |   |
// +   +-----------+   +
//   x       g       x
// +   +-----------+   +
// |   |           |   |
// |   |           |   |
// | e |           | c |
// |   |           |   |
// |   |           |   |
// +   +-----------+   +
//   x       d       x
//     +-----------+

// ------------------------------------------------------------------------------
func charToIndex(b byte) int {
	if ('a' <= b) && (b <= 'g') {
		return int(b - 'a')
	}
	err := fmt.Errorf("invalid segment %c", b)
	panic(err)
}

func runeToIndex(r byte) int {
	if ('a' <= r) && (r <= 'g') {
		return int(r - 'a')
	}
	err := fmt.Errorf("invalid segment %c", r)
	panic(err)
}

func segmentIsOn(x uint8, b byte) bool {
	return gobits.Uint8GetBit(x, charToIndex(b)) == 1
}

func flipSegmentsVertical(x uint8) uint8 {

	var (
		f = gobits.Uint8GetBit(x, charToIndex('f'))
		b = gobits.Uint8GetBit(x, charToIndex('b'))

		e = gobits.Uint8GetBit(x, charToIndex('e'))
		c = gobits.Uint8GetBit(x, charToIndex('c'))
	)

	// f <-> b
	x = gobits.Uint8SetBit(x, charToIndex('f'), b)
	x = gobits.Uint8SetBit(x, charToIndex('b'), f)

	// e <-> c
	x = gobits.Uint8SetBit(x, charToIndex('e'), c)
	x = gobits.Uint8SetBit(x, charToIndex('c'), e)

	return x
}

// ------------------------------------------------------------------------------
func printableSSD(x uint8, prefix string) string {
	const (
		filler = ' '

		segmentCharOff = '.'
		segmentCharOn  = 'X'
	)
	const (
		xn = 4
		yn = 3
	)
	bss := make([][]byte, (1 + yn + 1 + yn + 1))
	for y := range bss {
		bs := make([]byte, (1 + xn + 1))
		for x := range bs {
			bs[x] = filler
		}
		bss[y] = bs
	}

	var c byte

	makeFillChar := func(segmentOn bool) byte {
		if segmentOn {
			return segmentCharOn
		}
		return segmentCharOff
	}

	c = makeFillChar(segmentIsOn(x, 'a'))
	for x := 0; x < xn; x++ {
		bss[0][1+x] = c
	}

	c = makeFillChar(segmentIsOn(x, 'b'))
	for y := 0; y < yn; y++ {
		bss[1+y][1+xn] = c
	}

	c = makeFillChar(segmentIsOn(x, 'c'))
	for y := 0; y < yn; y++ {
		bss[yn+2+y][1+xn] = c
	}

	c = makeFillChar(segmentIsOn(x, 'd'))
	for x := 0; x < xn; x++ {
		bss[2+2*yn][1+x] = c
	}

	c = makeFillChar(segmentIsOn(x, 'e'))
	for y := 0; y < yn; y++ {
		bss[yn+2+y][0] = c
	}

	c = makeFillChar(segmentIsOn(x, 'f'))
	for y := 0; y < yn; y++ {
		bss[1+y][0] = c
	}

	c = makeFillChar(segmentIsOn(x, 'g'))
	for x := 0; x < xn; x++ {
		bss[1+yn][1+x] = c
	}

	var b strings.Builder
	for _, bs := range bss {
		b.WriteString(prefix)
		b.WriteString(string(bs))
		b.WriteByte('\n')
	}
	return b.String()
}

func ParseSSD(s string) (uint8, error) {
	var x uint8
	for _, r := range s {
		if ('a' <= r) && (r <= 'g') {
			x |= 1 << int(r-'a')
		} else {
			if r == '_' {
				continue
			}
			return 0, fmt.Errorf("ssd value has invalid rune %U", r)
		}
	}
	return x, nil
}

func MustParseSSD(s string) uint8 {
	x, err := ParseSSD(s)
	if err != nil {
		panic(err)
	}
	return x
}

func makeDigitsMap() map[int]uint8 {
	xs := []uint8{
		0:  MustParseSSD("bcdefg"),
		1:  MustParseSSD("ef_a"),
		2:  MustParseSSD("ef_g"),
		3:  MustParseSSD("ef_d"),
		4:  MustParseSSD("ef_ag"),
		5:  MustParseSSD("ef_ad"),
		6:  MustParseSSD("ef_gd"),
		7:  MustParseSSD("ef_ab"),
		8:  MustParseSSD("ef_gb"),
		9:  MustParseSSD("ef_gc"),
		10: MustParseSSD("ef_dc"),
		11: MustParseSSD("ef_agd"),
		12: MustParseSSD("ef_abg"),
		13: MustParseSSD("ef_abd"),
		14: MustParseSSD("ef_gbd"),
		15: MustParseSSD("ef_agc"),
		16: MustParseSSD("ef_acd"),
		17: MustParseSSD("ef_gcd"),
		18: MustParseSSD("ef_abg_d"),
		19: MustParseSSD("ef_a_gcd"),
	}
	m := make(map[int]uint8)

	for i, x := range xs {
		if i == 0 {
			m[i] = x
		} else {
			m[i] = x
			m[-i] = flipSegmentsVertical(x)
		}
	}

	return m
}

var digitsMap = makeDigitsMap()

func makeImagesSSD() {

	var size float64 = 512 // height

	var (
		nx = 10
		ny = 18
	)

	var (
		//lk = 0.25
		lk = 0.0

		w = size / (float64(ny) + lk)

		lineWidth = w * lk
	)

	var (
		width = ceilInt(w*float64(nx) + lineWidth)

		//height = ceilInt(w*float64(ny) + lineWidth)
		height = ceilInt(size) // w*float64(ny) + lineWidth)
	)

	fmt.Println(width, height)

	c := gg.NewContext(width, height)

	c.Translate(lineWidth/2, lineWidth/2)
	c.Scale(w, w)

	c.SetColor(color.White)
	//c.SetColor(color.Black)
	c.Clear()

	c.SetLineWidth(lineWidth)

	c.SetLineCap(gg.LineCapButt)
	//c.SetLineCap(gg.LineCapRound)
	//c.SetLineCap(gg.LineCapSquare)

	//c.SetLineJoin(gg.LineJoinRound)
	c.SetLineJoin(gg.LineJoinBevel)

	m := map[string][]Point2f{
		"a": []Point2f{Pt(1, 1), Pt(2, 0), Pt(8, 0), Pt(9, 1), Pt(8, 2), Pt(2, 2)},
		"b": []Point2f{Pt(9, 1), Pt(10, 2), Pt(10, 8), Pt(9, 9), Pt(8, 8), Pt(8, 2)},
		"c": []Point2f{Pt(9, 9), Pt(10, 10), Pt(10, 16), Pt(9, 17), Pt(8, 16), Pt(8, 10)},
		"d": []Point2f{Pt(9, 17), Pt(8, 18), Pt(2, 18), Pt(1, 17), Pt(2, 16), Pt(8, 16)},
		"e": []Point2f{Pt(1, 17), Pt(0, 16), Pt(0, 10), Pt(1, 9), Pt(2, 10), Pt(2, 16)},
		"f": []Point2f{Pt(1, 9), Pt(0, 8), Pt(0, 2), Pt(1, 1), Pt(2, 2), Pt(2, 8)},
		"g": []Point2f{Pt(1, 9), Pt(2, 8), Pt(8, 8), Pt(9, 9), Pt(8, 10), Pt(2, 10)},
	}

	drawSegment(c, ShrinkPolygon(m["a"]), true)
	drawSegment(c, ShrinkPolygon(m["b"]), false)
	drawSegment(c, ShrinkPolygon(m["c"]), true)
	drawSegment(c, ShrinkPolygon(m["d"]), true)
	drawSegment(c, ShrinkPolygon(m["e"]), false)
	drawSegment(c, ShrinkPolygon(m["f"]), true)
	drawSegment(c, ShrinkPolygon(m["g"]), true)

	err := c.SavePNG("images/ssd.png")
	checkError(err)
}

var (
	//drawSegment = drawSegment1
	drawSegment = drawSegment2
)

func drawSegment1(c *gg.Context, ps []Point2f, enabled bool) {

	hexagone.GGDrawPolygon(c, ps)

	if enabled {
		c.SetHexColor("#FF0000")
	} else {
		c.SetHexColor("#DDDDDD")
	}

	if true {
		c.FillPreserve()

		c.SetColor(color.White)
		//c.SetHexColor("#0000FF")
		c.Stroke()
	} else {
		c.Fill()
	}
}

func drawSegment2(c *gg.Context, ps []Point2f, enabled bool) {

	hexagone.GGDrawPolygon(c, ps)

	if enabled {
		c.SetHexColor("#FF0000")
	} else {
		c.SetHexColor("#DDDDDD")
	}
	c.Fill()
}

func lineNorm(a, b Point2f) Point2f {
	n := Point2f{
		X: (b.Y - a.Y),
		Y: -(b.X - a.X),
	}
	hyp := math.Hypot(n.X, n.Y)
	n = n.DivScalar(hyp)
	//n = pointMulScalar(n, -1)
	return n
}

func pointNorm(p Point2f) Point2f {
	return p.DivScalar(math.Hypot(p.X, p.Y))
}

//------------------------------------------------------------------------------

func drawSSD(c *gg.Context, x0, y0 float64, digitHeight float64, v uint8) error {

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

	c.Translate(x0, y0)
	c.Scale(dx, dy)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineWidth(lineWidth)

	var (
		// x1 = 0.5
		// x2 = 3.5

		// y1 = 1.0
		// y2 = 4.0
		// y3 = 7.0

		x1 = 1.0
		x2 = 3.0

		y1 = 2.0
		y2 = 4.0
		y3 = 6.0
	)

	c.SetColor(color.Black)

	segmentWidth := 0.5

	if segmentIsOn(v, 'a') {
		drawSegmentTrig(c, Pt(x1, y1), Pt(x2, y1), segmentWidth)
	}
	if segmentIsOn(v, 'b') {
		drawSegmentTrig(c, Pt(x2, y1), Pt(x2, y2), segmentWidth)
	}
	if segmentIsOn(v, 'c') {
		drawSegmentTrig(c, Pt(x2, y2), Pt(x2, y3), segmentWidth)
	}
	if segmentIsOn(v, 'd') {
		drawSegmentTrig(c, Pt(x2, y3), Pt(x1, y3), segmentWidth)
	}
	if segmentIsOn(v, 'e') {
		drawSegmentTrig(c, Pt(x1, y3), Pt(x1, y2), segmentWidth)
	}
	if segmentIsOn(v, 'f') {
		drawSegmentTrig(c, Pt(x1, y2), Pt(x1, y1), segmentWidth)
	}
	if segmentIsOn(v, 'g') {
		drawSegmentTrig(c, Pt(x1, y2), Pt(x2, y2), segmentWidth)
	}

	return nil
}

func drawSegmentTrig(c *gg.Context, a, b Point2f, width float64) {

	if true {
		const (
			koef = 0.03
			//koef = 0.04
		)
		var (
			a1 = geom.PtLerp(a, b, koef)
			b1 = geom.PtLerp(b, a, koef)
		)
		a = a1
		b = b1
	}

	var (
		da     = math.Pi / 4
		radius = width / math.Sqrt2
	)

	// Point a:
	var (
		deltaA = b.Sub(a)
		angleA = math.Atan2(deltaA.Y, deltaA.X)

		a1 = a.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleA+da)))
		a2 = a.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleA-da)))
	)

	// Point b:

	var (
		deltaB = a.Sub(b)
		angleB = math.Atan2(deltaB.Y, deltaB.X)

		b1 = b.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleB+da)))
		b2 = b.Add(geom.PolarToCartesian(geom.ShPolar(radius, angleB-da)))
	)

	c.MoveTo(a.X, a.Y)
	c.LineTo(a1.X, a1.Y)
	c.LineTo(b2.X, b2.Y)
	c.LineTo(b.X, b.Y)
	c.LineTo(b1.X, b1.Y)
	c.LineTo(a2.X, a2.Y)
	c.LineTo(a.X, a.Y)

	c.Fill()
}
