package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	helloWorld()
}

func helloWorld() error {

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	c := gg.NewContext(512, 512)

	// fill background
	c.SetColor(color.White)
	c.Clear()

	// Draw anchor point
	c.SetRGB(1, 0, 0)
	c.DrawCircle(100, 100, 5)
	c.Fill()

	c.SetColor(color.Black)
	c.SetFontFace(face)

	c.DrawString("Hello, World!", 100, 100)

	return c.SavePNG("result.png")
}

func drawHelloWorld(c *gg.Context, x0, y0 float64, size float64) error {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	face := truetype.NewFace(font, &truetype.Options{Size: size})

	c.SetFontFace(face)

	text := "Hello, World!"

	textWidth, textHeight := c.MeasureString(text)

	c.DrawRectangle(x0, y0, textWidth, textHeight)
	c.SetRGB(0.9, 0.9, 0.9)
	c.Fill()

	c.SetColor(color.Black)
	c.DrawString(text, x0, y0+size)
	//c.Stroke()

	// Draw anchor point
	c.SetRGB(1, 0, 0)
	c.DrawCircle(x0, y0+size, 5)
	c.Fill()

	return nil
}
