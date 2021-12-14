package stack

// LargestRectangleArea 求柱状表能够表示的最大面积
// no 84
//  heights = [2,1,5,6,2,3] 返回 2 * 5 = 10 在5和6的位置得到最大值
// 遇到比top大的，入栈，实际上是保存最大的右边比自己大的值
//
//
func LargestRectangleArea(heights []int) int {
	return 0
}

// dp[i] 表示每个位置能拿到的最大值；dp[i] = right-left -1
// left[i] 和right[i] 表示i的左边和右边大于i的最靠边的索引
// 会超时
func LargestRectangleAreaDP(heights []int) int {
	if heights == nil {
		return 0
	}
	left := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		for j := i; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			} else {
				left[i] = j
			}
		}
	}

	right := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			if heights[j] < heights[i] {
				break
			} else {
				right[i] = j
			}
		}
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		tmp := (right[i] - left[i] + 1) * heights[i]
		if tmp > maxArea {
			maxArea = tmp
		}

	}

	return maxArea
}

// 85
// no 496
// 503
// 316
// 402
// no 739
// 962
