package sort

import "sort"

/*
基数排序有两种方法：

这三种排序算法都利用了桶的概念，但对桶的使用方法上有明显差异：

基数排序：根据键值的每位数字来分配桶；
计数排序：每个桶只存储单一键值；
桶排序：每个桶存储一定范围的数值；
*/
// Swap ...
func Swap(in []int, x, y int) {
	tmp := in[x]
	in[x] = in[y]
	in[y] = tmp
}

//MergeSort ...
func MergeSort(in []int, start, end int) {

	if start < end {
		m := (end + start) / 2
		MergeSort(in, start, m)
		MergeSort(in, m+1, end)
		merge(in, start, m, end)
	}
}

// 使用index，模拟空间移动
func merge(in []int, start, mid, end int) {
	n := end - start + 1
	tmp := make([]int, n)

	i, j := start, mid+1
	k := 0
	for i <= mid && j <= end {
		if in[i] <= in[j] {
			tmp[k] = in[i]
			i++
		} else {
			tmp[k] = in[j]
			j++
		}
		k++
	}

	for j <= end {
		tmp[k] = in[j]
		j++
		k++
	}

	for i <= mid {
		tmp[k] = in[i]
		i++
		k++
	}
	// 此处赋值要加偏移
	for i := range tmp {
		in[i+start] = tmp[i]
	}
}

// Quicksort ...
func Quicksort(in []int, left, right int) {
	if left < right {
		pivot := partition(in, left, right)
		Quicksort(in, left, pivot-1)
		Quicksort(in, pivot+1, right)
	}
}

func partition(in []int, left, right int) int {

	pivot := (left + right) / 2
	tmp := in[pivot]

	// 将pivot移动到第一个，方便腾出空位用于元素交换
	in[pivot] = in[left]

	for left < right {
		for left < right && in[right] >= tmp {
			right--
		}

		in[left] = in[right]

		for left < right && in[left] <= tmp {
			left++
		}

		in[right] = in[left]

	}
	// pivot值归位
	in[left] = tmp
	// 真正的pivot
	return left
}

//HeapSort ...
func HeapSort(in []int) {
	MakeHeap(in)
	// 排序
	for i := len(in) - 1; i > 0; i-- {
		// 首尾交换
		Swap(in, 0, i)
		// 调整前i个
		AdjustHeap(in, 0, i)
	}
}

//MakeHeap ...
func MakeHeap(in []int) {
	// 构建队
	for i := len(in)/2 - 1; i >= 0; i-- {
		AdjustHeap(in, i, len(in))
	}
}

// AdjustHeap ...
// 调整第i个元素的位置,建立大顶堆
func AdjustHeap(in []int, i int, length int) {
	tmp := in[i]

	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && in[k] < in[k+1] {
			k++
		}

		if tmp < in[k] {
			in[i] = in[k]
			i = k
		} else {
			break
		}
	}

	in[i] = tmp
}

//BubbleSort ...
func BubbleSort(in []int) {
	if in == nil {
		return
	}
	n := len(in)
	if n < 2 {
		return
	}
	// 总共n-1轮
	for i := 0; i < n-1; i++ {
		// 每次需要计算新的j的终点
		flag := false
		for j := 0; j < n-1-i; j++ {
			if in[j] > in[j+1] {
				Swap(in, j, j+1)
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}

// SortColors
// no 75
//  数组中相同的元素排在一起 [2,0,2,1,1,0]
func SortColors(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	i, j := 0, len(nums)-1

	for p := i; p <= j; {
		if nums[p] == 2 {
			nums[p], nums[j] = nums[j], nums[p]
			j--
		}

		if nums[p] == 0 {
			nums[p], nums[i] = nums[i], nums[p]
			i++
			p++
		} else if nums[p] == 1 {
			p++
		}
	}
}

// no 324 小大小大排序 集合一定有解 o(n) 或者o(1)
// [1,1,1,1,4,5,6,7] => [1,4,1,5,1,6,1,7]
func WiggleSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	sort.Ints(nums)

	j := len(nums) - 1
	for i := 1; i < len(nums); i += 2 {
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}
}
