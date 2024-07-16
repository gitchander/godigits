package gorey

import (
	"github.com/gitchander/cairo"

	"github.com/gitchander/godigits/geom"
)

type Object interface {
	IsObject()

	Draw(c *cairo.Canvas, r geom.Rectangle2f, level int)
}
