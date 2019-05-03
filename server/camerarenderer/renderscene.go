package camerarenderer

import (
	"mapserver/coords"
)

const (
	NORTH_EAST = iota
	SOUTH_EAST = iota
	NORTH_WEST = iota
	SOUTH_WEST = iota
)

const (
	UP   = iota
	DOWN = iota
)

func (r *Renderer) GetNodeName(x, y, z int) string {
	return "" //TODO
}

func (r *Renderer) IsOccupied(x, y, z int) bool {
	coord := coords.GetMapBlockCoordsFromPlain(x, y, z)
	mb, err := r.BlockAccessor.GetMapBlock(coord)

	if err != nil {
		panic(err)
	}

	if mb == nil || mb.IsEmpty() {
		return false
	}

	return false //TODO
}

func (r *Renderer) RenderScene(x, y, z int, direction, zdirection int) ([]byte, error) {
	//10 mapblocks = 160 blocks distance
	// = 10^3 mapblocks = 1000

	return nil, nil
}
