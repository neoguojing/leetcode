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

// Trap ...
// no 42
// 给定一个数组，每个数代表从左到右墙的高度，求出能装多少单位的水。也就是图中蓝色正方形的个数。
//按列求，找出左边最高的和右边最高的列
// 1. 最两端的列不用考虑，因为一定不会有水。所以下标从 1 到 length - 2
// 2.较矮的墙的高度大于当前列的墙的高度，当前列可以蓄水
// 3.较矮的墙的高度小于当前列的墙的高度，当前列不会有水
// 4.较矮的墙的高度等于当前列的墙的高度，当前列不会有水
// 如下动态规划
func Trap(height []int) int {
	maxLeft := make([]int, len(height))
	for i := 1; i < len(height)-1; i++ {
		maxLeft[i] = utils.Max(maxLeft[i-1], height[i-1])
	}
	maxRight := make([]int, len(height))
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = utils.Max(maxRight[i+1], height[i+1])
	}
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		tmp := utils.Min(maxLeft[i], maxRight[i])
		if tmp > height[i] {
			sum += tmp - height[i]
		}
	}
	return sum
}
