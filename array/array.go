package array

import "fmt"

/*RemoveDuplicates ...
 no 26
返回非重复数字的个数，并且把 nums 里重复的数字也去掉。
例如，nums = [ 1, 1, 2 ] ，那么就返回 2 ，并且把 nums 变成 [ 1, 2 ]。
双指针
 1 1 2 2 3 3
*/
func RemoveDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
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

//RemoveElement ...
// no 27
// 只不过这个是去除给定的值。
//遇到删除的元素，用末尾元素覆盖，长度减一，需要防止末尾元素即要删除的元素
func RemoveElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	end := len(nums)
	i := 0
	//end 会变化
	for i < end {
		if nums[i] == val {
			nums[i] = nums[end-1]
			end--
		} else {
			//防止末尾的数字和当前数字相同
			i++
		}
	}
	nums = nums[:end]
	fmt.Println(nums)
	return end
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

// Search
// no 704
func Search(nums []int, target int) int {
	return 0
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
			return mid
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

// 53 MaxSubArray
// 给一个数组，找出一个连续的子数组，长度任意，和最大
// todo

// 63 MaxSubArray
// 数组代表一个数字，[ 1, 2, 3 ] 就代表 123，然后给它加上 1，输出新的数组。数组每个位置只保存 1 位，也就是 0 到 9
// todo

// 75
// 给一个数组，含有的数只可能 0，1，2 中的一个，然后把这些数字从小到大排序。

// NumOfSubarrays 获取数组中所有子数组中和为奇数的子数组的个数 (子数组元素必须连续)
// no 1524
// dp[j]
func NumOfSubarrays(arr []int) int {
	var module int = 1e9 + 7
	odd, even, ret := 0, 0, 0

	for _, v := range arr {
		if v%2 == 0 {
			// 遇到偶数，则原先是偶数的子数组还是偶数
			even++
		} else {
			tmp := even
			even = odd
			odd = tmp
			odd++
		}

		ret = (ret%module + odd%module) % module
	}

	return ret
}
