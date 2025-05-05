package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Greed struct {
	Cells [][]Object
}

var _ Object = Greed{}

func (Greed) IsObject() {}

func (v Greed) Draw(c *cairo.Canvas, r Bounds, level int) {

	var (
		w = minFloat64(r.Dx(), r.Dy())

		margin = geom.MakeFrame1(0.05).MulScalar(w)
		//margin = geom.MakeFrame4(0.1, 0.0, 0.0, 0.0).MulScalar(w)

		sep = geom.Pt2f(0.02, 0.02).MulScalar(w)
	)

	var (
		yn = len(v.Cells)
		xn = 0
	)
	for _, rowCells := range v.Cells {
		xn = maxInt(xn, len(rowCells))
	}

	var (
		dy = (r.Dy() - (margin.Top + margin.Bottom) + sep.Y) / float64(yn)
		dx = (r.Dx() - (margin.Left + margin.Right) + sep.X) / float64(xn)

		cellSize = geom.Pt2f(dx, dy).Sub(sep)
	)

	p0 := r.Min.Add(geom.Pt2f(margin.Left, margin.Top))

	drawAreas := true

	for yi, rowCells := range v.Cells {
		for xi, cell := range rowCells {

			p := p0.Add(geom.Pt2f(float64(xi)*dx, float64(yi)*dy))
			cr := geom.Bounds{
				Min: p,
				Max: p.Add(cellSize),
			}

			if drawAreas {

				if true {
					if ((yi + xi) % 2) == 0 {
						cairoSetSourceColor(c, palette1[0])
					} else {
						cairoSetSourceColor(c, palette1[1])
					}
				} else {
					if ((yi + xi) % 2) == 0 {
						cairoSetSourceColor(c, Black)
					} else {
						cairoSetSourceColor(c, White)
					}
				}

				cairoRectangle(c, cr)
				c.Fill()
			}

			if cell != nil {
				cell.Draw(c, cr, level+1)
			}
		}
	}
}
