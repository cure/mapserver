package mapblockrenderer

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"mapserver/colormapping"
	"mapserver/coords"
	"mapserver/blockaccessor"
	"time"
	"math"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type MapBlockRenderer struct {
	accessor           *blockaccessor.BlockAccessor
	colors             *colormapping.ColorMapping
	enableShadow       bool
	enableTransparency bool
}

func NewMapBlockRenderer(accessor *blockaccessor.BlockAccessor, colors *colormapping.ColorMapping) *MapBlockRenderer {
	return &MapBlockRenderer{
		accessor:           accessor,
		colors:             colors,
		enableShadow:       true,
		enableTransparency: false,
	}
}

const (
	IMG_SCALE                         = 16
	IMG_SIZE                          = IMG_SCALE * 16
	EXPECTED_BLOCKS_PER_FLAT_MAPBLOCK = 16 * 16
)

func IsViewBlocking(nodeName string) bool {
	if nodeName == "" {
		return false
	}

	if nodeName == "vacuum:vacuum" {
		return false
	}

	if nodeName == "air" {
		return false
	}

	return true
}

func clamp(num int) uint8 {
	if num < 0 {
		return 0
	}

	if num > 255 {
		return 255
	}

	return uint8(num)
}

func addColorComponent(c *color.RGBA, value int) *color.RGBA {
	return &color.RGBA{
		R: clamp(int(c.R) + value),
		G: clamp(int(c.G) + value),
		B: clamp(int(c.B) + value),
		A: clamp(int(c.A) + value),
	}
}

func (r *MapBlockRenderer) Render(pos1, pos2 *coords.MapBlockCoords) (*image.NRGBA, error) {
	if pos1.X != pos2.X {
		return nil, errors.New("X does not line up")
	}

	if pos1.Z != pos2.Z {
		return nil, errors.New("Z does not line up")
	}

	renderedMapblocks.Inc()
	timer := prometheus.NewTimer(renderDuration)
	defer timer.ObserveDuration()

	start := time.Now()
	defer func() {
		t := time.Now()
		elapsed := t.Sub(start)
		log.WithFields(logrus.Fields{"elapsed": elapsed}).Debug("Rendering completed")
	}()

	upLeft := image.Point{0, 0}
	lowRight := image.Point{IMG_SIZE, IMG_SIZE}
	img := image.NewNRGBA(image.Rectangle{upLeft, lowRight})

	maxY := pos1.Y
	minY := pos2.Y

	if minY > maxY {
		maxY, minY = minY, maxY
	}

	foundBlocks := 0
	xzOccupationMap := make([][]bool, 16)
	for x := range xzOccupationMap {
		xzOccupationMap[x] = make([]bool, 16)
	}

	fromX := pos1.X * 16
	toX := fromX + 16

	fromZ := pos1.Z * 16
	toZ := fromZ + 16

	toY := minY*16
	fromY := (maxY*16) + 16

	for x := fromX; x < toX; x++ {
		for z := fromZ; z < toZ; z++ {
			for y := fromY; y >= toY; y-- {
				x_mod := int(math.Abs(float64(x%16)))
				z_mod := int(math.Abs(float64(z%16)))

				if xzOccupationMap[x_mod][z_mod] {
					break
				}

				block, err := r.accessor.GetBlock(x,y,z)
				if err != nil {
					return nil, err
				}

				if block == nil || block.Name == "" {
					continue
				}

				c := r.colors.GetColor(block.Name, block.Param2)

				if c == nil {
					continue
				}

				if r.enableShadow {
					left, err := r.accessor.GetBlock(x-1,y,z)
					if err != nil {
						return nil, err
					}

					leftAbove, err := r.accessor.GetBlock(x-1,y+1,z)
					if err != nil {
						return nil, err
					}

					top, err := r.accessor.GetBlock(x,y,z-1)
					if err != nil {
						return nil, err
					}

					topAbove, err := r.accessor.GetBlock(x,y+1,z-1)
					if err != nil {
						return nil, err
					}

					if leftAbove != nil && IsViewBlocking(leftAbove.Name) {
						//add shadow
						c = addColorComponent(c, -10)
					}

					if topAbove != nil && IsViewBlocking(topAbove.Name) {
						//add shadow
						c = addColorComponent(c, -10)
					}

					if left != nil && !IsViewBlocking(left.Name) {
						//add light
						c = addColorComponent(c, 10)
					}

					if top != nil && !IsViewBlocking(top.Name) {
						//add light
						c = addColorComponent(c, 10)
					}
				}

				imgX := x * IMG_SCALE
				imgY := (15 - z) * IMG_SCALE

				rect := image.Rect(
					imgX, imgY,
					imgX+IMG_SCALE, imgY+IMG_SCALE,
				)

				if c.A != 0xFF || !r.enableTransparency {
					//not transparent, mark as rendered
					foundBlocks++
					xzOccupationMap[x_mod][z_mod] = true
				}

				draw.Draw(img, rect, &image.Uniform{c}, image.ZP, draw.Src)

				if foundBlocks == EXPECTED_BLOCKS_PER_FLAT_MAPBLOCK {
					return img, nil
				}
			}
		}
	}

	if foundBlocks == 0 {
		return nil, nil
	}

	return img, nil
}
