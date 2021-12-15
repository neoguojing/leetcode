package stack

import "fmt"

// LargestRectangleArea 求柱状表能够表示的最大面积
// no 84
//  heights = [2,1,5,6,2,3] 返回 2 * 5 = 10 在5和6的位置得到最大值
func LargestRectangleArea(heights []int) int {

	stack := make(Stack, 0)
	heights = append(heights, 0)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 && heights[i] < heights[stack.Top()] {
			cur := stack.Pop()
			width := i - 1 - stack.Top() //是i-1当前元素的右边界（因为小与当前元素），stack.Top()是左边界（栈递增，所以是top小于当前元素）
			if len(stack) == 0 {         // 栈为空，则元素没有左边界或者 实际等价于左边界是-1  i-1-（-1），
				width = i
			}
			area := width * heights[cur]
			if area > maxArea {
				maxArea = area
			}
		}

		stack.Push(i)
	}
	return maxArea
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

// Trap
// no 42 栈解法
// [0,1,0,2,1,0,1,3,2,1,2,1]
// 能够接住的最多的雨水
// 需要找到当前柱子i,左右两边大于i对应值的最近的值
// 遇到大的值出栈，所以是一个递减的栈
// 左右接不了水，所以不考虑
func Trap(height []int) int {
	stack := make(Stack, 0)
	total := 0
	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack.Top()] {
			cur := stack.Pop()
			if len(stack) == 0 {
				break
			}
			smaller := height[i]
			if height[i] > height[stack.Top()] {
				smaller = height[stack.Top()]
			}
			total += (smaller - height[cur]) * (i - 1 - stack.Top())
		}
		stack.Push(i)
	}

	return total
}

// 85
// NextGreaterElement 暴力法
// no 496
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	for i, v := range nums1 {
		foundV := false
		j := 0
		for ; j < len(nums2); j++ {
			if nums2[j] == v {
				foundV = true
			}

			if foundV && nums2[j] > v {
				nums1[i] = nums2[j]
				break
			}
		}

		if j == len(nums2) {
			nums1[i] = -1
		}
	}

	return nums1
}

func NextGreaterElementStack(nums1 []int, nums2 []int) []int {

	stack := make(Stack, 0)
	nextMap := map[int]int{}

	for i := 0; i < len(nums2); i++ {
		for len(stack) != 0 && nums2[i] > stack.Top() {
			nextMap[stack.Pop()] = nums2[i]
		}
		stack.Push(nums2[i])
	}

	fmt.Println(nextMap)

	for i := 0; i < len(nums1); i++ {
		if v, ok := nextMap[nums1[i]]; !ok {
			nums1[i] = -1
		} else {
			nums1[i] = v
		}
	}

	return nums1
}

// NextGreaterElementsWithCircle 同上，但是nums是可以从末尾跳到第一个继续遍历的
// 503
func NextGreaterElementsWithCircle(nums []int) []int {
	if nums == nil {
		return nums
	}

	stack := make(Stack, 0)
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	for i := 0; i < len(nums)*2; i++ {
		for len(stack) != 0 && nums[stack.Top()] < nums[i%len(nums)] {
			res[stack.Pop()] = nums[i%len(nums)]
		}
		stack.Push(i % len(nums))
	}

	return res
}

// 316
// 402
// no 739
// 962
