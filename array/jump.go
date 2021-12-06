package array

// CanJump ...
// no 55
//从数组的第 0 个位置开始跳，跳的距离小于等于数组上对应的数.判断是否能跳到最后一个位置
//  [2,3,1,1,4]
func CanJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	isUpdate := false
	positon := len(nums) - 1
	for positon != 0 {
		isUpdate = false
		for i := 0; i < positon; i++ {
			if nums[i] >= (positon - i) {
				positon = i
				isUpdate = true
				break
			}
		}
		//若无法向前走，则表示无法到达末尾
		if !isUpdate {
			return false
		}
	}

	return true
}

// 贪心算法
func CanJump2(nums []int) bool {
	far := 0

	for i := 0; i < len(nums); i++ {
		// 若最远距离比当前步数还小，表示无法继续跳
		if far < i {
			return false
		}
		// 更新能跳的最远距离
		if far < nums[i]+i {
			far = nums[i] + i
		}

	}

	return true
}

// 动态规划
func CanJump3(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		for step := 0; step <= nums[i] && (i+step) < len(nums); step++ {
			if dp[i+step] {
				dp[i] = true
				goto HEAR
			}
		}
	HEAR:
	}

	return dp[0]
}

// Jump ...
// no 45
// 从数组的第 0 个位置开始跳，跳的距离小于等于数组上对应的数。求出跳到最后个位置需要的最短步数
// 2，3，1，1，4
//1 贪心算法
// 2.从末尾开始，每次找能跳到当前位置的最远位置,直到到达0的位置
func Jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	positon := len(nums) - 1
	steps := 0
	for positon != 0 {
		for i := 0; i < positon; i++ {
			if nums[i] >= (positon - i) { //i 从0开始，保证每次找到了跳的最远的位置
				positon = i
				steps++
				break
			}
		}
	}

	return steps
}

// 贪心策略
func Jump2(nums []int) int {
	jump, curEnd, curFartest := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		// 当前能够跳的最远距离（每步的最优解）
		if curFartest < i+nums[i] {
			curFartest = i + nums[i]
		}
		// 一旦到了当前的end，则继续下一个目标
		if i == curEnd {
			jump++
			curEnd = curFartest
		}
	}

	return jump
}

// 动态规划
// dp[n-1] = max
// dp[i] = min(dp[i+k] +1)
func Jump3(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = len(nums)
	}
	dp[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		min := len(nums)
		for step := 0; step <= nums[i] && (i+step) < len(nums); step++ {
			if dp[i+step]+1 < min {
				min = dp[i+step] + 1
			}
		}
		dp[i] = min
	}

	return dp[0]
}

// no 1036
// no 1871
