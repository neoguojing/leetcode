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
