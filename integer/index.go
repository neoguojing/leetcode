package integer

import (
	"container/heap"
	"fmt"
)

const MAX_INT32 = int32(^uint32(0) >> 1)
const MIN_INT32 = ^MAX_INT32

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	signed := func(s byte) int {
		if s == '-' {
			return -1
		}
		return 1
	}
	ret := 0
	sign := byte('+')
	hasNumBefore := false
	hasSigned := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+':
			if hasNumBefore {
				goto OUT
			}
			if hasSigned {
				goto OUT
			}
			hasSigned = true
			sign = s[i]
		case '-':
			if hasNumBefore {
				goto OUT
			}
			if hasSigned {
				goto OUT
			}
			hasSigned = true
			sign = s[i]
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			hasNumBefore = true
			if signed(sign) > 0 {
				if ret > int(MAX_INT32)/10 {
					return int(MAX_INT32)
				}

				if ret == int(MAX_INT32)/10 && int(s[i]-'0') > 7 {
					return int(MAX_INT32)
				}
			} else {
				if -ret < int(MIN_INT32)/10 {
					return int(MIN_INT32)
				}

				if -ret == int(MIN_INT32)/10 && int(s[i]-'0') > 8 {
					return int(MIN_INT32)
				}
			}
			ret = ret*10 + int(s[i]-'0')

		case '.':
			goto OUT
		case ' ':
			if hasNumBefore {
				goto OUT
			}
			if hasSigned {
				goto OUT
			}
			continue
		default:
			if !hasNumBefore {
				return 0
			}
			goto OUT
		}
	}
OUT:
	if sign == '-' {
		ret = -ret
	}

	return ret
}

// NthUglyNumber ...
// no 264 返回第n个丑数
func NthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1

	idx := make(map[int]int, 3)
	idx[2] = 1
	idx[3] = 1
	idx[5] = 1
	for i := 2; i <= n; i++ {
		a := dp[idx[2]] * 2
		b := dp[idx[3]] * 3
		c := dp[idx[5]] * 5
		min := a
		if min > b {
			min = b
		}

		if min > c {
			min = c
		}

		dp[i] = min

		if dp[i] == a {
			idx[2]++
		}
		if dp[i] == b {
			idx[3]++
		}
		if dp[i] == c {
			idx[5]++
		}
	}
	fmt.Println(dp)
	return dp[n]
}

// IsUgly pa'm
// no 263
//
func IsUgly(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%2 != 0 && n%3 != 0 && n%5 != 0 {
		return false
	}

	if n%5 == 0 {
		for ; n%5 == 0; n = n / 5 {
		}
	}

	if n%3 == 0 {
		for ; n%3 == 0; n = n / 3 {
		}
	}

	if n%2 == 0 {
		for ; n%2 == 0; n = n / 2 {
		}
	}

	if n != 1 {
		return false
	}

	return true
}

// NthUglyNumber1：质数自己定义
// no 1201
func NthUglyNumber1(n int, a int, b int, c int) int {
	if n < a && n < b && n < c {
		return 0
	}

	idx := make(map[int]int, 3)
	idx[a] = 1
	idx[b] = 1
	idx[c] = 1

	dp := make([]int, n+2)
	dp[1] = 1

	for i := 2; i <= n+1; i++ {

		x := idx[a] * a
		y := idx[b] * b
		z := idx[c] * c

		min := x
		if min > y {
			min = y
		}
		if min > z {
			min = z
		}
		fmt.Println(idx[a])
		dp[i] = min
		if dp[i] == x {
			idx[a]++
		}
		if dp[i] == y {
			idx[b]++
		}
		if dp[i] == z {
			idx[c]++
		}
	}
	fmt.Println(dp)
	return dp[n+1]
}

// IsHappy
// no 202
func IsHappy(n int) bool {
	if n == 1 {
		return true
	}
	set := make(map[int]bool)
	set[n] = true
	for n != 1 {
		tmp := 0
		for n/10 != 0 {
			tmp += (n % 10) * (n % 10)
			n = n / 10
		}
		n = tmp + n*n
		if _, ok := set[n]; ok {
			return false
		}
		set[n] = true
	}

	return true
}

// CountPrimes 统计小于n的质数的个数
// no 204
func CountPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	// table索引为质数本身，值为是否质数
	table := make([]byte, n)
	cnt := 0
	for i := 2; i < n; i++ {

		if table[i] == 0 {
			cnt++
			// k为倍数
			k := 2
			for i*k < n {
				table[i*k] = 1
				k++
			}
		}
	}

	return cnt
}

// 1201

// NthSuperUglyNumber
// 313
// dp[i][j] 第i个primes在第j轮的值
func NthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	if len(primes) == 0 {
		return 1
	}

	heapInt := make([]int, len(primes))
	copy(heapInt, primes)
	h := IntHeap(heapInt)
	heap.Init(&h)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
		dp[i][1] = 1
	}

	for i := 2; i <= n; i++ {
		heap.Fix(&h, 0)
	}

	return dp[n]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 279
// 258
// 1954
// 279
