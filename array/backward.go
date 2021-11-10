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

func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var minPrice = 0

	dfs(n, flights, src, dst, k, minPrice, &minPrice)
	return minPrice
}

func dfs(n int, flights [][]int, src int, dst int, k int, curPrice int, minPrice *int) {
	if curPrice > *minPrice {
		return
	}

	if k == 0 && src != dst {
		return
	}

	if k >= 0 && curPrice < *minPrice && src == dst {
		*minPrice = curPrice
		return
	}

	for i := 0; i < n; i++ {

		if flights[i][0] == src && flights[i][1] == dst {
			curPrice += flights[i][2]
			if curPrice < *minPrice {
				*minPrice = curPrice
			}
			continue
		}
		if flights[i][0] == src {
			k--
			curPrice += flights[i][2]
			dfs(n, flights, flights[i][1], dst, k, curPrice, minPrice)
			k++
			curPrice -= flights[i][2]
		}
	}
}

// 1524
