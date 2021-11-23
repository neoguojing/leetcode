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

// 977
