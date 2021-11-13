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

func FindCheapestPriceBellWithKStop(n int, flights [][]int, src int, dst int, k int) int {
	INF := 100000

	dist := make([][]int, k+2)
	for o := 0; o < k+2; o++ {
		dist[o] = make([]int, n)
		for i := range dist[o] {
			fmt.Println(i)
			if i == src {
				dist[o][i] = 0
			} else {
				dist[o][i] = INF
			}

		}
	}
	// 限制路径长度
	for i := 1; i <= k+1; i++ {
		for _, v := range flights {
			if dist[i][v[1]] > dist[i-1][v[0]]+v[2] {
				dist[i][v[1]] = dist[i-1][v[0]] + v[2]
			}
		}
	}
	if dist[k+1][dst] != INF {
		return dist[k+1][dst]
	}
	return -1
}
