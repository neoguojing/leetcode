package array

import (
	"fmt"
	"sort"
)

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
// 找到最便宜的航班路线
// n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
// 深度优先遍历
// 超时-未解决
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
<<<<<<< HEAD
	var minPrice = 100000
	dfs(findFlight(src, flights), flights, src, dst, k, 0, &minPrice)
=======
	sort.Sort(Flight(flights))
	minPrice := -1
	path := map[int]int{}

	dfs(path, findFlight(src, flights), flights, src, dst, k, 0, &minPrice)
>>>>>>> a5bd1d8cec45f19a96e7e5a9a52ea85c6b3e6e62

	return minPrice
}

<<<<<<< HEAD
=======
type Flight [][]int

func (a Flight) Len() int           { return len(a) }
func (a Flight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Flight) Less(i, j int) bool { return a[i][2] < a[j][2] }

>>>>>>> a5bd1d8cec45f19a96e7e5a9a52ea85c6b3e6e62
func findFlight(src int, flights [][]int) [][]int {
	ret := make([][]int, 0)
	for i := range flights {
		if flights[i][0] == src {
			ret = append(ret, flights[i])
		}
	}
	return ret
}
<<<<<<< HEAD
func dfs(srcFlights [][]int, flights [][]int, src int, dst int, k int, curPrice int, minPrice *int) {
	if curPrice > *minPrice {
		return
=======
func dfs(path map[int]int, srcFlights [][]int, flights [][]int, src int, dst int, k int, curPrice int, minPrice *int) bool {
	if len(flights) == 4851 {
		fmt.Println(len(srcFlights))
	}
	// 剪枝，防止路径上的环路
	if _, ok := path[src]; !ok {
		path[src] = 1
	} else {
		return false
>>>>>>> a5bd1d8cec45f19a96e7e5a9a52ea85c6b3e6e62
	}

	if k < 0 {
		delete(path, src)
		return false
	}

	if *minPrice != -1 && curPrice > *minPrice {
		delete(path, src)
		return false
	}

<<<<<<< HEAD
	for i := 0; i < len(srcFlights); i++ {

		if flights[i][0] == src {
			if flights[i][1] != dst {
				k--
			}
			curPrice += flights[i][2]
			dfs(findFlight(flights[i][1], flights), flights, flights[i][1], dst, k, curPrice, minPrice)
			curPrice -= flights[i][2]
			if flights[i][1] != dst {
=======
	if src == dst && (curPrice < *minPrice || *minPrice == -1) {
		delete(path, src)
		*minPrice = curPrice
		if len(flights) == 4851 {
			fmt.Println(curPrice)
		}
		return true
	}
	var ret bool = false
	for i := 0; i < len(srcFlights); i++ {

		if srcFlights[i][0] == src && (srcFlights[i][2] < *minPrice || *minPrice == -1) {
			path[srcFlights[i][0]] = 1
			if srcFlights[i][1] != dst {
				k--
			}
			curPrice += srcFlights[i][2]
			tmp := dfs(path, findFlight(srcFlights[i][1], flights), flights, srcFlights[i][1], dst, k, curPrice, minPrice)
			ret = ret || tmp
			curPrice -= srcFlights[i][2]
			if srcFlights[i][1] != dst {
>>>>>>> a5bd1d8cec45f19a96e7e5a9a52ea85c6b3e6e62
				k++
			}
		}
	}
	delete(path, src)
	return ret
}

// 1524
