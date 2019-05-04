package camerarenderer

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"

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

const (
	IMG_HEIGHT = 640
	IMG_WIDTH  = 800
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

	upLeft := image.Point{0, 0}
	lowRight := image.Point{IMG_WIDTH, IMG_HEIGHT}
	img := image.NewNRGBA(image.Rectangle{upLeft, lowRight})

	c := color.RGBA{R: 100, G: 100, B: 100, A: 200}

	rect := image.Rect(
		0, 0,
		IMG_WIDTH, IMG_HEIGHT,
	)

	draw.Draw(img, rect, &image.Uniform{c}, image.ZP, draw.Src)

	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	return buf.Bytes(), nil
}
