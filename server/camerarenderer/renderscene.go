package camerarenderer

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

func (r *Renderer) RenderScene(x, y, z, direction, zdirection int) ([]byte, error) {
	return nil, nil
}
