package sort

/*
基数排序有两种方法：

这三种排序算法都利用了桶的概念，但对桶的使用方法上有明显差异：

基数排序：根据键值的每位数字来分配桶；
计数排序：每个桶只存储单一键值；
桶排序：每个桶存储一定范围的数值；
*/

//MergeSort ...
func MergeSort(in []int, start, end int) {

	if start < end {
		m := (end + start) / 2
		MergeSort(in, start, m)
		MergeSort(in, m+1, end)
		merge(in, start, m, end)
	}
}

// 使用index，模拟空间yidong
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
