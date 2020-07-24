package sort

import "sort"

// BucketSort ...
// 桶排序
// 最好O（n）
func BucketSort(in []int, bucketSize int) []int {
	if in == nil || len(in) == 0 {
		return in
	}

	max := in[0]
	min := in[0]
	for _, v := range in {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}
	//	求需要多少个桶
	bucketNum := (max - min) / bucketSize
	buckets := make([][]int, bucketNum)
	//		初始化桶
	for i := 0; i < len(buckets); i++ {
		buckets[i] = make([]int, 0)
	}
	//将数据插入桶
	for i := 0; i < len(in); i++ {
		index := (in[i] - min) / bucketSize
		buckets[index] = append(buckets[index], in[i])
	}

	ret := make([]int, 0)
	for i := 0; i < len(buckets); i++ {
		sort.Ints(buckets[i])
		for _, v := range buckets[i] {
			ret = append(ret, v)
		}
	}

	return ret
}
