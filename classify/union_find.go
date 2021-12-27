package classify

import (
	"math"
)

// NumIslands ...
//  no 200
// 一个二维数组，把 1 看做陆地，把 0 看做大海，陆地相连组成一个岛屿。把数组以外的区域也看做是大海，问总共有多少个岛屿
func NumIslands(grid [][]int) int {
	nums := 0
	m := len(grid)
	n := len(grid[0])
	trees := MakeSet()
	for i := range grid {
		for j := range grid[i] {
			//为1的值创建集合
			if grid[i][j] == '1' {
				elem := i*n + j
				trees.NewElem(elem)
				nums++
			}
		}
	}

	for k, v := range trees {
		if v == nil {
			continue
		}

		//计算坐标
		elem := k.(int)
		i := elem / n
		j := elem % n

		//左面的值为1，且不属于一个集合，则合并
		if j-1 >= 0 && j-1 < n && trees[i*n+j-1] != nil {
			if v.Parent.Val != trees[i*n+j-1].Parent.Val {
				trees.Union(k.(int), trees[i*n+j-1].Val)
				nums--
			}

		}
		//上面的值为1
		if i-1 >= 0 && i-1 < m && trees[(i-1)*n+j] != nil {
			if v.Parent.Val != trees[(i-1)*n+j].Parent.Val {
				trees.Union(k.(int), trees[(i-1)*n+j].Val)
				nums--
			}
		}
		// 右面的值为1
		if j+1 >= 0 && j+1 < n && trees[i*n+j+1] != nil {
			if v.Parent.Val != trees[i*n+j+1].Parent.Val {

				trees.Union(k.(int), trees[i*n+j+1].Val)
				nums--
			}
		}
		// 下面的只为1
		if i+1 >= 0 && i+1 < m && trees[(i+1)*n+j] != nil {
			if v.Parent.Val != trees[(i+1)*n+j].Parent.Val {

				trees.Union(k.(int), trees[(i+1)*n+j].Val)
				nums--
			}
		}
	}
	set := map[int]int{}
	for _, v := range trees {
		set[v.Parent.Val] = 1
	}
	return len(set)
}

// Solve ...
// no 130
// 有一点点像围棋，把被 X 围起来的 O 变成 X，边界的 O
// 一定不会被围起来。如果 O 和边界的 O 连通起来，那么这些 O 就都算作不被围起来
func Solve(board [][]string) [][]string {
	m := len(board)
	n := len(board[0])
	aliveSet := MakeSet()
	// 创建不需要重置的集合
	aliveSet.NewElem(-1)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == "O" {
				cur := i*n + j
				aliveSet.NewElem(cur)
				// 若为边角元素，则合并到-1的集合
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					aliveSet.Union(-1, cur)
				} else {
					// 非边角元素，则与当前元素合并
					if board[i-1][j] == "O" {
						tmp := (i-1)*n + j
						aliveSet.NewElem(tmp)

						aliveSet.Union(cur, tmp)
					}
					if board[i+1][j] == "O" {
						tmp := (i+1)*n + j
						aliveSet.NewElem(tmp)

						aliveSet.Union(cur, tmp)
					}
					if board[i][j-1] == "O" {
						tmp := i*n + j - 1
						aliveSet.NewElem(tmp)

						aliveSet.Union(cur, tmp)
					}
					if board[i][j+1] == "O" {
						tmp := i*n + j + 1
						aliveSet.NewElem(tmp)

						aliveSet.Union(cur, tmp)
					}
				}

			}
		}
	}

	for i := range board {
		for j := range board[0] {
			// 若不在-1集合中则重置
			if board[i][j] == "O" {
				if aliveSet.Find(i*n+j).Parent.Val != -1 {
					board[i][j] = "X"
				}
			}
		}
	}

	return board
}

//LongestConsecutive 要求o(n)
// 128 最长连续子序列
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	setMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		setMap[nums[i]] = nums[i]
	}

	for i := range nums {
		_, ok1 := setMap[nums[i]-1]
		if ok1 {
			setMap[nums[i]] = nums[i] - 1
		}
	}

	for i := range nums {
		find(nums[i], setMap)
	}

	countMap := map[int]int{}
	max := 0
	for _, v := range setMap {
		countMap[v]++
		if countMap[v] > max {
			max = countMap[v]
		}
	}

	return max
}

func find(x int, setMap map[int]int) int {
	parent, ok := setMap[x]
	if !ok {
		return math.MaxInt32
	}
	if parent != x {
		setMap[x] = find(parent, setMap)
	}

	return setMap[x]
}
