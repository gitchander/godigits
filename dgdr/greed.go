package dgdr

import (
	"github.com/fogleman/gg"
	"github.com/gitchander/cairo"
)

var DrawGreedEnable = false

func DrawGreedGG(c *gg.Context, nx, ny int, greedWidth float64) {

	if !DrawGreedEnable {
		return
	}

	c.Push()
	defer c.Pop()

	// Horizontal lines
	{
		var (
			fX0 = float64(0)
			fX1 = float64(nx)
		)
		for i := 0; i <= ny; i++ {
			fY := float64(i)
			c.MoveTo(fX0, fY)
			c.LineTo(fX1, fY)
		}
	}

	// Vertical lines
	{
		var (
			fY0 = float64(0)
			fY1 = float64(ny)
		)
		for i := 0; i <= nx; i++ {
			fX := float64(i)
			c.MoveTo(fX, fY0)
			c.LineTo(fX, fY1)
		}
	}

	c.SetLineWidth(greedWidth)
	c.SetRGB(0, 0, 1)
	c.Stroke()
}

func DrawGreedCairo(c *cairo.Canvas, nx, ny int, greedWidth float64) {

	if !DrawGreedEnable {
		return
	}

	c.Save()
	defer c.Restore()

	// Horizontal lines
	{
		var (
			fX0 = float64(0)
			fX1 = float64(nx)
		)
		for i := 0; i <= ny; i++ {
			fY := float64(i)
			c.MoveTo(fX0, fY)
			c.LineTo(fX1, fY)
		}
	}

	// Vertical lines
	{
		var (
			fY0 = float64(0)
			fY1 = float64(ny)
		)
		for i := 0; i <= nx; i++ {
			fX := float64(i)
			c.MoveTo(fX, fY0)
			c.LineTo(fX, fY1)
		}
	}

	c.SetLineWidth(greedWidth)
	c.SetSourceRGB(0, 0, 1)
	c.Stroke()
}
