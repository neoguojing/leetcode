package combination

import "sort"

//Permute ...
// no 46
//就是给定几个数，然后输出他们所有排列的可能。数字不重复
func Permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	ret := make([][]int, 0)
	tmp := make([]int, 0)
	permute(nums, tmp, &ret)
	return ret
}

func permute(nums []int, tmp []int, ret *[][]int) {
	if len(tmp) == len(nums) {
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contain(tmp, nums[i]) {
			continue
		}

		tmp = append(tmp, nums[i])
		permute(nums, tmp, ret)
		tmp = tmp[:len(tmp)-1]
	}
}

func contain(in []int, target int) bool {
	for _, v := range in {
		if target == v {
			return true
		}
	}

	return false
}

//PermuteUnique ...
// no 47
//就是给定几个数，然后输出他们所有排列的可能。有重复数字
// 排序先
func PermuteUnique(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	sort.Ints(nums)

	ret := make([][]int, 0)
	tmp := make([]int, 0)
	permuteUnique(nums, tmp, &ret)
	return ret
}

func permuteUnique(nums []int, tmp []int, ret *[][]int) {
	if len(tmp) == len(nums) {
		out := make([]int, len(tmp))
		for i, v := range tmp {
			out[i] = nums[v]
		}
		*ret = append(*ret, out)
		return
	}

	for i := 0; i < len(nums); i++ {
		//判断上一个元素和当前相等，且上一个元素不在临时路径中，则跳过，
		if i > 0 && !contain(tmp, i-1) && nums[i] == nums[i-1] {
			continue
		}
		// 相同下表的元素跳过
		if contain(tmp, i) {
			continue
		}
		// tmp 保存元素的下标
		tmp = append(tmp, i)
		permuteUnique(nums, tmp, ret)
		tmp = tmp[:len(tmp)-1]
	}
}

// no 60
// 是给一个 n，不是输出它的全排列，而是把所有组合从从小到大排列后，输出第 k 个。
