package gorey

import (
	"image"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/dgdr"
	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/utils/digits"
)

type DigitV4 struct {
	Value int
}

var _ Object = DigitV4{}

func (DigitV4) IsObject() {}

func (d DigitV4) Draw(c *cairo.Canvas, r geom.Rectangle2f, level int) {

	const (
		usePoints = false
	)

	const (
		lineWidthRel = 0.8
		//lineWidthRel = 1.0
	)

	var (
		marginRel = geom.MakeFrame1(0)
		//marginRel = MakeFrame1(5)
	)

	var (
		vm = vmin(r)

		marginAbs = marginRel.MulScalar(vm)
	)

	//bunchSize := image.Pt()

	var (
		nx = 30
		ny = 19

		greedWidth = 0.04
	)

	dSize := image.Pt(nx, ny)

	aspectRatio := float64(dSize.Y) / float64(dSize.X)

	var (
		r1 = r.Shrink(marginAbs)
		r2 = subRectByAspectRatio(r1, aspectRatio)

		cellSize = geom.Point2f{
			X: r2.Dx() / float64(dSize.X),
			Y: r2.Dy() / float64(dSize.Y),
		}

		wLineWidth = lineWidthRel
		radius     = wLineWidth
	)

	//wLineWidth = 1.0

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

	c.Save()
	defer c.Restore()

	c.Translate(r2.Min.X, r2.Min.Y)
	c.Scale(cellSize.X, cellSize.Y)

	dgdr.DrawGreedCairo(c, nx, ny, greedWidth)

	cairoSetSourceColor(c, Black)
	c.SetLineWidth(wLineWidth)
	c.SetLineCap(cairo.LINE_CAP_ROUND)

	//--------------------------------------------------------------------------

	const base = 4
	const maxValue = base*base - 1

	// val := d.Value
	// //val = (val+maxValue)%(2*maxValue+1) - maxValue
	// negative := false
	// if val < 0 {
	// 	negative = true
	// 	val = -val
	// }

	//var ny int

	//--------------------------------------------------------------------------

	x0 := 2.0
	y0 := 1.0

	dx := 6.0
	dy := 2.0

	var (
		lineDx1 = -1.5
		lineDx2 = 1.5

		// lineDx1 = -2.0
		// lineDx2 = 2.0

		// lineDx1 = -3.0
		// lineDx2 = 0.0

		//--------------------------------------
		// lineDy1 = 0.0
		// lineDy2 = 0.0

		// lineDy1 = -1.0
		// lineDy2 = 0.0

		// lineDy1 = 0.0
		// lineDy2 = -2.0

		lineDy1 = 0.0
		lineDy2 = -1.5
		//--------------------------------------
	)

	x1 := x0
	y1 := y0 + 4

	y2 := y1 + 5
	y3 := y2 + 7

	var (
		circleDx = 4.0
		circleDy = 4.0

		// circleDx = 3.0
		// circleDy = 5.0
	)

	rd := digits.MustNewRestDigiter(-3, +3)

	bus, _ := digits.CalcDigitsN(rd, d.Value, 10)
	if false {
		const nm = 2
		for len(bus) < nm {
			bus = append(bus, 0)
		}
		if len(bus) > nm {
			bus = bus[:nm]
		}
	}

	for _, bunch := range bus {

		x2 := x1 + dx

		c.MoveTo(x1, y1)

		switch 0 {
		case 0:
			BesierQuad(c, geom.Pt2f(x1, y1), geom.Pt2f(x2, y0), geom.Pt2f(x2, y1))
		case 1:
			BesierCubic(c, geom.Pt2f(x1, y1), geom.Pt2f(x1+4, y0+1), geom.Pt2f(x2, y0+2), geom.Pt2f(x2, y1))
		}

		c.LineTo(x2, y3)
		c.Stroke()

		negative := false
		ny := bunch
		if ny < 0 {
			negative = true
			ny = -ny
		}

		if !usePoints {
			y := y2
			for j := 0; j < ny; j++ {
				c.MoveTo(x2+lineDx1, y+lineDy1)
				c.LineTo(x2+lineDx2, y+lineDy2)
				y += dy
			}
			c.Stroke()

			if negative {
				c.Circle(x1+circleDx, y0+circleDy, radius)
				c.Fill()
			}

		} else {
			dy := 3.0
			y := y2 - 1
			for j := 0; j < ny; j++ {
				c.Circle(x2-3, y, radius)
				y += dy
			}
			c.Fill()

			if negative {
				c.MoveTo(x1+circleDx-1, y0)
				c.LineTo(x1+circleDx-1, y0+5)

				// c.MoveTo(x1+circleDx-3, y0+5)
				// c.LineTo(x1+circleDx+1, y0+5)

				// c.MoveTo(x1+circleDx-3, y0+1)
				// c.LineTo(x1+circleDx-1, y0+5)
			}
			c.Stroke()
		}

		x1 += dx
	}
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}
