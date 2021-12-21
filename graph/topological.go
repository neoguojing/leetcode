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
		for _, edge := range prerequisites {
			if edge[1] == point {
				ins[edge[0]]--
				cnt++
				if ins[edge[0]] == 0 {
					q.Push(edge[0])
				}
			}
		}
	}

	return cnt == numCourses
}
