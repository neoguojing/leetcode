package combination

import (
	"fmt"
	"sort"
)

//NextPermutation ...
// no 31
// 给定一个数，然后将这些数字的位置重新排列，得到一个刚好比原数字大的一种排列。如果没有比原数字大的，就升序输出
func NextPermutation(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}

	pos := 0
	//查找递减的第一个元素
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pos = i - 1
			break
		}
	}

	if pos == 0 {
		sort.Ints(nums)
		return nums
	}

	justIndex := 0
	posVal := nums[pos]

	remain := nums[pos+1:]
	sort.Ints(remain)
	for i := 0; i < len(remain)-1; i++ {
		if remain[i] > posVal {
			justIndex = i
			break
		}
	}
	fmt.Println(pos)

	nums[pos] = remain[justIndex]
	remain[justIndex] = posVal
	nums = append(nums[:pos+1], remain...)
	return nums
}
