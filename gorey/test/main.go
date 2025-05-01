package main

import (
	"image"
	"image/color"
	"path/filepath"

	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
	"github.com/gitchander/godigits/gorey"
	"github.com/gitchander/godigits/utils"
	"github.com/gitchander/godigits/utils/colorf"
)

func main() {
	checkError(makeObjectImage())
	checkError(makeDigitImage())
	checkError(makeDigitsImage())
	checkError(makeDigitsImageMatrix())
}

// parts of speech

func makeObjectImage() error {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	filename := filepath.Join(dirName, "result.png")

	var (
		size = image.Pt(456, 456)
		//size = image.Pt(512, 512)
	)

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

	r := geom.Rectangle2f{
		Max: geom.Pt2f(float64(size.X), float64(size.Y)),
	}

	if true {
		colorf.CairoSetSourceColor(c, color.White)
		cairoRectangle(c, r)
		c.Fill()
	}

	var o gorey.Object

	o = gorey.Invert{
		Content: gorey.HBox{
			Objects: []gorey.Object{
				gorey.RoundRect{
					Content: gorey.Caret{
						Content: gorey.Infinity{},
					},
				},
				gorey.Rectangle{
					Content: gorey.Triangle(gorey.Circle{}),
				},
			},
		},
	}

	gorey.DrawLevelArea = false

	if o != nil {
		o.Draw(c, r, 0)
	}

	return surface.WriteToPNG(filename)
}

func makeDigitImage() error {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	filename := filepath.Join(dirName, "result_digit.png")

	var (
		size = image.Pt(456, 456)
		//size = image.Pt(512, 512)
	)

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

	r := geom.Rectangle2f{
		Max: geom.Pt2f(float64(size.X), float64(size.Y)),
	}

	if true {
		colorf.CairoSetSourceColor(c, color.White)
		cairoRectangle(c, r)
		c.Fill()
	}

	var o gorey.Object

	o = gorey.DigitV1{
		Value: -19,
	}

	// o = gorey.DigitV4{
	// 	Value: 100,
	// }

	// o = gorey.DigitV4{
	// 	Value: -3 * (1 + 7 + 7*7 + 7*7*7),
	// }

	// o = gorey.RoundRect{
	// 	Content: gorey.Circle{
	// 		Content: gorey.Infinity{},
	// 	},
	// }

	if o != nil {
		o.Draw(c, r, 0)
	}

	return surface.WriteToPNG(filename)
}

func makeDigitsImage() error {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	filename := filepath.Join(dirName, "digits.png")
	digits := utils.MakeIntsByInterval(-20, 21)

	ds := image.Pt(30, 19).Mul(3)

	size := image.Pt(ds.X*len(digits), ds.Y)

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

	// if true {
	// 	c.SetSourceRGB(1, 1, 1)
	// 	c.Rectangle(0, 0, float64(n*len(digits)), scale)
	// 	c.Fill()
	// }

	var o gorey.Object

	rt := geom.Rectangle2f{
		Max: geom.Pt2f(float64(ds.X), float64(ds.Y)),
	}

	for i, digit := range digits {
		// c.Save()
		// c.Translate(float64(n*i), 0)
		// c.Scale(scale, scale)

		cr := rt.Add(geom.Pt2f(float64(i)*rt.Dx(), 0))

		if true {
			if (i % 2) == 0 {
				c.SetSourceRGB(1, 1, 1)
			} else {
				c.SetSourceRGB(0.7, 0.9, 1)
			}
			c.Rectangle(0, 0, 1, 1)
			c.Fill()
		}

		o = gorey.DigitV4{
			Value: digit,
		}

		o.Draw(c, cr, 0)
	}

	return surface.WriteToPNG(filename)
}

func makeDigitsImageMatrix() error {

	dirName := "images"
	utils.MustMkdirIfNotExist(dirName)

	filename := filepath.Join(dirName, "digits_matrix.png")

	var (
		xn = 4
		yn = 4
	)

	var (
		makeDigit func(v int) gorey.Object
		ds        image.Point
	)

	switch 2 {

	case 1:
		makeDigit = func(v int) gorey.Object {
			return gorey.DigitV1{
				Value: v,
			}
		}
		ds = image.Pt(64, 64)

	case 2:
		makeDigit = func(v int) gorey.Object {
			return gorey.DigitV2{
				Value: v,
			}
		}
		ds = image.Pt(164, 164)

	case 3:
		// makeDigit = func(v int) Object {
		// 	return DigitV3{
		// 		Value: v,
		// 	}
		// }
		// x := 64
		// ds = image.Pt(x, 2*x)

	case 4:
		makeDigit = func(v int) gorey.Object {
			return gorey.DigitV4{
				Value: v,
			}
		}
		ds = image.Pt(30, 19).Mul(3)
	}

	size := image.Pt(xn*ds.X, yn*ds.Y)

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

	if true {
		c.SetSourceRGB(1, 1, 1)
		c.Rectangle(0, 0, float64(size.X), float64(size.Y))
		c.Fill()
	}

	var o gorey.Object

	var (
		//intNexter IntNexter = SerialIntNexter()
		intNexter IntNexter = SerialIntNexterInit(0)
		//intNexter IntNexter = RandomIntNexter()
	)

	rt := geom.Rectangle2f{
		Max: geom.Pt2f(float64(ds.X), float64(ds.Y)),
	}

	for y := 0; y < yn; y++ {
		for x := 0; x < xn; x++ {

			p := geom.Pt2f(float64(x*ds.X), float64(y*ds.Y))
			cr := rt.Add(p)

			if true {
				if ((x + y) % 2) == 0 {
					c.SetSourceRGB(1, 1, 1)
				} else {
					c.SetSourceRGB(0.7, 0.9, 1)
				}
				cairoRectangle(c, cr)
				c.Fill()
			}

			v := intNexter.NextInt()

			o = makeDigit(v)

			o.Draw(c, cr, 0)
		}
	}

	return surface.WriteToPNG(filename)
}

func cairoRectangle(c *cairo.Canvas, r geom.Rectangle2f) {
	c.Rectangle(r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}
