package bitmap

type BitMap struct {
	bitMap []uint8
	size   uint64
}

func NewBitMap(size uint64) *BitMap {
	return &BitMap{
		size:   size,
		bitMap: make([]uint8, size/8+1),
	}
}

func (b *BitMap) Set(value uint64) {
	if value > b.size {
		return
	}
	bucket := value / 8
	index := value % 8
	b.bitMap[bucket] |= 1 << index
}

func (b *BitMap) Unset(value uint64) {
	if value > b.size || !b.IsExist(value) {
		return
	}
	bucket := value / 8
	index := value % 8
	//b.bitMap[bucket] ^= 1 << index
	b.bitMap[bucket] &= 0 << index
}

func (b *BitMap) IsExist(value uint64) bool {
	if value > b.size {
		return false
	}
	bucket := value / 8
	index := value % 8

	return (1 << index) == (b.bitMap[bucket] & (1 << index))
}
