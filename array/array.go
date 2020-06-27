package array

/*RemoveDuplicates ...
返回非重复数字的个数，并且把 nums 里重复的数字也去掉。
例如，nums = [ 1, 1, 2 ] ，那么就返回 2 ，并且把 nums 变成 [ 1, 2 ]。
双指针
 1 1 2 2 3 3
*/
func RemoveDuplicates(nums []int) int {
	//j 记录非重复元素的插入位置
	i, j := 1, 0
	for ; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	nums = nums[0 : j+1]
	return j + 1
}

/*SearchRange ...
no 34
找到目标值的第一次出现和最后一次出现的位置，同样要求 log ( n ) 下完成
数组已经有序
二分查找
查找左侧边界和右侧边界
*/
func SearchRange(nums []int, target int) []int {
	ret := make([]int, 2)
	left := FindLeftBound(nums, target)
	right := FindRightBound(nums, target)
	ret[0] = left
	ret[1] = right

	return ret
}

//FindLeftBound ...
//查找target的左侧边界
func FindLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	//寻找左侧边界
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}

	if left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}

//FindRightBound ...
//查找target的右侧边界
func FindRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	//寻找左侧边界
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}

//SearchInsert ...
//no 35
//给定一个有序数组，依旧是二分查找，不同之处是如果没有找到指定数字，需要返回这个数字应该插入的位置。
func SearchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] == target {
			return start
		}
	}

	if start == 0 {
		return 0
	}

	if end == len(nums)-1 {
		return len(nums)
	}

	return start
}
