package combination

// Combine ...
// no 77
// 给定 n ，k ，表示从 { 1, 2, 3 ... n } 中选 k 个数，输出所有可能，并且选出数字从小到大排列，每个数字只能用一次。
func Combine(n, k int) [][]int {

	tmp := make([]int, 0)
	ret := make([][]int, 0)
	combine(n, k, 1, tmp, &ret)
	return ret
}

func combine(n int, k, idx int, tmp []int, ret *[][]int) {
	if k == len(tmp) {
		row := make([]int, k)
		copy(row, tmp)
		*ret = append(*ret, row)
		return
	}

	for i := idx; i <= n; i++ {
		tmp = append(tmp, i)
		combine(n, k, i+1, tmp, ret)
		tmp = tmp[:len(tmp)-1]
	}
}
