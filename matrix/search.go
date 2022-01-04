package matrix

import "container/heap"

/*Exist ...
no 79
DSF算法
1.需要位图标记走过的位置
2.设置边界，防止越界
3.设置方向，指导遍历
4.使用递归，处理子问题：递归结束条件：找到完整的字符借宿

[["a","b","c"],["a","e","d"],["a","f","g"]]
"abcdefg"

*/

func Exist(board [][]byte, word string) bool {
	if board == nil {
		return false
	}

	m := len(board)
	n := len(board[0])

	if m*n == 0 {
		return false
	}

	if m*n < len(word) {
		return false
	}

	bitMap := initVisitedGraph(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//有一个成功则返回
			if existRecur(board, word, 0, i, j, bitMap) {
				return true
			}
		}
	}
	return false
}

func existRecur(board [][]byte, word string, index int, x, y int, bitMap [][]bool) bool {
	m := len(board)
	n := len(board[0])
	//所有字符比较完毕，表示该路径可以用
	if index >= len(word) {
		return true
	}

	//越界表示该路不通
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}

	//已经访问过，则该路径不能用
	if bitMap[x][y] {
		return false
	}

	//当前字符不想等，则该路径不能用
	if board[x][y] != word[index] {
		return false
	}

	//设置当前位置为已访问
	bitMap[x][y] = true

	ret := false
	//像上遍历，返回true，则改路径通，结束递归
	ret = existRecur(board, word, index+1, x-1, y, bitMap)
	if ret {
		return true
	}

	ret = existRecur(board, word, index+1, x+1, y, bitMap)
	if ret {
		return true
	}

	ret = existRecur(board, word, index+1, x, y-1, bitMap)
	if ret {
		return true
	}

	ret = existRecur(board, word, index+1, x, y+1, bitMap)
	if ret {
		return true
	}
	//若四个方向均没有可用路径，则重置访问map，返回false
	bitMap[x][y] = false
	return false
}

// no 74
//判断一个矩阵中是否存在某个数，矩阵是有序的。
// todo

// SearchMatrix 矩阵的每行（左-》右）和每列（上-》下）都是有序的，快速搜索某个值
// 240 二分查找
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m := len(matrix) - 1
	n := len(matrix[0]) - 1
	if target < matrix[0][0] || target > matrix[m][n] {
		return false
	}

	return searchMatrixHelper(matrix, 0, m, 0, n, target)
}

func searchMatrixHelper(matrix [][]int, xb, xe, yb, ye, target int) bool {

	for xb <= xe && yb <= ye {
		mx := xb + (xe-xb)/2
		my := yb + (ye-yb)/2

		if matrix[mx][my] == target {
			return true
		} else if matrix[mx][my] > target {
			return searchMatrixHelper(matrix, xb, mx-1, yb, my-1, target) ||
				searchMatrixHelper(matrix, mx, xe, yb, my-1, target) ||
				searchMatrixHelper(matrix, xb, mx-1, my, ye, target)
		} else {

			return searchMatrixHelper(matrix, mx+1, xe, my+1, ye, target) ||
				searchMatrixHelper(matrix, xb, mx, my+1, ye, target) ||
				searchMatrixHelper(matrix, mx+1, xe, yb, my, target)

		}
	}

	return false
}

// 从右上角开始搜索
func SearchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m := len(matrix) - 1
	n := len(matrix[0]) - 1
	if target < matrix[0][0] || target > matrix[m][n] {
		return false
	}

	i, j := 0, n

	for i >= 0 && i <= m && j >= 0 && j <= n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}

// KthSmallest 在排序（非递减,左边小于右边，上小于下）矩阵中最小的第k个元素 时间复杂度小于o（N^2）
// 378
func KthSmallest(matrix [][]int, k int) int {
	if matrix == nil {
		return -1
	}
	pq := PriorityQueue{}
	for i := range matrix[0] {
		heap.Push(&pq, [3]int{0, i, matrix[0][i]})
	}

	for i := 0; i < k-1; i++ {
		tmp := heap.Pop(&pq).([3]int)
		if tmp[0] == len(matrix[0])-1 {
			continue
		}

		heap.Push(&pq, [3]int{tmp[0] + 1, tmp[1], matrix[tmp[0]+1][tmp[1]]})
	}

	return heap.Pop(&pq).([3]int)[2]
}

type PriorityQueue [][3]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i][2] < pq[j][2]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([3]int))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
