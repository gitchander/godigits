package main

import (
	"image/color"

	"github.com/gitchander/cairo"
)

const maxColorComponent = 0xffff

// Don don :)
type DonCanvas struct {
	c *cairo.Canvas
}

func NewDonCanvas(c *cairo.Canvas) *DonCanvas {
	return &DonCanvas{
		c: c,
	}
}

func (dc *DonCanvas) MoveTo(p Point2f) {
	dc.c.MoveTo(p.X, p.Y)
}

func (dc *DonCanvas) LineTo(p Point2f) {
	dc.c.LineTo(p.X, p.Y)
}

func (dc *DonCanvas) SetSourceColor(cl color.Color) {

	uc := color.NRGBA64Model.Convert(cl).(color.NRGBA64)

	var (
		fR = float64(uc.R) / maxColorComponent
		fG = float64(uc.G) / maxColorComponent
		fB = float64(uc.B) / maxColorComponent
		fA = float64(uc.A) / maxColorComponent
	)
	if uc.A == maxColorComponent {
		dc.c.SetSourceRGB(fR, fG, fB)
	} else {
		dc.c.SetSourceRGBA(fR, fG, fB, fA)
	}
}

func (dc *DonCanvas) Circle(c *cairo.Canvas, center Point2f, radius float64) {
	dc.c.Arc(center.X, center.Y, radius, 0, tau)
}

// Quadratic Bézier
func (dc *DonCanvas) BesierQuad(p0, p1, p2 Point2f) {

	const koef = 2.0 / 3.0

	var (
		P01 = PtLerp(p0, p1, koef)
		P21 = PtLerp(p2, p1, koef)
	)

	dc.BesierCubic(p0, P01, P21, p2)
}

func (dc *DonCanvas) BesierCubic(p0, p1, p2, p3 Point2f) {
	dc.LineTo(p0)
	dc.c.CurveTo(p1.X, p1.Y, p2.X, p2.Y, p3.X, p3.Y)
}
