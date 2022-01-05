package array

import (
	"sort"
)

//TwoSum ...
// no 1
func TwoSum(nums []int, target int) []int {
	anotherMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		anotherMap[nums[i]] = i
	}
	result := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		another := target - nums[i]
		if index, ok := anotherMap[another]; ok && index != i {
			result[0] = i
			result[1] = index

		}
	}

	return result
}

//ThreeSum ...
// no 15 排序和双指针
func ThreeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	if nums == nil || len(nums) == 0 {
		return result
	}

	if nums[len(nums)-1] < 0 || nums[0] > 0 {
		return result
	}

	for i := 0; i < len(nums)-2; i++ {
		another := target - nums[i]
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			m, n := i+1, len(nums)-1
			for m < n {
				tmp := nums[m] + nums[n]
				if tmp == another {
					result = append(result, []int{nums[i], nums[m], nums[n]})

					for m < n && nums[m] == nums[m+1] {
						m++
					}

					for m < n && nums[n] == nums[n-1] {
						n--
					}

					m++
					n--
				}

				if tmp < another {
					m++
				}

				if tmp > another {
					n--
				}
			}
		}
	}
	return result
}

// FourSumCount 四个数组相加，返回组合中和等于0 的组合个数；数组长度相等
// no 454
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumMap := map[int]int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			sumMap[sum] += 1
		}
	}

	res := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			sum := nums3[i] + nums4[j]
			if sum == 0 {
				res += sumMap[0]
			} else {
				res += sumMap[-sum]
			}
		}
	}

	return res
}
