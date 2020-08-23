package bits

// BitSet ...
type BitSet []int64

// NewBitSet ...
func NewBitSet(n int) BitSet {
	return make(BitSet, n/64+1)
}

// SetBit ...
func (b BitSet) SetBit(index int) {
	segIndex := index / 64
	posIndex := index % 64
	b[segIndex] |= 1 << posIndex
}

// UnSetBit ...
func (b BitSet) UnSetBit(index int) {
	segIndex := index / 64
	posIndex := index % 64
	b[segIndex] ^= 1 << posIndex
}

// IsBitSet ...
func (b BitSet) IsBitSet(index int) bool {
	segIndex := index / 64
	posIndex := index % 64
	seg := b[segIndex]

	return (seg | (1 << posIndex)) == seg
}
