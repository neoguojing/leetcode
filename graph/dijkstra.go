package graph

import "fmt"

//FindCheapestPrice ...
// no 787
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

	minK := make([]int, k)
	stop := 1
	minK[0] = adj[src][dst]
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
			if j == dst {
				minK[stop] = dist[dst]
			}
			if stop == k {
				break
			}

			stop++
		}

	}

	if dist[dst] != INF {
		return dist[dst]
	}
	fmt.Println(minK)
	return -1
}
