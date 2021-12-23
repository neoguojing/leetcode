package array

import "sort"

func iniDP(m, n int, tag int) [][]int {
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = tag
		}
	}

	return visited
}

// MajorityElement o(1) space
// 169 返回数组中出现次数超过一半的元素
func MajorityElement(nums []int) int {
	cntMap := map[int]int{}

	for i := range nums {
		if _, ok := cntMap[nums[i]]; !ok {
			cntMap[i] = 1
		} else {
			cntMap[i]++
		}
	}
	ret := 0
	for k, cnt := range cntMap {
		if cnt > len(cntMap)/2 {
			ret = k
		}
	}

	return ret
}

func MajorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
