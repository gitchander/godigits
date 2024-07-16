package main

import (
	"fmt"
	"image"
	"math"

	"github.com/gitchander/cairo"
)

func testPredatorCountdown() error {

	filename := "result.png"

	//n := 512

	size := image.Pt(512, 640).Add(image.Pt(30, 30))

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer c.Destroy()

	// scale := float64(n)

	// if true {
	// 	c.SetSourceRGB(1, 1, 1)
	// 	c.Rectangle(0, 0, scale, scale)
	// 	c.Fill()
	// }

	//drawInd1(c)
	//drawInd2(c)
	//drawInd3(c)
	//drawInd4(c)
	testUniversalIndicator(c)

	return surface.WriteToPNG(filename)
}

func drawPredatorCountdown() error {

	filename := "result.png"

	const (
		ratioDy = 1.3175965665236051
		//ratioDy        = math.Phi

		//ratioLineWidth = 0.05785123966942149
		ratioLineWidth = 0.06

		ratioRadius = 0.3218884120171674

		//ratioSeparator = 0.27
		ratioSeparator = 0.3
	)

	var (
		margin           = Pt2f(10, 10)
		indicatorsNumber = 4
	)

	var (
		indicatorWidth = 150.0

		indicatorHeight = indicatorWidth * ratioDy
		lineWidth       = indicatorWidth * ratioLineWidth
		radius          = indicatorWidth * ratioRadius
		separatorWidth  = indicatorWidth * ratioSeparator
	)

	var (
		indicatorAndSeparatorWidth = indicatorWidth + separatorWidth
		offsetX                    = indicatorAndSeparatorWidth
	)

	var indicatorsWidth = indicatorWidth * float64(indicatorsNumber)
	if indicatorsNumber > 0 {
		indicatorsWidth += separatorWidth * float64(indicatorsNumber-1)
	}

	size := image.Point{
		X: int(math.Ceil(2*margin.X + indicatorsWidth)),
		Y: int(math.Ceil(2*margin.Y + indicatorHeight)),
	}

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer c.Destroy()

	indicatorSize := Pt2f(indicatorWidth, indicatorHeight)

	r := Rectangle2f{
		Min: margin,
		Max: margin.Add(indicatorSize),
	}

	rnd := NewRandNow()

	if false {
		for i := 0; i < indicatorsNumber; i++ {
			var (
				v  = uint16(rnd.Uint32())
				rr = r.Add(Pt2f(float64(i)*offsetX, 0))
			)
			v = 0xFFFF
			drawUniversalIndicator(c, v, rr, lineWidth, radius)
		}
	} else {
		for i := 0; i < indicatorsNumber; i++ {
			var (
				v  = uint16(rnd.Uint32())
				rr = r.Add(Pt2f(float64(i)*offsetX, 0))
			)
			v = 0xffff
			drawIndicator(c, (i%4)+1, v, rr, lineWidth, radius)
		}
	}

	return surface.WriteToPNG(filename)
}

func drawPredatorCountdownOneSegment() error {

	filename := "result.png"

	size := image.Point{
		X: 512,
		Y: 512,
	}

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer c.Destroy()

	ps := []Point2f{
		Pt2f(100, 100),
		Pt2f(300, 150),
	}

	lineWidth := 100.0

	c.SetSourceRGB(0.8, 0, 0)

	drawSegment(c, ps[0], ps[1], lineWidth)

	return surface.WriteToPNG(filename)
}

type Segment struct {
	Enable bool
	A      Point2f
	B      Point2f
}

func reverseSegments(segments []Segment) {
	i, j := 0, (len(segments) - 1)
	for i < j {
		segments[i], segments[j] = segments[j], segments[i]
		i, j = (i + 1), (j - 1)
	}
}

func normalizeSegments(segments []Segment, lineWidth float64, r Rectangle2f) {

	var roundKoef = math.Pow10(6)

	var (
		normalizeX = func(x float64) float64 {
			return normalize(x, r.Min.X, r.Max.X)
		}
		normalizeY = func(y float64) float64 {
			return normalize(y, r.Min.Y, r.Max.Y)
		}
		round = func(a float64) float64 {
			return math.Round(a*roundKoef) / roundKoef
		}
		roundPoint = func(p Point2f) Point2f {
			return Point2f{
				X: round(p.X),
				Y: round(p.Y),
			}
		}
	)

	ns := make([]Segment, len(segments))

	for i, segment := range segments {
		var (
			a = Point2f{
				X: normalizeX(segment.A.X),
				Y: normalizeY(segment.A.Y),
			}
			b = Point2f{
				X: normalizeX(segment.B.X),
				Y: normalizeY(segment.B.Y),
			}
		)

		a = roundPoint(a)
		b = roundPoint(b)

		ns[i] = Segment{
			A: a,
			B: b,
		}
	}

	for _, segment := range ns {
		fmt.Printf("{ Enable: %t, A: Pt2f(%v, %v), B: Pt2f(%v, %v) },\n", true,
			segment.A.X, segment.A.Y,
			segment.B.X, segment.B.Y,
		)
	}

	ratioDy := r.Dy() / r.Dx()
	ratioLineWidth := lineWidth / r.Dx()

	fmt.Println("ratioDy ", ratioDy)
	fmt.Println("ratioLineWidth:", ratioLineWidth)
}

func drawSegments(dc *cairo.Canvas, segments []Segment, value uint16,
	r Rectangle2f, lineWidth float64, radius float64) {

	var (
		dx = r.Dx()
		dy = r.Dy()
	)

	var (
		x1 = r.Min.X
		y1 = r.Min.Y

		x2 = x1 + dx
		y2 = y1 + dy
	)

	debug := false

	dc.SetSourceRGB(0.1, 0.1, 0.1)

	if debug {
		dc.Rectangle(x1, y1, dx, dy)
	} else {
		curveRectanglePath(dc, x1, y1, x2, y2, radius)
	}

	if true {
		dc.Fill()
	} else {
		dc.SetLineCap(cairo.LINE_CAP_ROUND)
		dc.SetLineWidth(lineWidth)
		dc.Stroke()
	}

	for i, segment := range segments {
		if uint16GetBit(value, i) == 1 {

			var (
				a = Point2f{
					X: x1 + segment.A.X*dx,
					Y: y1 + segment.A.Y*dy,
				}
				b = Point2f{
					X: x1 + segment.B.X*dx,
					Y: y1 + segment.B.Y*dy,
				}
			)

			dc.SetSourceRGB(0.8, 0, 0)
			drawSegment(dc, a, b, lineWidth)
		}
	}
}

var (
	//drawSegment = drawSegment1
	drawSegment = drawSegment2
)

func drawSegment1(c *cairo.Canvas, a, b Point2f, lineWidth float64) {

	c.MoveTo(a.X, a.Y)
	c.LineTo(b.X, b.Y)

	c.SetSourceRGB(0.8, 0, 0)
	c.SetLineWidth(lineWidth)
	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.Stroke()
}

func drawSegment2(c *cairo.Canvas, a, b Point2f, lineWidth float64) {

	rlw := lineWidth / 2

	var (
		deltaEdge = b.Sub(a).MulScalar(rlw / Distance(a, b))
	)

	var (
		deltaA = Pt2f(deltaEdge.Y, -deltaEdge.X)

		a12 = a.Sub(deltaEdge)

		a1 = a12.Add(deltaA)
		a2 = a12.Sub(deltaA)
	)

	var (
		deltaB = Pt2f(-deltaEdge.Y, deltaEdge.X)

		b12 = b.Add(deltaEdge)

		b2 = b12.Sub(deltaB)
		b1 = b12.Add(deltaB)
	)

	var (
		a1b2 = a1.Add(b2).DivScalar(2)
		b1a2 = b1.Add(a2).DivScalar(2)
	)

	const (
		// deltaFactor: [0.0 ... 1.0]
		deltaFactor = 0.4
	)

	var (
		ae1 = a12.Add(deltaA.MulScalar(deltaFactor))
		ae2 = a12.Sub(deltaA.MulScalar(deltaFactor))

		be1 = b12.Add(deltaB.MulScalar(deltaFactor))
		be2 = b12.Sub(deltaB.MulScalar(deltaFactor))
	)

	dc := NewDonCanvas(c)

	dc.BesierQuad(ae1, a1, a1b2)
	dc.BesierQuad(a1b2, b2, be2)
	dc.BesierQuad(be1, b1, b1a2)
	dc.BesierQuad(b1a2, a2, ae2)

	c.Fill()

	if false {
		circleRadius := rlw * 0.1

		c.SetSourceRGB(0, 0.6, 0)
		dc.Circle(c, a, circleRadius)
		c.Fill()

		c.SetSourceRGB(1, 1, 0)
		dc.Circle(c, b, circleRadius)
		c.Fill()

		p := a12
		c.SetSourceRGB(0, 0, 0)
		dc.Circle(c, p, circleRadius)
		c.Fill()
	}
}

// ------------------------------------------------------------------------------
func drawInd1(dc *cairo.Canvas) error {

	if true {
		const (
			//filename = "images/ind1_02.png"
			filename = "images/ind1_02_curves.png"
		)
		m, err := cairo.NewSurfaceFromPNG(filename)
		if err != nil {
			return err
		}

		dc.SetSourceSurface(m, 0, 0)
		dc.Paint()
	}

	radius := 150.0

	size := Pt2f(512, 640)

	r := Rectangle2f{
		Min: Point2f{
			X: 23.0,
			Y: 23.0,
		},
		Max: Point2f{
			X: size.X - 23.0,
			Y: size.Y - 3.0,
		},
	}

	ratioRadius := radius / r.Dx()
	fmt.Println("ratioRadius", ratioRadius)

	debug := true

	dc.SetSourceRGB(0, 0, 0)
	if debug {
		dc.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
	} else {
		curveRectanglePath(dc, r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, radius)
	}

	if debug {
		dc.SetLineWidth(3)
		dc.Stroke()
	} else {
		dc.Fill()
	}

	segments := []Segment{
		{
			Enable: true,
			A:      Pt2f(240, 90),
			B:      Pt2f(243, 168),
		},
		{
			Enable: true,
			A:      Pt2f(278, 242),
			B:      Pt2f(335, 282),
		},
		{
			Enable: true,
			A:      Pt2f(177, 344),
			B:      Pt2f(222, 396),
		},
		{
			Enable: true,
			A:      Pt2f(334, 344),
			B:      Pt2f(294, 404),
		},
		{
			Enable: true,
			A:      Pt2f(252, 480),
			B:      Pt2f(264, 568),
		},
		{
			Enable: true,
			A:      Pt2f(320, 460),
			B:      Pt2f(322, 548),
		},
	}

	for _, segment := range segments {
		if segment.Enable {
			var (
				a = segment.A
				b = segment.B
			)
			dc.MoveTo(a.X, a.Y)
			dc.LineTo(b.X, b.Y)
		}
	}

	if debug {
		dc.SetSourceRGBA(0, 0, 1, 0.5)
	} else {
		dc.SetSourceRGBA(1, 0, 0, 1)
	}

	lineWidth := 35.0
	dc.SetLineWidth(lineWidth)
	dc.SetLineCap(cairo.LINE_CAP_ROUND)
	dc.Stroke()

	normalizeSegments(segments, lineWidth, r)

	return nil
}

// ------------------------------------------------------------------------------
func drawInd2(dc *cairo.Canvas) error {

	if true {
		const (
			//filename = "images/ind2_02.png"
			filename = "images/ind2_02_curves.png"
		)
		m, err := cairo.NewSurfaceFromPNG(filename)
		if err != nil {
			return err
		}

		dc.SetSourceSurface(m, 0, 0)
		dc.Paint()
	}

	radius := 150.0

	size := Pt2f(512, 640)

	// var (
	// 	x1 = 23.0
	// 	y1 = 23.0

	// 	x2 = size.X - 2
	// 	y2 = size.Y + 4
	// )

	r := Rectangle2f{
		Min: Point2f{
			X: 23.0,
			Y: 23.0,
		},
		Max: Point2f{
			X: size.X - 2,
			Y: size.Y + 4,
		},
	}

	debug := true

	dc.SetSourceRGB(0, 0, 0)
	if debug {
		dc.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
	} else {
		curveRectanglePath(dc, r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, radius)
	}

	dc.SetLineWidth(3)

	if debug {
		dc.Stroke()
	} else {
		dc.Fill()
	}

	segments := []Segment{
		{
			Enable: true,
			A:      Pt2f(250, 104),
			B:      Pt2f(256, 190),
		},
		{
			Enable: true,
			A:      Pt2f(342, 226),
			B:      Pt2f(410, 220),
		},
		{
			Enable: true,
			A:      Pt2f(296, 284),
			B:      Pt2f(350, 310),
		},
		{
			Enable: true,
			A:      Pt2f(302, 424),
			B:      Pt2f(358, 378),
		},
		{
			Enable: true,
			A:      Pt2f(125, 487),
			B:      Pt2f(200, 487),
		},
		{
			Enable: true,
			A:      Pt2f(262, 492),
			B:      Pt2f(284, 572),
		},
	}

	for _, segment := range segments {
		if segment.Enable {
			var (
				a = segment.A
				b = segment.B
			)
			dc.MoveTo(a.X, a.Y)
			dc.LineTo(b.X, b.Y)
		}
	}

	if debug {
		dc.SetSourceRGBA(0, 0, 1, 0.5)
	} else {
		dc.SetSourceRGBA(1, 0, 0, 1)
	}

	lineWidth := 35.0
	dc.SetLineWidth(lineWidth)
	dc.SetLineCap(cairo.LINE_CAP_ROUND)
	dc.Stroke()

	normalizeSegments(segments, lineWidth, r)

	return nil
}

// ------------------------------------------------------------------------------
func drawInd3(dc *cairo.Canvas) error {

	if true {
		const (
			//filename = "images/ind3_03.png"
			filename = "images/ind3_03_curves.png"
		)
		m, err := cairo.NewSurfaceFromPNG(filename)
		if err != nil {
			return err
		}

		dc.SetSourceSurface(m, 0, 0)
		dc.Paint()
	}

	radius := 150.0

	size := Pt2f(512, 640)

	r := Rectangle2f{
		Min: Point2f{
			X: 23.0,
			Y: 23.0,
		},
		Max: Point2f{
			X: size.X - 4,
			Y: size.Y,
		},
	}

	debug := true

	dc.SetSourceRGB(0, 0, 0)
	if debug {
		dc.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
	} else {
		curveRectanglePath(dc, r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, radius)
	}

	dc.SetLineWidth(3)

	if debug {
		dc.Stroke()
	} else {
		dc.Fill()
	}

	segments := []Segment{
		{
			Enable: true,
			A:      Pt2f(178, 110),
			B:      Pt2f(186, 192),
		},
		{
			Enable: true,
			A:      Pt2f(262, 136),
			B:      Pt2f(266, 210),
		},
		{
			Enable: true,
			A:      Pt2f(312, 274),
			B:      Pt2f(334, 330),
		},
		{
			Enable: true,
			A:      Pt2f(189, 396),
			B:      Pt2f(240, 438),
		},
		{
			Enable: true,
			A:      Pt2f(322, 434),
			B:      Pt2f(376, 378),
		},
		{
			Enable: true,
			A:      Pt2f(286, 492),
			B:      Pt2f(296, 558),
		},
	}

	for _, segment := range segments {
		if segment.Enable {
			var (
				a = segment.A
				b = segment.B
			)
			dc.MoveTo(a.X, a.Y)
			dc.LineTo(b.X, b.Y)
		}
	}

	if debug {
		dc.SetSourceRGBA(0, 0, 1, 0.5)
	} else {
		dc.SetSourceRGBA(1, 0, 0, 1)
	}
	lineWidth := 35.0
	dc.SetLineWidth(lineWidth)
	dc.SetLineCap(cairo.LINE_CAP_ROUND)
	dc.Stroke()

	normalizeSegments(segments, lineWidth, r)

	return nil
}

// ------------------------------------------------------------------------------
func drawInd4(dc *cairo.Canvas) error {

	if true {
		const (
			filename = "images/ind4_02.png"
			//filename = "images/ind4_02_curves.png"
		)
		m, err := cairo.NewSurfaceFromPNG(filename)
		if err != nil {
			return err
		}

		dc.SetSourceSurface(m, 0, 0)
		dc.Paint()
	}

	radius := 150.0

	size := Pt2f(512, 640)

	// var (
	// 	x1 = 12.0
	// 	y1 = 30.0

	// 	x2 = size.X - 16
	// 	y2 = size.Y - 12
	// )

	r := Rectangle2f{
		Min: Point2f{
			X: 12.0,
			Y: 30.0,
		},
		Max: Point2f{
			X: size.X - 16,
			Y: size.Y - 12,
		},
	}

	debug := true

	dc.SetSourceRGB(0, 0, 0)
	if debug {
		dc.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
	} else {
		curveRectanglePath(dc, r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, radius)
	}

	dc.SetLineWidth(3)

	if debug {
		dc.Stroke()
	} else {
		dc.Fill()
	}

	segments := []Segment{
		{
			Enable: true,
			A:      Pt2f(232, 112),
			B:      Pt2f(236, 180),
		},
		{
			Enable: true,
			A:      Pt2f(128, 208),
			B:      Pt2f(176, 232),
		},
		{
			Enable: true,
			A:      Pt2f(328, 204),
			B:      Pt2f(392, 210),
		},
		{
			Enable: true,
			A:      Pt2f(180, 318),
			B:      Pt2f(218, 254),
		},
		{
			Enable: true,
			A:      Pt2f(288, 246),
			B:      Pt2f(332, 294),
		},
		{
			Enable: true,
			A:      Pt2f(170, 384),
			B:      Pt2f(216, 434),
		},
		// {
		// 	Enable: true,
		// 	A:      Pt2f(322, 434),
		// 	B:      Pt2f(376, 378),
		// },
		// {
		// 	Enable: true,
		// 	A:      Pt2f(286, 492),
		// 	B:      Pt2f(296, 558),
		// },
	}

	for _, segment := range segments {
		if segment.Enable {
			var (
				a = segment.A
				b = segment.B
			)
			dc.MoveTo(a.X, a.Y)
			dc.LineTo(b.X, b.Y)
		}
	}

	if debug {
		dc.SetSourceRGBA(0, 0, 1, 0.5)
	} else {
		dc.SetSourceRGBA(1, 0, 0, 1)
	}

	lineWidth := 26.0
	dc.SetLineWidth(lineWidth)
	dc.SetLineCap(cairo.LINE_CAP_ROUND)
	dc.Stroke()

	normalizeSegments(segments, lineWidth, r)

	return nil
}

func testUniversalIndicator(dc *cairo.Canvas) error {

	if false {
		const (
			filename = "images/universal_indicator.png"
		)
		m, err := cairo.NewSurfaceFromPNG(filename)
		if err != nil {
			return err
		}

		dc.SetSourceSurface(m, 0, 0)
		dc.Paint()
	}

	radius := 150.0

	size := Pt2f(512, 640)

	// var (
	// 	x1 = 12.0
	// 	y1 = 30.0

	// 	x2 = size.X - 16
	// 	y2 = size.Y - 12
	// )

	r := Rectangle2f{
		Min: Point2f{
			X: 32.0,
			Y: 30.0,
		},
		Max: Point2f{
			X: size.X - 34,
			Y: size.Y - 16,
		},
	}

	debug := true

	dc.SetSourceRGB(0, 0, 0)
	if debug {
		dc.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
	} else {
		curveRectanglePath(dc, r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, radius)
	}

	dc.SetLineWidth(3)

	if debug {
		dc.Stroke()
	} else {
		dc.Fill()
	}

	segments := []Segment{
		{
			Enable: false,
			A:      Pt2f(185, 102),
			B:      Pt2f(190, 170),
		},
		{
			Enable: true,
			A:      Pt2f(248, 112),
			B:      Pt2f(254, 180),
		},
		{
			Enable: true,
			A:      Pt2f(122, 208),
			B:      Pt2f(172, 228),
		},
		{
			Enable: true,
			A:      Pt2f(322, 202),
			B:      Pt2f(386, 208),
		},
		{
			Enable: true,
			A:      Pt2f(182, 316),
			B:      Pt2f(226, 252),
		},
		{
			Enable: true,
			A:      Pt2f(288, 246),
			B:      Pt2f(332, 294),
		},
	}

	for i, segment := range segments {

		//shift := Pt2f(0, 0)
		shift := Pt2f(0, -6)
		//shift := Pt2f(-2, -5)
		//shift := Pt2f(-16, 4)

		var (
			a = segment.A
			b = segment.B
		)

		a = a.Add(shift)
		b = b.Add(shift)

		segments[i] = Segment{
			Enable: true,
			A:      a,
			B:      b,
		}
	}

	center := r.Center()
	downSegments := make([]Segment, len(segments))
	for i, segment := range segments {

		var (
			a = segment.A
			b = segment.B
		)

		a = center.MulScalar(2).Sub(a)
		b = center.MulScalar(2).Sub(b)

		downSegments[i] = Segment{
			Enable: true,
			A:      a,
			B:      b,
		}
	}

	reverseSegments(downSegments)

	segments = append(segments, downSegments...)

	for _, segment := range segments {
		if segment.Enable {
			var (
				a = segment.A
				b = segment.B
			)
			dc.MoveTo(a.X, a.Y)
			dc.LineTo(b.X, b.Y)
		}
	}

	if debug {
		dc.SetSourceRGBA(0, 0, 1, 0.5)
	} else {
		dc.SetSourceRGBA(1, 0, 0, 1)
	}

	lineWidth := 26.0
	dc.SetLineWidth(lineWidth)
	dc.SetLineCap(cairo.LINE_CAP_ROUND)
	dc.Stroke()

	normalizeSegments(segments, lineWidth, r)

	return nil
}

func drawIndicator(dc *cairo.Canvas, number int, value uint16,
	r Rectangle2f, lineWidth float64, radius float64) error {

	m := map[int][]Segment{
		// segment 1
		1: []Segment{
			{Enable: true, A: Pt2f(0.465665, 0.109121), B: Pt2f(0.472103, 0.236156)},
			{Enable: true, A: Pt2f(0.54721, 0.356678), B: Pt2f(0.669528, 0.421824)},
			{Enable: true, A: Pt2f(0.330472, 0.522801), B: Pt2f(0.427039, 0.607492)},
			{Enable: true, A: Pt2f(0.667382, 0.522801), B: Pt2f(0.581545, 0.620521)},
			{Enable: true, A: Pt2f(0.491416, 0.7443), B: Pt2f(0.517167, 0.887622)},
			{Enable: true, A: Pt2f(0.637339, 0.711726), B: Pt2f(0.641631, 0.855049)},
		},
		// segment 2
		2: []Segment{
			{Enable: true, A: Pt2f(0.466119, 0.130435), B: Pt2f(0.478439, 0.268921)},
			{Enable: true, A: Pt2f(0.655031, 0.326892), B: Pt2f(0.794661, 0.31723)},
			{Enable: true, A: Pt2f(0.560575, 0.42029), B: Pt2f(0.671458, 0.462158)},
			{Enable: true, A: Pt2f(0.572895, 0.645733), B: Pt2f(0.687885, 0.571659)},
			{Enable: true, A: Pt2f(0.209446, 0.747182), B: Pt2f(0.36345, 0.747182)},
			{Enable: true, A: Pt2f(0.49076, 0.755233), B: Pt2f(0.535934, 0.884058)},
		},
		// segment 3
		3: []Segment{
			{Enable: true, A: Pt2f(0.319588, 0.141005), B: Pt2f(0.336082, 0.273906)},
			{Enable: true, A: Pt2f(0.492784, 0.183144), B: Pt2f(0.501031, 0.303079)},
			{Enable: true, A: Pt2f(0.595876, 0.406807), B: Pt2f(0.641237, 0.497569)},
			{Enable: true, A: Pt2f(0.342268, 0.604538), B: Pt2f(0.447423, 0.672609)},
			{Enable: true, A: Pt2f(0.616495, 0.666126), B: Pt2f(0.727835, 0.575365)},
			{Enable: true, A: Pt2f(0.542268, 0.76013), B: Pt2f(0.562887, 0.867099)},
		},
		// segment 4
		4: []Segment{
			{Enable: true, A: Pt2f(0.454545, 0.137124), B: Pt2f(0.46281, 0.250836)},
			{Enable: true, A: Pt2f(0.239669, 0.297659), B: Pt2f(0.338843, 0.337793)},
			{Enable: true, A: Pt2f(0.652893, 0.29097), B: Pt2f(0.785124, 0.301003)},
			{Enable: true, A: Pt2f(0.347107, 0.481605), B: Pt2f(0.42562, 0.374582)},
			{Enable: true, A: Pt2f(0.570248, 0.361204), B: Pt2f(0.661157, 0.441472)},
			{Enable: true, A: Pt2f(0.326446, 0.591973), B: Pt2f(0.421488, 0.675585)},
		},
	}

	segments := m[number]

	drawSegments(dc, segments, value, r, lineWidth, radius)

	return nil
}

func drawUniversalIndicator(dc *cairo.Canvas, value uint16,
	r Rectangle2f, lineWidth float64, radius float64) error {

	segments := []Segment{
		{Enable: true, A: Pt2f(0.343049, 0.111111), B: Pt2f(0.35426, 0.225589)},
		{Enable: true, A: Pt2f(0.484305, 0.127946), B: Pt2f(0.497758, 0.242424)},
		{Enable: true, A: Pt2f(0.201794, 0.289562), B: Pt2f(0.313901, 0.323232)},
		{Enable: true, A: Pt2f(0.650224, 0.279461), B: Pt2f(0.793722, 0.289562)},
		{Enable: true, A: Pt2f(0.336323, 0.47138), B: Pt2f(0.434978, 0.363636)},
		{Enable: true, A: Pt2f(0.573991, 0.353535), B: Pt2f(0.672646, 0.434343)},
		{Enable: true, A: Pt2f(0.426009, 0.646465), B: Pt2f(0.327354, 0.565657)},
		{Enable: true, A: Pt2f(0.663677, 0.52862), B: Pt2f(0.565022, 0.636364)},
		{Enable: true, A: Pt2f(0.349776, 0.720539), B: Pt2f(0.206278, 0.710438)},
		{Enable: true, A: Pt2f(0.798206, 0.710438), B: Pt2f(0.686099, 0.676768)},
		{Enable: true, A: Pt2f(0.515695, 0.872054), B: Pt2f(0.502242, 0.757576)},
		{Enable: true, A: Pt2f(0.656951, 0.888889), B: Pt2f(0.64574, 0.774411)},
	}

	drawSegments(dc, segments, value, r, lineWidth, radius)

	return nil
}

// ------------------------------------------------------------------------------
func curveRectanglePath(dc *cairo.Canvas, x0, y0, x1, y1 float64, radius float64) {

	var (
		dx = x1 - x0
		dy = y1 - y0
	)

	if (dx <= 0) || (dy <= 0) {
		return
	}

	dc.NewPath()

	if dx/2 < radius {
		if dy/2 < radius {
			dc.MoveTo(x0, (y0+y1)/2)
			dc.CurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			dc.CurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			dc.CurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			dc.CurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			dc.MoveTo(x0, y0+radius)
			dc.CurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			dc.CurveTo(x1, y0, x1, y0, x1, y0+radius)
			dc.LineTo(x1, y1-radius)
			dc.CurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			dc.CurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	} else {
		if dy/2 < radius {
			dc.MoveTo(x0, (y0+y1)/2)
			dc.CurveTo(x0, y0, x0, y0, x0+radius, y0)
			dc.LineTo(x1-radius, y0)
			dc.CurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			dc.CurveTo(x1, y1, x1, y1, x1-radius, y1)
			dc.LineTo(x0+radius, y1)
			dc.CurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			dc.MoveTo(x0, y0+radius)
			dc.CurveTo(x0, y0, x0, y0, x0+radius, y0)
			dc.LineTo(x1-radius, y0)
			dc.CurveTo(x1, y0, x1, y0, x1, y0+radius)
			dc.LineTo(x1, y1-radius)
			dc.CurveTo(x1, y1, x1, y1, x1-radius, y1)
			dc.LineTo(x0+radius, y1)
			dc.CurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	}

	dc.ClosePath()
}
