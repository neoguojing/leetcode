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
		bit[i] += val
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

//RangeSum  i+1->j的和
func (bit BITree) RangeSum(i, j int) int {
	if i > j {
		return -1
	}
	sum := 0
	i = i + 1
	j = j + 1

	for ; j > i; j = getParent(j) {
		sum += bit[j]
	}
	for ; i > j; i = getParent(i) {
		sum -= bit[i]
	}

	return sum
}

func (bit BITree) Get(i int) int {
	return bit.RangeSum(i-1, i)
}

func (bit BITree) Set(i, val int) {
	bit.Add(i, val-bit.Get(i))
}

func (bit BITree) GetArray() []int {
	ret := make([]int, len(bit)-1)
	for i := len(bit) - 1; i > 0; i-- {
		j := getNext(i)
		if j < len(bit) {
			ret[j-1] = bit[j] - bit[i]
		}

	}
	return ret
}
