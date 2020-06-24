package combination

import "fmt"

/*Subsets ...
no 78
给一个数组，输出这个数组的所有子数组
1.回溯法：取一个数，有两种情况，该数在集合和不在集合中;set[n+1] = set[n] + 0|1
2.位表示法：
*/
func Subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	result := make([][]int, 0)
	tmp := make([]int, 0)
	subsets(nums, 0, tmp, result)
	return result
}

func subsets(nums []int, index int, tmp []int, result [][]int) [][]int {
	result = append(result, tmp)
	for i := index; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		result = subsets(nums, i+1, tmp, result)
		tmpLen := len(tmp)
		if tmpLen == 0 {
			tmp = make([]int, 0)
		} else if tmpLen == 1 {
			tmp = make([]int, 0)
		} else {
			tmp = tmp[0 : tmpLen-1]
		}

	}

	fmt.Println(result)
	return result
}
