package classify

// NumIslands ...
//  no 200
// 一个二维数组，把 1 看做陆地，把 0 看做大海，陆地相连组成一个岛屿。把数组以外的区域也看做是大海，问总共有多少个岛屿
func NumIslands(grid [][]int) int {
	nums := 0
	m := len(grid)
	n := len(grid[0])
	trees := make([]*UnionFind, m*n)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				elem := i*n + j
				trees[elem] = MakeSet(elem)
				nums++
			}
		}
	}

	for _, v := range trees {
		if v == nil {
			continue
		}

		elem := v.Val
		i := elem / n
		j := elem % n

		if i*n+j-1 > 0 && trees[i*n+j-1] != nil {
			if v.Parent.Val != trees[i*n+j-1].Parent.Val {
				Union(v, trees[i*n+j-1])
				nums--
			}

		}

		if (i-1)*n+j > 0 && trees[(i-1)*n+j] != nil {
			if v.Parent.Val != trees[(i-1)*n+j].Parent.Val {

				Union(v, trees[(i-1)*n+j])
				nums--
			}
		}

		if i*n+j+1 < m*n && trees[i*n+j+1] != nil {
			if v.Parent.Val != trees[i*n+j+1].Parent.Val {

				Union(v, trees[i*n+j+1])
				nums--
			}
		}

		if (i+1)*n+j < m*n && trees[(i+1)*n+j] != nil {
			if v.Parent.Val != trees[(i+1)*n+j].Parent.Val {

				Union(v, trees[(i+1)*n+j])
				nums--
			}
		}
	}

	return nums
}

// Solve ...
// no 130
// 有一点点像围棋，把被 X 围起来的 O 变成 X，边界的 O
// 一定不会被围起来。如果 O 和边界的 O 连通起来，那么这些 O 就都算作不被围起来
func Solve(board [][]string) [][]string {
	m := len(board)
	n := len(board[0])

	aliveSet := MakeSet(-1)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == "O" {
				cur := MakeSet(i*n + j)
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					Union(aliveSet, cur)
				} else {
					if board[i-1][j] == "O" {
						tmp := MakeSet((i-1)*n + j)
						Union(cur, tmp)
					}
					if board[i+1][j] == "O" {
						tmp := MakeSet((i+1)*n + j)
						Union(cur, tmp)
					}
					if board[i][j-1] == "O" {
						tmp := MakeSet(i*n + j - 1)
						Union(cur, tmp)
					}
					if board[i][j+1] == "O" {
						tmp := MakeSet(i*n + j + 1)
						Union(cur, tmp)
					}
				}

			}
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == "O" {
				if Find(MakeSet(i*n+j)).Parent.Val != -1 {
					board[i][j] = "X"
				}
			}
		}
	}

	return board
}
