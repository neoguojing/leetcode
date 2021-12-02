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
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			hi = mid - 1
		}
	}
	if lo > len(nums)-1 || nums[lo] != target {
		return -1
	}
	return lo
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
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}

	if lo == 0 {
		return 0
	}

	if hi == len(nums)-1 {
		return len(nums)
	}

	return lo
}

// SearchInRotateArray
// no 33
func SearchInRotateArraySimple(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			hi = mid
		} else {
			hi--
		}
	}

	rot := lo //等价于一个偏移，即旋转的次数
	lo, hi = 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		realMid := (mid + rot) % len(nums)
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}

func SearchInRotateArray(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	lo, hi := 0, len(nums)-1

	return searchInRotate(nums, lo, hi, target)
}

func searchInRotate(nums []int, lo, hi, target int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2

	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		if nums[mid] >= nums[lo] {
			ret := binarySearch(nums, lo, mid-1, target)
			if ret != -1 {
				return ret
			} else {
				return searchInRotate(nums, mid+1, hi, target)
			}
		} else {
			return searchInRotate(nums, lo, mid-1, target)
		}
	} else {
		if nums[mid] <= nums[hi] {
			ret := binarySearch(nums, mid+1, hi, target)
			if ret != -1 {
				return ret
			} else {
				return searchInRotate(nums, lo, mid-1, target)
			}
		} else {
			return searchInRotate(nums, mid+1, hi, target)
		}
	}

}

func binarySearch(nums []int, lo, hi, target int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			hi = mid - 1
		}
	}

	if lo >= len(nums) || nums[lo] != target {
		return -1
	}

	return lo
}

//SearchInRotateArray2 num中包含重复元素
// no 81
/*
[1,0,1,1,1]
0
*/
func SearchInRotateArrayWithDup(nums []int, target int) bool {
	if nums == nil {
		return false
	}

	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] == nums[lo] && nums[mid] == nums[hi] {
			hi--
			lo++
		} else if nums[mid] <= nums[hi] {
			// mid和target相比较的两种情况
			if nums[mid] < target && nums[hi] >= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			if nums[mid] > target && nums[lo] <= target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}

	}

	return false
}
func SearchInRotateArray2(nums []int, target int) bool {
	if nums == nil {
		return false
	}

	lo, hi := 0, len(nums)-1

	return searchInRotate2(nums, lo, hi, target) != -1
}

func searchInRotate2(nums []int, lo, hi, target int) int {
	if lo > hi {
		return -1
	}

	mid := lo + (hi-lo)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] == nums[lo] && nums[mid] == nums[hi] {
		ret := searchInRotate2(nums, lo, mid-1, target)
		if ret != -1 {
			return ret
		} else {
			return searchInRotate2(nums, mid+1, hi, target)
		}
	}

	if nums[mid] > target {
		if nums[mid] == nums[hi] && nums[mid] > nums[lo] {
			return searchInRotate2(nums, lo, mid-1, target)
		}

		if nums[mid] >= nums[lo] {
			ret := binarySearch(nums, lo, mid-1, target)
			if ret != -1 {
				return ret
			} else {
				return searchInRotate2(nums, mid+1, hi, target)
			}
		} else if nums[mid] < nums[lo] {
			return searchInRotate2(nums, lo, mid-1, target)
		}
	} else {

		if nums[mid] == nums[lo] && nums[mid] < nums[hi] {
			return searchInRotate2(nums, mid+1, hi, target)
		}

		if nums[mid] <= nums[hi] {
			ret := binarySearch(nums, mid+1, hi, target)
			if ret != -1 {
				return ret
			} else {
				return searchInRotate2(nums, lo, mid-1, target)
			}
		} else if nums[mid] > nums[hi] {
			return searchInRotate2(nums, mid+1, hi, target)
		}
	}

	return -1
}

// FindMin
// no 153
func FindMin2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	lo, hi := 0, len(nums)-1
	mid := 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] >= nums[lo] && nums[mid] <= nums[hi] {
			return nums[lo]
		}

		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			hi = mid
		} else {

		}
	}

	return nums[lo]
}

func FindMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	lo, hi := 0, len(nums)-1
	mid := lo + (hi-lo)/2
	return searchMinInRotate(nums, lo, hi, mid)
}

func searchMinInRotate(nums []int, lo, hi, mid int) int {
	if lo == hi {
		return nums[lo]
	}

	if nums[mid] >= nums[lo] && nums[mid] <= nums[hi] {
		return nums[lo]
	}

	if nums[mid] >= nums[lo] && nums[mid] > nums[hi] {
		lo = mid + 1
		mid := lo + (hi-lo)/2
		return searchMinInRotate(nums, lo, hi, mid)
	}

	if nums[mid] <= nums[lo] && nums[mid] <= nums[hi] {
		hi = mid
		mid := lo + (hi-lo)/2
		return searchMinInRotate(nums, lo, hi, mid)
	}

	return nums[lo]
}

// 154
func FindMinWithDup(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}
	lo, hi, mid := 0, len(nums)-1, 0

	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else if nums[mid] < nums[hi] {
			hi = mid
		} else {
			hi--
		}
	}

	return nums[lo]
}

// 322
// 1329
// 1798
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
