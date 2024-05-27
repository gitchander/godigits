package main

import (
	"github.com/fogleman/gg"

	"github.com/gitchander/godigits/bal81"
	"github.com/gitchander/godigits/geom"
)

func main() {

	nA := 50
	dA := float64(nA)

	d := bal81.Digit1{
		A: dA,
		B: dA * 0.22,
		C: dA * 0.22,
	}.DigitDrawer()

	c := gg.NewContext(nA*8, nA*4)

	if true {
		c.SetRGB(1, 1, 1)
		c.Clear()
	}

	c.SetRGB(0, 0, 0)

	b := geom.MakeBounds(0, 0, 2*dA, 4*dA)
	bsh := geom.MakePoint(2*dA, 0)

	for i := 0; i < 4; i++ {
		d.DrawDigit(c, b, 0)
		b = b.Add(bsh)
	}

	c.SavePNG("result.png")
}
