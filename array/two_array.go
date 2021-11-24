package array

import (
	"sort"
)

//FindMedianSortedArrays ...
// no 4
//已知两个有序数组，找到两个数组合并后的中位数
//
func FindMedianSortedArrays(nums1, nums2 []int) float32 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	mid := len(nums1) / 2
	if len(nums1)%2 == 0 {
		return float32((nums1[mid] + nums1[mid-1])) / 2.0
	}

	return float32(nums1[mid])
}

// Merge 将两个排序的数组合并,排序结果存储在nums1中
// no 88
func Merge(nums1 []int, m int, nums2 []int, n int) {
	j := n - 1
	k := m + n - 1
	i := m - 1
	for ; k >= 0 && i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	for ; i >= 0; i-- {
		nums1[k] = nums1[i]
		k--
	}

	for ; j >= 0; j-- {
		nums1[k] = nums2[j]
		k--
	}

}

// SortedSquares ...
// no 977
func SortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	m := 0
	for ; m < len(nums) && nums[m] < 0; m++ {
	}
	pos := nums[m:]
	negs := nums[:m]
	ret := make([]int, len(nums))
	i := 0
	j := len(negs) - 1
	k := 0
	for ; k < len(nums) && i < len(pos) && j >= 0; k++ {

		if pos[i]*pos[i] < negs[j]*negs[j] {
			ret[k] = pos[i] * pos[i]
			i++
		} else {
			ret[k] = negs[j] * negs[j]
			j--
		}
	}

	for ; i < len(pos); i++ {
		ret[k] = pos[i] * pos[i]
		k++
	}

	for ; j >= 0; j-- {
		ret[k] = negs[j] * negs[j]
		k++
	}

	return ret
}
