package matrix

import (
	"leetcode/utils"
)

//UniquePaths ...
//no 62
//机器人从左上角走到右下角，只能向右或者向下走。输出总共有多少种走法。
// 动态规划：dp[i][j] = dp[i-1][j] + dp[i][j-1],当i=0，dp[0][j] = 1 同理dp[i][0]=1
func UniquePaths(m, n int) int {
	dp := iniDP(m, n, 0)
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

//UniquePathsWithObstacles ...
//no 63
//增加了一些不能走的格子，用 1 表示。还是输出从左上角到右下角总共有多少种走法
// 动态规划：同上，当遇到障碍是的dp值为0
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := iniDP(m, n, 0)

	hasObstacle := false
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			hasObstacle = true
		}
		if hasObstacle {
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}
	hasObstacle = false
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			hasObstacle = true
		}
		if hasObstacle {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

//MinPathSum ...
//no 64
//依旧是62题的扩展，这个是输出从左上角到右下角，路径的数字加起来和最小是多少。
// dp[i][j] = min(dp[i][j-1],dp[i-1][j])+grid[i][j]
func MinPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	dp := iniDP(m, n, 0)
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// LongestIncreasingPath 求矩阵的最长递增路径
// 329 拓扑排序
func LongestIncreasingPath(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])
	inDegree := make([][]int, m)
	for i := range matrix {
		inDegree[i] = make([]int, n)
		for j := range matrix[i] {
			if i-1 >= 0 && matrix[i-1][j] < matrix[i][j] {
				inDegree[i][j]++
			}

			if i+1 < m && matrix[i+1][j] < matrix[i][j] {
				inDegree[i][j]++
			}

			if j+1 < n && matrix[i][j+1] < matrix[i][j] {
				inDegree[i][j]++
			}

			if j-1 >= 0 && matrix[i][j-1] < matrix[i][j] {
				inDegree[i][j]++
			}
		}
	}

	q := utils.NewQueue()
	for i := range matrix {
		for j := range matrix[i] {
			if inDegree[i][j] == 0 {
				q.Push([]int{i, j})
			}
		}
	}
	max := 0
	for !q.Empty() {

		for k := 0; k < q.Len(); k++ {
			pos := q.Pop().([]int)
			i := pos[0]
			j := pos[1]

			if i-1 >= 0 && matrix[i-1][j] > matrix[i][j] {
				inDegree[i-1][j]--
				if inDegree[i-1][j] == 0 {
					q.Push((i-1)*m + j)
				}
			}

			if i+1 < m && matrix[i+1][j] > matrix[i][j] {
				inDegree[i+1][j]--
				if inDegree[i+1][j] == 0 {
					q.Push((i+1)*m + j)
				}
			}

			if j+1 < n && matrix[i][j+1] > matrix[i][j] {
				inDegree[i][j+1]--
				if inDegree[i][j+1] == 0 {
					q.Push((i)*m + j + 1)
				}
			}

			if j-1 >= 0 && matrix[i][j-1] > matrix[i][j] {
				inDegree[i][j-1]--
				if inDegree[i][j-1] == 0 {
					q.Push((i)*m + j - 1)
				}
			}
		}
		max++
	}
	return max
}
