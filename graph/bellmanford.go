package graph

import "fmt"

// FindCheapestPrice
// no 787
// dist[j] > dist[i] + adj[i,j]
func FindCheapestPriceBell(n int, flights [][]int, src int, dst int, k int) int {
	INF := 100000

	dist := make([]int, n)
	for i := range dist {
		if i == src {
			dist[i] = 0
		} else {
			dist[i] = INF
		}

	}
	// 限制路径长度
	for i := 0; i <= k; i++ {
		tmp := make([]int, n)
		copy(tmp, dist)
		for _, v := range flights {
			if dist[v[0]] != INF && tmp[v[1]] > dist[v[0]]+v[2] {
				tmp[v[1]] = dist[v[0]] + v[2]
			}
		}
		fmt.Println(tmp)
		dist = tmp
	}

	if dist[dst] != INF {
		return dist[dst]
	}
	return -1
}
