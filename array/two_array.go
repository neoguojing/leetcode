package array

import (
	"sort"
)

//FindMedianSortedArrays ...
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
