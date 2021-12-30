package array

import (
	"math"
	"sort"
)

func iniDP(m, n int, tag int) [][]int {
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = tag
		}
	}

	return visited
}

// MajorityElement o(1) space
// 169 返回数组中出现次数超过一半的元素
func MajorityElement(nums []int) int {
	cntMap := map[int]int{}

	for i := range nums {
		if _, ok := cntMap[nums[i]]; !ok {
			cntMap[i] = 1
		} else {
			cntMap[i]++
		}
	}
	ret := 0
	for k, cnt := range cntMap {
		if cnt > len(cntMap)/2 {
			ret = k
		}
	}

	return ret
}

func MajorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 摩尔投票法
func MajorityElement3(nums []int) int {
	cadidate := int(math.MinInt32)
	count := 0

	for _, v := range nums {
		if count == 0 {
			cadidate = v
		}

		if v == cadidate {
			count += 1
		} else {
			count -= 1
		}
	}

	return cadidate
}

// MajorityElementII
// 229
func MajorityElementII(nums []int) []int {
	return nil
}

// ContainsDuplicate 数组中是否有重复值
// 217
// [1,2,3,1]
func ContainsDuplicate(nums []int) bool {
	set := map[int]bool{}
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = true
	}

	return false
}

// 219
// 220

// MoveZeroes
// 283 原地将0移动到末尾，并不改变元素相对顺序
// [0,1,0,3,12]
func MoveZeroes(nums []int) {
	start := 0
	for i := 0; i < len(nums); i++ {
		for ; start < len(nums) && nums[start] != 0; start++ {
		}

		if nums[i] != 0 && start < i {
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
}

func MoveZeroes2(nums []int) {
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[start], nums[i] = nums[i], nums[start]
			start++
		}
	}
}

// LengthOfLIS 求最长增长序列 不需要连续  要求o(nlogn)
// no 300
//  [0,1,0,3,2,3]
// 耐心排序
func LengthOfLIS(nums []int) int {

	piles := []int{}

	for _, num := range nums {
		isSet := false
		for i := range piles {
			if piles[i] >= num {
				piles[i] = num
				isSet = true
				break
			}
		}

		if !isSet {
			piles = append(piles, num)
		}
	}

	return len(piles)
}

// 动态规划
// dp[i] 表示以i为结尾的最长子串，初始全为1 1 dp[i] = dp[i-1] +1
func LengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	max := dp[0]
	for i := range dp {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
