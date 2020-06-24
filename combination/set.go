package combination

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
	//result = append(result, make([]int, 0))
	for i := rang nums{
		result = append(result,subsets(nums,i)) 
	}
	return result
}

func subsets(nums []int,index int) [][]int {
	if len(nums) == 0 {
		
		return 
	}
	return result
}