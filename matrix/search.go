package matrix

/*Exist ...
DSF算法
1.需要位图标记走过的位置
2.设置边界，防止越界
3.设置方向，指导遍历
4.使用递归，处理子问题：递归结束条件：找到完整的字符借宿
*/
func Exist(board [][]rune, word string) bool {
	if board == nil {
		return false
	}

	m := len(board)
	n := len(board[0])

	if m*n == 0 {
		return false
	}

	if m+n < len(word) {
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

func existRecur(board [][]rune, word string, index int, x, y int, bitMap [][]bool) bool {
	m := len(board)
	n := len(board[0])
	//越界表示该路不通
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	//所有字符比较完毕，表示该路径可以用
	if index >= len(word) {
		return true
	}
	//已经访问过，则该路径不能用
	if bitMap[x][y] {
		return false
	}

	//当前字符不想等，则该路径不能用
	if board[x][y] != rune(word[index]) {
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
