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
