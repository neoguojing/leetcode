package bits

// BloomFilter ...
type BloomFilter struct {
	hashNum int
	b       BitSet
}

type hashFunc func(data []byte) int

var hashs []hashFunc

// NewBloomFilter ...
func NewBloomFilter(n int, k int) *BloomFilter {
	return &BloomFilter{
		hashNum: k,
		b:       NewBitSet(n),
	}
}

// Put 添加一条记录
func (filter *BloomFilter) Put(data []byte) {
	l := len(filter.b)
	for i := 0; i < filter.hashNum; i++ {
		filter.b.SetBit(hashs[i](data) % l)
	}
}

// Has 推测记录是否已存在
func (filter *BloomFilter) Has(data []byte) bool {
	l := len(filter.b)

	for i := 0; i < filter.hashNum; i++ {
		if !filter.b.IsBitSet(hashs[i](data) % l) {
			return false
		}
	}

	return true
}
