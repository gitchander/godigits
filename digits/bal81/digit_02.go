package bal81

import (
	"math"

	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/hexagone"
	"github.com/gitchander/godigits/utils/gobits"
)

type Digit2 struct{}

var _ dgdr.DigitDrawer = Digit2{}

func (Digit2) Width(height float64) (width float64) {
	width = height / 2
	return
}

func (d Digit2) DrawDigit(c *gg.Context, x, y float64, digitHeight float64, digit int) {

	var (
		nx = 4
		ny = 8

		//digitWidth = d.Width(digitHeight)

		w = digitHeight / float64(ny)

		greedWidth = 0.02 * w
		lineWidth  = 0.001 * w
	)

	c.Push()
	defer c.Pop()

	c.Translate(x, y)
	c.Scale(w, w)

	dgdr.DrawGreedGG(c, nx, ny, greedWidth)

	c.SetLineCap(gg.LineCapRound)
	c.SetLineJoin(gg.LineJoinRound)
	c.SetLineWidth(lineWidth)

	center := geom.Pt2f(2, 4)

	var (
		// radius = 1.75
		// v      = digitSegmentsV1(digit)

		// radius = 2.0
		// v      = digitSegmentsV1(digit)

		// radius = 1.5
		// v      = digitSegmentsV2(digit)

		radius = float64(nx) / math.Sqrt(3)
		v      = digitSegmentsV1(digit)
	)

	const segmentAngle = math.Pi / 3
	lineWidth = radius * 0.175
	vertexes := hexagone.HexagoneVertexes(center, radius)

	drawSegment := func(a, b geom.Point2f) {
		a, b = hexagone.ShrinkPoints(a, b, 0.03)
		drawSegmentGG(c, a, b, lineWidth, segmentAngle)

		// lineWidth := w / 2.5
		// shrinkWidth := lineWidth * 0.04
		// a, b = hexagone.ShrinkPoints(a, b, shrinkWidth)
		// c.MoveTo(a.X, a.Y)
		// c.LineTo(b.X, b.Y)
		// //c.SetLineCap(gg.LineCapRound)
		// c.SetLineCap(gg.LineCapSquare)
		// c.SetLineWidth(lineWidth)
		// c.Stroke()
	}

	// Rays:
	for i, vertex := range vertexes {
		if gobits.GetBit(v, i) == 1 {
			drawSegment(center, vertex)
		}
	}

	// Edges:
	for i, vertex := range vertexes {
		if gobits.GetBit(v, 6+i) == 1 {
			next := vertexes[(i+1)%len(vertexes)]
			drawSegment(vertex, next)
		}
	}
}

func drawSegmentGG(c *gg.Context, a, b geom.Point2f, lineWidth float64, angle float64) {
	hexagone.GGDrawSegment(c, a, b, lineWidth, angle)
}

func digitSegmentsV1(digit int) uint16 {

	bs := CalcDigitsBal3(digit)

	var x uint16

	var (
		ons = []int{0, 3}

		positives = []int{6, 1, 2, 8}
		negatives = []int{11, 5, 4, 9}
	)

	for _, on := range ons {
		x = gobits.SetBit(x, on, 1)
	}
	n := minInt(4, len(bs))
	for i := 0; i < n; i++ {
		if bs[i] == 1 {
			x = gobits.SetBit(x, positives[i], 1)
		} else if bs[i] == -1 {
			x = gobits.SetBit(x, negatives[i], 1)
		}
	}

	return x
}

func digitSegmentsV2(digit int) uint16 {

	if (digit < -13) || (13 < digit) {
		return 0
	}

	bs := CalcDigitsBal3(digit)

	var x uint16

	//--------------------------------------------------------------------------

	// var (
	// 	ons = []int{0, 2, 4}

	// 	positives = []int{6, 8, 10}
	// 	negatives = []int{7, 9, 11}
	// )

	//--------------------------------------------------------------------------

	// var (
	// 	ons = []int{0, 2, 4}

	// 	positives = []int{6, 8, 10}
	// 	negatives = []int{11, 7, 9}
	// )

	//--------------------------------------------------------------------------

	// var (
	// 	ons = []int{0, 2, 4}

	// 	positives = []int{8, 10, 6}
	// 	negatives = []int{9, 11, 7}
	// )

	//--------------------------------------------------------------------------

	var (
		ons = []int{1, 3, 5}

		positives = []int{6, 8, 10}
		negatives = []int{11, 7, 9}
	)

	//--------------------------------------------------------------------------

	// var (
	// 	ons = []int{6, 7, 8, 9, 10, 11}

	// 	positives = []int{0, 2, 4}
	// 	negatives = []int{3, 5, 1}
	// )

	//--------------------------------------------------------------------------

	for _, on := range ons {
		x = gobits.SetBit(x, on, 1)
	}
	n := minInt(3, len(bs))
	for i := 0; i < n; i++ {
		if bs[i] == 1 {
			x = gobits.SetBit(x, positives[i], 1)
		} else if bs[i] == -1 {
			x = gobits.SetBit(x, negatives[i], 1)
		}
	}

	return x
}
