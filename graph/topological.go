package graph

import "leetcode/utils"

// CanFinish 是否能完成课程 prerequisites表示要完成前面的课程，必须先完成后面的课程
// 207
func CanFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	ins := make([]int, numCourses)

	for _, edge := range prerequisites {
		ins[edge[0]]++
	}

	q := utils.NewQueue()
	cnt := 0
	for i, in := range ins {
		if in == 0 {
			q.Push(i)
		}
	}

	if q.Empty() {
		return false
	}

	for !q.Empty() {
		point := q.Pop()
		cnt++
		for _, edge := range prerequisites {
			if edge[1] == point {
				ins[edge[0]]--
				if ins[edge[0]] == 0 {
					q.Push(edge[0])
				}
			}
		}
	}

	return cnt == numCourses
}

// FindOrder 排序
// 210
func FindOrder(numCourses int, prerequisites [][]int) []int {
	ret := []int{}
	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			ret = append(ret, i)
		}
		return ret
	}

	ins := make([]int, numCourses)

	for _, edge := range prerequisites {
		ins[edge[0]]++
	}

	q := utils.NewQueue()
	cnt := 0
	for i, in := range ins {
		if in == 0 {
			q.Push(i)
		}
	}

	if q.Empty() {
		return ret
	}

	for !q.Empty() {
		point := q.Pop()
		cnt++
		ret = append(ret, point.(int))
		for _, edge := range prerequisites {
			if edge[1] == point {
				ins[edge[0]]--
				if ins[edge[0]] == 0 {
					q.Push(edge[0])
				}
			}
		}
	}

	if cnt != numCourses {
		return []int{}
	}
	return ret
}

// 630
// 310
