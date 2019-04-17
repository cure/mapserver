package camerarenderer

import (
	"mapserver/colormapping"
	"mapserver/mapblockaccessor"
)

type Renderer struct {
	BlockAccessor *mapblockaccessor.MapBlockAccessor
	Colormapping  *colormapping.ColorMapping
}

func NewRenderer(BlockAccessor *mapblockaccessor.MapBlockAccessor,
	Colormapping *colormapping.ColorMapping) *Renderer {

	return &Renderer{
		BlockAccessor: BlockAccessor,
		Colormapping:  Colormapping,
	}
}
