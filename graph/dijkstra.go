package graph

import "fmt"

//FindCheapestPrice ...
// no 787
// 通过k剪枝未完成，未解决
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	dist := make([]int, n)
	path := make([]int, n)
	set := make(map[int]bool)
	INF := 100000
	// 构建矩阵
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
		for j := range adj[i] {
			adj[i][j] = INF
		}
	}
	for _, arc := range flights {
		adj[arc[0]][arc[1]] = arc[2]
	}

	// init
	set[src] = true
	for i, w := range adj[src] {
		dist[i] = w
		path[i] = src

	}
	dist[src] = 0
	if k == 0 {
		if adj[src][dst] != INF {
			return adj[src][dst]
		} else {
			return -1
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(dist)
		min := INF
		selectNode := 0
		for j := 0; j < n; j++ {
			if !set[j] && dist[j] < min {

				min = dist[j]
				selectNode = j

			}

		}

		set[selectNode] = true
		for j := 0; j < n; j++ {
			if !set[j] && min+adj[selectNode][j] < dist[j] {
				dist[j] = min + adj[selectNode][j]
				path[j] = selectNode

			}
		}
		// stop++
	}

	if dist[dst] != INF {
		outPut := fmt.Sprintf("%v", dst)
		for i := path[dst]; i != src; i = path[i] {
			outPut = fmt.Sprintf("%v->%v", i, outPut)
		}
		outPut = fmt.Sprintf("%v->%v", src, outPut)
		fmt.Println(outPut)
		return dist[dst]
	}
	return -1
}
