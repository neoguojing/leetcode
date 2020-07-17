package array

// CanJump ...
// no 55
//从数组的第 0 个位置开始跳，跳的距离小于等于数组上对应的数.判断是否能跳到最后一个位置
//
func CanJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
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

// Jump ...
// no 45
// 从数组的第 0 个位置开始跳，跳的距离小于等于数组上对应的数。求出跳到最后个位置需要的最短步数
// 2，3，1，1，4
//1 贪心算法
// 2.从末尾开始，每次找能跳到当前位置的最远位置,直到到达0的位置
func Jump(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	positon := len(nums) - 1
	steps := 0
	for positon != 0 {
		for i := 0; i < positon; i++ {
			if nums[i] >= (positon - i) {
				positon = i
				steps++
				break
			}
		}
	}

	return steps
}
