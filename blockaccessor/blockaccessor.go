package blockaccessor

import (
	"mapserver/coords"
	"mapserver/mapblockaccessor"
	"math"
)

func New(mba *mapblockaccessor.MapBlockAccessor) *BlockAccessor {
	return &BlockAccessor{mba: mba}
}

type BlockAccessor struct {
	mba *mapblockaccessor.MapBlockAccessor
}

type Block struct {
	Name string
	Param2 int
}

func (this *BlockAccessor) GetBlock(x, y, z int) (*Block, error) {

	mbc := coords.NewMapBlockCoordsFromBlock(x, y, z)
	mapblock, err := this.mba.GetMapBlock(mbc)

	if err != nil {
		return nil, err
	}

	if mapblock == nil {
		return nil, nil
	}

	relx := int(math.Abs(float64(x % 16)))
	rely := int(math.Abs(float64(y % 16)))
	relz := int(math.Abs(float64(z % 16)))

	block := Block{
		Name: mapblock.GetNodeName(relx, rely, relz),
		Param2: mapblock.GetParam2(relx, rely, relz),
	}

	return &block, nil
}

// TODO: GetMeta()
