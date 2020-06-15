package array

import (
	"sort"
)

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

//排序和双指针
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
