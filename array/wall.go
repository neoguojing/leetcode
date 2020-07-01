package array

import (
	"leetcode/utils"
)

//MaxArea ...
// no 11
// 每个数组代表一个高度，选两个任意的柱子往里边倒水，能最多倒多少水。
//首位两个指针，遍历求最大
func MaxArea(height []int) int {
	if height == nil {
		return 0
	}

	if len(height) == 0 {
		return 0
	}

	maxArea := 0
	start := 0
	end := len(height) - 1

	for start < end {
		w := end - start
		h := utils.Min(height[start], height[end])

		tmpArea := h * w
		maxArea = utils.Max(tmpArea, maxArea)

		if height[start] < height[end] {
			start++
		} else {
			end--
		}

	}

	return maxArea
}
