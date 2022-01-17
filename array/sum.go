package array

import (
	"leetcode/utils"
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

// SubarraySum 和为k的子数组个数
// no 560
func SubarraySum(nums []int, k int) int {
	cnt := 0
	if len(nums) == 0 {
		return cnt
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}

			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}

func SubarraySumWithBitree(nums []int, k int) int {
	cnt := 0
	if len(nums) == 0 {
		return cnt
	}
	bit := utils.NewBiTree(nums)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := bit.RangeSum(i-1, j)
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}

func SubarraySumWithMap(nums []int, k int) int {
	cnt := 0
	sum := 0
	if len(nums) == 0 {
		return cnt
	}

	cntMap := map[int]int{}
	cntMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if n, ok := cntMap[sum-k]; ok {
			cnt += n
		}
		cntMap[sum] = 1
	}
	return cnt
}

// FindTargetSumWays,在数字前面放置+或者-号，让所有数字的和等于target；返回可能的组合个数
// no 494
func FindTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	cnt := 0
	findTargetSumWayshelper(nums, target, 0, 0, &cnt)
	return cnt
}

func findTargetSumWayshelper(nums []int, target, cur, sum int, cnt *int) {
	if cur == len(nums) {
		if sum == target {
			*cnt++
		}
		return
	}

	findTargetSumWayshelper(nums, target, cur+1, sum+nums[cur], cnt)
	findTargetSumWayshelper(nums, target, cur+1, sum-nums[cur], cnt)
}

//FindTargetSumWayWithsDP
func FindTargetSumWayWithsDP(nums []int, target int) int {

	return 0
}

// CanPartition
// 416 是否能够将集合分成两个组，每组的和相等？ 所有值都是正数
// 转换为任意个数字的和等于sum/2
// 动态规划
func CanPartition(nums []int) bool {

	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > sum/2 {
			return false
		}
		if nums[i] == sum/2 {
			return true
		}
	}

	return false
}
