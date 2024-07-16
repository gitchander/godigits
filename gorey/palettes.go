package gorey

import (
	"image/color"

	"github.com/gitchander/godigits/utils/colorf"
)

var (
	Black = color.Black
	White = color.White

	//Red   = MkfColor(1, 0, 0, 1)
	// Green = MkfColor(0, 0.5, 0, 1)
	// Blue  = MkfColor(0, 0, 1, 1)

	Red   = colorf.MustParseColor("#f00")
	Green = colorf.MustParseColor("#070")
	Blue  = colorf.MustParseColor("#00f")

	// Tp - transparent
	RedTp50 = colorf.MakeColorf(1, 0, 0, 0.5)
)

var (
	palette1 = []color.Color{
		colorf.MakeColorf(0.9, 0.1, 0, 0.5),
		colorf.MakeColorf(0, 0.3, 0.7, 0.5),
	}

	paletteContentAreas1 = []color.Color{
		colorf.MakeColorf(1.0, 0.5, 0.5, 1.0),
		colorf.MakeColorf(0.4, 1.0, 0.5, 1.0),
		colorf.MakeColorf(0.4, 0.6, 1.0, 1.0),
		colorf.MakeColorf(1.0, 1.0, 0.5, 1.0),
		colorf.MakeColorf(0.8, 0.5, 1.0, 1.0),
		colorf.MakeColorf(0.5, 1.0, 1.0, 1.0),
	}

	paletteContentAreas2 = []color.Color{
		colorf.MakeColorf(1.0, 0.0, 0.0, 0.5),
		colorf.MakeColorf(0.0, 1.0, 0.0, 0.5),
		colorf.MakeColorf(0.0, 0.0, 1.0, 0.5),
	}

	paletteContentAreas3 = []color.Color{
		colorf.MakeColorf(0.5, 0.5, 0.5, 1.0),
		colorf.MakeColorf(1.0, 1.0, 1.0, 1.0),
	}

	// Google Colors Color Palette
	// http://www.color-hex.com/color-palette/1872
	palette2 = []color.Color{
		colorf.MustParseColor("#008744"),
		colorf.MustParseColor("#0057e7"),
		colorf.MustParseColor("#d62d20"),
		colorf.MustParseColor("#ffa700"),
		colorf.MustParseColor("#ffffff"),
	}
)

func levelToColorPalette(palette []color.Color, level int) color.Color {
	return palette[level%(len(palette))]
}

func levelToColor(level int) color.Color {
	var (
		palette = paletteContentAreas1
		//palette = paletteContentAreas3
	)
	return levelToColorPalette(palette, level)
}
