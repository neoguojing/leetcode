package tree

type BITree []int

func getParent(i int) int {
	return i - i&(-i)
}

func getNext(i int) int {
	return i + i&(-i)
}

// size 为实际数组size
func NewBiTree(nums []int) BITree {
	bit := make(BITree, len(nums)+1)
	for i := range nums {
		bit.Add(i, nums[i])
	}
	return bit
}

// Add i为数组索引
func (bit BITree) Add(i int, val int) {
	i = i + 1
	for ; i < len(bit); i = getNext(i) {
		bit[i] = val
	}
}

//GetSum  0->i的和
func (bit BITree) GetSum(i int) int {
	i = i + 1
	sum := 0
	for i > 0 {
		sum += bit[i]
		i = getParent(i)
	}

	return sum
}

//RangeSum  i->j的和
func (bit BITree) RangeSum(i, j int) int {
	if i > j {
		return -1
	}

	i = i + 1
	sum := 0
	for i > 0 {
		sum += bit[i]
		i -= getNext(i)
	}

	return sum
}
