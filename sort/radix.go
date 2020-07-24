package sort

// RadixSort ...
// 基数排序
// 时间复杂度：o(kn)
func RadixSort(in []int) []int {
	// 求最大位数
	d := maxBits(in)
	//用于求每位的值
	raidx := 1
	//计数器，存储对应桶的元素个数
	counter := make([]int, 10)
	// 中间结果存储
	tmp := make([]int, len(in))
	for i := 1; i <= d; i++ {
		// 计数器重置
		for j := range counter {
			counter[j] = 0
		}
		// 统计每个桶的数据个数
		for j := 0; j < len(in); j++ {
			k := (in[j] / raidx) % 10
			counter[k]++
		}
		// 把桶的计数，摊平到输入数据的空间
		// 假设：counter[0] = 0 counter[1] = 3 counter[2] = 1
		// 转换后：counter[0] = 0 counter[1] = 3 counter[2] = 4，因为每个通的值在tmp中是连续存储的，可以看出其实是求在tmp中的每个桶的右边界
		for j := 1; j < 10; j++ {
			counter[j] = counter[j-1] + counter[j]
		}
		//倒叙
		for j := len(in) - 1; j >= 0; j-- {
			k := (in[j] / raidx) % 10
			tmp[counter[k]-1] = in[j]
			counter[k]--
		}

		raidx = raidx * 10
	}

	return tmp
}

func maxBits(in []int) int {
	max := 0
	for _, v := range in {
		if v > max {
			max = v
		}
	}

	bitNum := 1
	p := 10
	for max > p {
		bitNum++
		max = max / 10
	}

	return bitNum
}
