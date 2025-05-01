package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Bounds = geom.Rectangle2f

type Object interface {
	IsObject()

	Draw(c *cairo.Canvas, b Bounds, level int)
}
