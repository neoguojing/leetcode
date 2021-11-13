package graph

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

	for m := 0; m < n; m++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][m]+dist[m][j] {
					dist[i][j] = dist[i][m] + dist[m][j]
				}
			}
		}
	}

	if dist[src][dst] != INF {
		return dist[src][dst]
	}
	return -1
}

// not work
func FindCheapestPriceFloydWithKStop(n int, flights [][]int, src int, dst int, k int) int {
	INF := 1000000
	dist := make([][][]int, k+2)
	for m := 0; m < k+2; m++ {
		dist[m] = make([][]int, n)
		for i := 0; i < n; i++ {
			dist[m][i] = make([]int, n)
			for j := 0; j < n; j++ {
				dist[m][i][j] = INF
			}
		}
	}
	for m := 0; m < k+2; m++ {
		for _, v := range flights {
			dist[m][v[0]][v[1]] = v[2]
		}
	}

	for m := 0; m < n; m++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for n := 1; n < k+2; n++ {
					if dist[n][i][j] > dist[n-1][i][m]+dist[n-1][m][j] {
						dist[n][i][j] = dist[n-1][i][m] + dist[n-1][m][j]
					}
				}

			}
		}
	}

	if dist[k][src][dst] != INF {
		return dist[k][src][dst]
	}
	return -1
}
