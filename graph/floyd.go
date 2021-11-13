package graph

import "fmt"

// FindCheapestPriceFloyd ...
// no 787
func FindCheapestPriceFloyd(n int, flights [][]int, src int, dst int, k int) int {
	INF := 1000000
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}

	for _, v := range flights {
		dist[v[0]][v[1]] = v[2]
	}
	fmt.Println(dist)

	for m := 0; m < n; m++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][m]+dist[m][j] {
					dist[i][j] = dist[i][m] + dist[m][j]
				}
			}
		}
	}

	fmt.Println(dist)
	if dist[src][dst] != INF {
		return dist[src][dst]
	}
	return -1
}
