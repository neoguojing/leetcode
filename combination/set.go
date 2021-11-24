package combination

import (
	"fmt"
	"sort"
)

/*Subsets ...
no 78
给一个数组，输出这个数组的所有子数组
1.回溯法：递归广度由选择的个数决定(数组长度),递归深度由子问题个数决定,子问题即顺序子数组的个数
取一个数，有两种情况，该数在集合和不在集合中
主函数调用递归函数
递归函数是一个多路递归
需要维护一个当前路径tmp，记录当前取值情况，方便回溯
需要维护一个变量记录当前迭代的可选路径中的某个值
1 2 3
 1  2  3
 1  2
2.位表示法：
*/
func Subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	result := make([][]int, 0)
	tmp := make([]int, 0)
	subsets(nums, 0, tmp, &result)
	return result
}

func subsets(nums []int, index int, tmp []int, result *[][]int) {
	//copy 防止内存被修改
	item := make([]int, len(tmp))
	copy(item, tmp)

	fmt.Println(item)
	fmt.Println(index)

	*result = append(*result, item)
	for i := index; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		subsets(nums, i+1, tmp, result)
		tmp = tmp[0 : len(tmp)-1]
	}
}

//SubsetsWithBit ...
func SubsetsWithBit(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	bitNum := len(nums)
	result := make([][]int, 0)
	result = append(result, make([]int, 0))
	//总共2的n次方种排列
	for i := 1; i < 1<<bitNum; i++ {
		item := make([]int, 0)
		for j, k := bitNum-1, 1; j >= 0 && k <= i; j-- {
			//利用&操作确定某位是否在其中
			if k&i != 0 {
				item = append(item, nums[j])
			}
			k = k << 1
		}
		result = append(result, item)
	}
	return result
}

//SubsetsWithDup ...
// no 90
// 去除重复元素 1.数组排序 2.
/*1 2 2
 */
func SubsetsWithDup(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	//对数组排序
	sort.Ints(nums)

	result := make([][]int, 0)
	tmp := make([]int, 0)
	subsetsWithDup(nums, 0, tmp, &result)
	return result
}

func subsetsWithDup(nums []int, index int, tmp []int, result *[][]int) {
	row := make([]int, len(tmp))
	copy(row, tmp)
	*result = append(*result, row)

	for i := index; i < len(nums); i++ {
		//当前数等于前一个数，且前一个数在该次迭代中存在
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		subsetsWithDup(nums, i+1, tmp, result)
		tmp = tmp[0 : len(tmp)-1]
	}
}

//IntervalIntersection  求两个排序集合的交集
// no 986
func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}

	ret := make([][]int, 0)
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		if firstList[i][1] < secondList[j][0] {
			i++
		} else if firstList[i][0] > secondList[j][1] {
			j++
		} else {
			l := firstList[i][0]
			if firstList[i][0] < secondList[j][0] {
				l = secondList[j][0]
			}
			r := firstList[i][1]
			if firstList[i][1] > secondList[j][1] {
				r = secondList[j][1]
				j++
			} else {
				i++
			}
			new := []int{l, r}
			ret = append(ret, new)
		}
	}

	return ret
}

// 数组集合取并集
// no 56
func Union(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	ret := make([][]int, 0)
	doUnion := false
	i := 1
	for ; i < len(intervals); i++ {
		if intervals[i-1][1] >= intervals[i][0] && intervals[i-1][0] <= intervals[i][1] {
			l := intervals[i-1][0]
			if l > intervals[i][0] {
				l = intervals[i][0]
			}
			r := intervals[i-1][1]
			if r < intervals[i][1] {
				r = intervals[i][1]
			}
			union := []int{l, r}
			ret = append(ret, union)
			doUnion = true
		} else {
			ret = append(ret, intervals[i-1])
		}
	}
	ret = append(ret, intervals[i-1])
	fmt.Println(ret)
	if !doUnion {
		return ret
	}

	return Union(ret)
}
