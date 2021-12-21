package graph

import (
	"leetcode/utils"
)

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

func FindOrderDFS(numCourses int, prerequisites [][]int) []int {
	visited := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	res := []int{}
	for i := 0; i < numCourses; i++ {
		if !dfs(graph, visited, i, &res) {
			return []int{}
		}
	}
	result := make([]int, 0)
	for i := numCourses - 1; i >= 0; i-- {
		result = append(result, res[i])
	}

	return result
}

func dfs(graph [][]int, visited []int, course int, res *[]int) bool {
	if visited[course] == 1 {
		return false
	} else if visited[course] == 2 {
		return true
	} else {
		visited[course] = 1
	}

	for i := 0; i < len(graph[course]); i++ {
		if !dfs(graph, visited, graph[course][i], res) {
			return false
		}
	}

	*res = append(*res, course)
	visited[course] = 2
	return true
}

// 630
// 310
