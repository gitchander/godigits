package gorey

import (
	"image"
	"image/color"
	"math"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type HexagoneOrientation int

const (
	Angled HexagoneOrientation = iota
	Flat
)

var halfSqrt3 = math.Sqrt(3) / 2

var hexagoneVertexesAngled = [6]geom.Point2f{
	geom.Pt2f(0.0, 1.0),
	geom.Pt2f(halfSqrt3, 0.5),
	geom.Pt2f(halfSqrt3, -0.5),
	geom.Pt2f(0.0, -1.0),
	geom.Pt2f(-halfSqrt3, -0.5),
	geom.Pt2f(-halfSqrt3, 0.5),
}

var hexagoneVertexesFlat = [6]geom.Point2f{
	geom.Pt2f(-0.5, halfSqrt3),
	geom.Pt2f(0.5, halfSqrt3),
	geom.Pt2f(1, 0),
	geom.Pt2f(0.5, -halfSqrt3),
	geom.Pt2f(-0.5, -halfSqrt3),
	geom.Pt2f(-1, 0),
}

const hexScaleFactor = 0.9

type hexagoneInfo struct {
	FillColor   color.Color
	StrokeColor color.Color
	Orientation HexagoneOrientation
	LineWidth   float64
}

func makeImages(info hexagoneInfo, size image.Point, filename string) error {

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer canvas.Destroy()

	err = drawHexagone(canvas, info)
	if err != nil {
		return err
	}

	if err = surface.WriteToPNG(filename); err != nil {
		return err
	}

	return nil
}

func drawHexagone(c *cairo.Canvas, info hexagoneInfo) error {

	var (
		surface = c.GetTarget()
		nx      = surface.GetWidth()
		ny      = surface.GetHeight()
	)

	radius := hexScaleFactor * (float64(minInt(nx, ny)) * 0.5)

	m := cairo.NewMatrix()
	m.InitIdendity()
	m.InitTranslate(float64(nx)*0.5, float64(ny)*0.5)
	m.Scale(radius, radius)
	c.SetMatrix(m)

	c.SetLineWidth(info.LineWidth)

	c.SetLineJoin(cairo.LINE_JOIN_ROUND)
	c.SetLineCap(cairo.LINE_CAP_ROUND)

	var points [6]geom.Point2f
	if info.Orientation == Angled {
		points = hexagoneVertexesAngled
	} else {
		points = hexagoneVertexesFlat
	}

	p := points[len(points)-1]
	c.MoveTo(p.X, p.Y)

	for _, p := range points {
		c.LineTo(p.X, p.Y)
	}

	cairoSetSourceColor(c, info.FillColor)
	c.FillPreserve()

	cairoSetSourceColor(c, info.StrokeColor)
	c.Stroke()

	return nil
}

// type Hexagone struct {
// 	Orientation HexagoneOrientation
// }

// var _ Object = Hexagone{}

// func (h Hexagone) IsObject() {}

// func (h Hexagone) Draw(*cairo.Canvas) {
// 	// todo
// }
