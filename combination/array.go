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
	//从后向前查找递减的第一个元素
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pos = i - 1
			break
		}
	}
	//找不到则返回升序序列
	if pos == 0 {
		sort.Ints(nums)
		return nums
	}

	justIndex := 0
	posVal := nums[pos]

	remain := nums[pos+1:]
	sort.Ints(remain)
	//在递减位置开始找到第一个比它大的元素
	for i := 0; i < len(remain)-1; i++ {
		if remain[i] > posVal {
			justIndex = i
			break
		}
	}
	fmt.Println(pos)
	//交换元素位置
	nums[pos] = remain[justIndex]
	remain[justIndex] = posVal
	nums = append(nums[:pos+1], remain...)
	return nums
}

//CombinationSum ...
// no 39
// 给几个数字,无重复，一个目标值，输出所有和等于目标值的组合,同一个数字可以被选择多次,且集合不能重复,
//回溯法：用减法，当剩余值小于0时，该路径走不通，等于时路径有效；需要记录当前路径（数组），以及当前路径的选择
//每个迭代选择集是所有的数字
// 2 3 6 7  7
func CombinationSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	tmp := make([]int, 0)
	result := make([][]int, 0)
	combinationSum(nums, target, tmp, &result, 0)
	return result

}

func combinationSum(nums []int, target int, tmp []int, result *[][]int, pos int) {
	//递归结束条件
	if target == 0 {
		row := make([]int, len(tmp))
		copy(row, tmp)
		*result = append(*result, row)
		return
	} else if target < 0 {
		return
	}
	//多路递归
	//i = pos 表明只能向后选择，避免向前选择导致的重复
	for i := pos; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		//此处pos 传入i，而不是传入i+1,表面一个数可以重复选择
		combinationSum(nums, target-nums[i], tmp, result, i)
		tmp = tmp[:len(tmp)-1]
	}

}

//CombinationSum2 ...
// no 40
// 给几个数字有重复，一个目标值，输出所有和等于目标值的组合,同一个数字只能被使用一次,且集合不能重复,
//回溯法：
func CombinationSum2(nums []int, target int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	tmp := make([]int, 0)
	result := make([][]int, 0)
	combinationSum2(nums, target, tmp, &result, 0)
	return result

}

func combinationSum2(nums []int, target int, tmp []int, result *[][]int, pos int) {
	//递归结束条件
	if target == 0 {
		row := make([]int, len(tmp))
		copy(row, tmp)
		*result = append(*result, row)
		return
	} else if target < 0 {
		return
	}
	//多路递归
	//i = pos 表明只能向后选择，避免向前选择导致的重复
	for i := pos; i < len(nums); i++ {
		//有重复元素的需要这个条件
		if i > pos && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		combinationSum2(nums, target-nums[i], tmp, result, i+1)
		tmp = tmp[:len(tmp)-1]
	}

}
