package array

import "fmt"

// no 518
// coins:找到合适的组合满足硬币需要
/*
amount = 5, coins = [1,2,5]
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
*/
func Change1(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}

	cur, target := 0, 0
	tmp := make([]int, 0)
	changeBackward(amount, cur, coins, tmp, &target)
	return target
}

func changeBackward(amount, cur int, coins []int, tmp []int, target *int) {
	if cur > amount {
		return
	}
	if cur == amount {
		fmt.Println(tmp)
		*target += 1
		return
	}
	for i := 0; i < len(coins); i++ {
		if coins[i] > amount {
			continue
		}
		if cur+coins[i] > amount {
			continue
		}

		cur += coins[i]
		tmp = append(tmp, coins[i])
		changeBackward(amount, cur, coins, tmp, target)
		cur -= coins[i]
		tmp = tmp[0 : len(tmp)-1]
	}
}

// FindCheapestPrice
// no 787
// 找到最偏移的航班路线
// n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
// 深度优先遍历
/*
5
[[1,2,10],[2,0,7],[1,3,8],[4,0,10],[3,4,2],[4,2,10],[0,3,3],[3,1,6],[2,4,5]]
0
4
1
*/
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var minPrice = 100000
	dfs(findFlight(src, flights), flights, src, dst, k, 0, &minPrice)

	if minPrice == 100000 {
		minPrice = -1
	}
	return minPrice
}

func findFlight(src int, flights [][]int) [][]int {
	ret := make([][]int, 0)
	for i := range flights {
		if flights[i][0] == src {
			ret = append(ret, flights[i])
		}
	}
	return ret
}
func dfs(srcFlights [][]int, flights [][]int, src int, dst int, k int, curPrice int, minPrice *int) {
	if curPrice > *minPrice {
		return
	}

	if k < 0 && src != dst {
		return
	}

	if k >= 0 && curPrice < *minPrice && src == dst {
		*minPrice = curPrice
		return
	}

	for i := 0; i < len(srcFlights); i++ {

		if flights[i][0] == src {
			if flights[i][1] != dst {
				k--
			}
			curPrice += flights[i][2]
			dfs(findFlight(flights[i][1], flights), flights, flights[i][1], dst, k, curPrice, minPrice)
			curPrice -= flights[i][2]
			if flights[i][1] != dst {
				k++
			}
		}
	}
}

// 1524
