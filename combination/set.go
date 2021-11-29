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

// Union 数组集合取并集 集合无序
// no 56
func Union(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 冒泡排序，安装第一个坐标排序
	for i := 0; i < len(intervals)-1; i++ {
		hasSwap := false
		for j := len(intervals) - 1; j > i; j-- {
			if intervals[j][0] < intervals[j-1][0] {
				tmp := intervals[j]
				intervals[j] = intervals[j-1]
				intervals[j-1] = tmp
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}
	ret := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		union := HasUnion(ret[len(ret)-1], intervals[i])
		if len(union) == 0 {
			ret = append(ret, intervals[i])
		} else {
			ret[len(ret)-1] = union
		}
	}
	return ret
}

func HasUnion(a, b []int) []int {
	if a[1] < b[0] || a[0] > b[1] {
		return nil
	}

	l := a[0]
	if l > b[0] {
		l = b[0]
	}
	r := a[1]
	if r < b[1] {
		r = b[1]
	}
	return []int{l, r}
}

// Insert
// no 57
func Insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	start := -1
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] < newInterval[0] {
			start = i
			continue
		}

		break
	}
	fmt.Println(start)
	ret := make([][]int, 0)
	rear := make([][]int, 0)
	if start == -1 {
		union := HasUnion(newInterval, intervals[0])
		if len(union) == 0 {
			ret = append(ret, newInterval)
			ret = append(ret, intervals...)
			return ret
		} else {
			ret = append(ret, union)
			rear = append(rear, intervals[1:]...)
		}
	} else {
		ret = append(ret, intervals[:start+1]...)
		rear = append(rear, newInterval)
		rear = append(rear, intervals[start+1:]...)
	}

	fmt.Println(ret, rear)

	for i := 0; i < len(rear); i++ {
		union := HasUnion(ret[len(ret)-1], rear[i])
		if len(union) == 0 {
			ret = append(ret, rear[i])
		} else {
			ret[len(ret)-1] = union
		}
	}

	return ret
}

// FindPoisonedDuration 反复放毒，求敌人的总中毒时间
// no 495
func FindPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	sets := make([][]int, 1)
	sets[0] = []int{timeSeries[0], timeSeries[0] + duration}

	for i := 1; i < len(timeSeries); i++ {
		union := HasUnion(sets[len(sets)-1], []int{timeSeries[i], timeSeries[i] + duration})
		if union == nil {
			sets = append(sets, []int{timeSeries[i], timeSeries[i] + duration})
		} else {
			sets[len(sets)-1] = union
		}
	}
	ret := 0
	for _, v := range sets {
		ret += v[1] - v[0]
	}

	return ret
}

// no 715
type RangeModule struct {
	data [][]int
}

func Constructor() RangeModule {
	return RangeModule{
		data: make([][]int, 0),
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	fmt.Println("add--", []int{left, right})
	defer this.Print()
	set := []int{left, right}
	if len(this.data) == 0 {
		this.data = append(this.data, set)
		return
	}

	rear := make([][]int, 0)
	if right < this.data[0][0] {
		rear = this.data
		this.data = append([][]int{}, set)
		this.data = append(this.data, rear...)
		return
	}

	if left > this.data[len(this.data)-1][1] {
		this.data = append(this.data, set)
	}

	start := -1
	for i := 0; i < len(this.data); i++ {
		if this.data[i][0] < set[0] {
			start = i
			continue
		}
		break
	}

	if start == -1 {
		union := this.hasUnionSection(set, this.data[0])
		rear = append(rear, this.data[1:]...)
		this.data = [][]int{union}

	} else {
		rear = append(rear, set)
		rear = append(rear, this.data[start+1:]...)
		this.data = append([][]int{}, this.data[:start+1]...)
	}

	for i := 0; i < len(rear); i++ {
		union := this.hasUnionSection(this.data[len(this.data)-1], rear[i])
		if len(union) == 0 {
			this.data = append(this.data, rear[i])
		} else {
			this.data[len(this.data)-1] = union
		}
	}

}

func (this *RangeModule) QueryRange(left int, right int) bool {
	fmt.Println("query--", []int{left, right})
	if len(this.data) == 0 {
		return false
	}
	// set := []int{left, right}
	i := 0
	for ; i < len(this.data); i++ {
		if left >= this.data[i][1] {
			continue
		}
		break
	}
	if i == 0 {
		if left >= this.data[i][0] && right <= this.data[i][1] {
			return true
		}
	} else if i == len(this.data) {
		return false
	} else {
		if left >= this.data[i][0] && right <= this.data[i][1] {
			return true
		}
	}

	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	fmt.Println("remove--", []int{left, right})
	if len(this.data) == 0 {
		return
	}
	defer this.Print()
	if left >= this.data[len(this.data)-1][1] {
		return
	}
	if right < this.data[0][0] {
		return
	}

	i := 0
	for ; i < len(this.data); i++ {
		if left >= this.data[i][1] {
			continue
		}
		break
	}
	fmt.Println(i)
	count := 0
	end := 0
	for j := i; j < len(this.data); j++ {
		// 完全删除的情况
		set, ok := this.calDivSection(this.data[j], []int{left, right})
		if ok {
			if len(set) == 0 {
				count++
				end = j
			} else if len(set) == 1 {
				this.data[j] = set[0]
			} else { //只有一种情况会裂变，只循环一次，所以修改数组没问题
				rear := set
				rear = append(rear, this.data[j+1:]...)
				this.data = append(this.data[:j], rear...)
			}
		} else {
			break
		}
	}

	start := end - count + 1
	if end == len(this.data)-1 {
		this.data = this.data[:start]
	} else {
		this.data = append(this.data[:start], this.data[end+1:]...)
	}

}

func (this *RangeModule) hasUnionSection(a []int, b []int) []int {
	if a[1] < b[0] || a[0] > b[1] {
		return nil
	}

	l := a[0]
	if l > b[0] {
		l = b[0]
	}
	r := a[1]
	if r < b[1] {
		r = b[1]
	}
	return []int{l, r}
}

func (this *RangeModule) hasInterSection(a []int, b []int) []int {
	if a[1] < b[0] || a[0] >= b[1] {
		return nil
	}

	l := a[0]
	if l < b[0] {
		l = b[0]
	}
	r := a[1]
	if r > b[1] {
		r = b[1]
	}
	return []int{l, r}
}

func (this *RangeModule) calDivSection(target []int, div []int) ([][]int, bool) {
	if div[0] >= target[1] || div[1] <= target[0] {
		return nil, false
	}

	ret := make([][]int, 0)
	if div[0] <= target[0] && div[1] >= target[1] {
		return ret, true
	} else if div[0] > target[0] && div[1] >= target[1] {
		tmp := []int{target[0], div[0]}
		ret = append(ret, tmp)
	} else if div[0] < target[0] && div[1] < target[1] {
		tmp := []int{div[1], target[1]}
		ret = append(ret, tmp)
	} else {
		ret = append(ret, []int{target[0], div[0]})
		ret = append(ret, []int{div[1], target[1]})
	}

	return ret, true
}

func (this *RangeModule) Print() {
	fmt.Println("result--", this.data)
}

// CanPlaceFlowers 种花，相邻的位置不能种花，传入n，问是否可以全部种植
// 605
func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 0 {
		return false
	}

	if n > (len(flowerbed)/2 + len(flowerbed)%2) {
		return false
	}

	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}

	return count >= n

}

// MaxNumberOfFamilies 电影院座位分配
// 1386
func MaxNumberOfFamilies(n int, reservedSeats [][]int) int {
	if reservedSeats == nil {
		return n * 2
	}
	count := n * 2

	setMap := map[int][]byte{}
	for j := 0; j < len(reservedSeats); j++ {
		_, ok := setMap[reservedSeats[j][0]]
		if !ok {
			setMap[reservedSeats[j][0]] = []byte{0, 0, 0, 0}
		}

		if reservedSeats[j][1] >= 2 && reservedSeats[j][1] <= 3 {
			setMap[reservedSeats[j][0]][0] = 1
		} else if reservedSeats[j][1] >= 4 && reservedSeats[j][1] <= 5 {
			setMap[reservedSeats[j][0]][1] = 1
		} else if reservedSeats[j][1] >= 6 && reservedSeats[j][1] <= 7 {
			setMap[reservedSeats[j][0]][2] = 1
		} else if reservedSeats[j][1] >= 8 && reservedSeats[j][1] <= 9 {
			setMap[reservedSeats[j][0]][3] = 1
		}

	}

	for _, v := range setMap {
		if (v[0] == 1 && v[2] == 1) ||
			(v[1] == 1 && v[3] == 1) ||
			(v[1] == 1 && v[2] == 1) {
			count -= 2
			continue
		} else if v[0] == 0 && v[1] == 0 && v[2] == 0 && v[3] == 0 {

		} else {
			count -= 1
		}
	}

	return count
}

// 1540
// 735
