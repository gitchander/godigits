package dgdr

import (
	"github.com/fogleman/gg"
	"github.com/gitchander/cairo"
)

// DrawOver
// DrawerOn

type DrawerCairo interface {
	DrawCairo(c *cairo.Canvas)
}

type DrawerGG interface {
	DrawGG(c *gg.Context)
}

type DrawerCairoGG interface {
	DrawerCairo
	DrawerGG
}
