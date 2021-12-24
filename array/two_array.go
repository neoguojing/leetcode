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

// Intersect 求两个数组的交集
// 350
func Intersect(nums1 []int, nums2 []int) []int {
	a, b := len(nums1), len(nums2)
	num1Map := map[int]int{}
	for i := range nums1 {
		num1Map[nums1[i]]++
	}
	num2Map := map[int]int{}
	for i := range nums2 {
		num2Map[nums2[i]]++

	}

	ret := []int{}
	if a <= b {
		for _, k := range nums1 {
			if _, ok := num2Map[k]; ok {
				ret = append(ret, k)
				num2Map[k]--
				if num2Map[k] == 0 {
					delete(num2Map, k)
				}

			}
		}
	} else if a > b {
		for _, k := range nums2 {
			if _, ok := num1Map[k]; ok {
				ret = append(ret, k)
				num1Map[k]--
				if num1Map[k] == 0 {
					delete(num1Map, k)
				}
			}
		}
	}

	return ret
}
