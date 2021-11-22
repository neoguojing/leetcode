package integer

import (
	"container/heap"
	"fmt"
	"math"
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
// no 1201 该解法超时
func NthUglyNumber1(n int, a int, b int, c int) int {
	if n < a && n < b && n < c {
		return 0
	}

	next := 0

	idx := make(map[int]int, 3)
	idx[a] = 1
	idx[b] = 1
	idx[c] = 1

	for i := 2; i <= n; i++ {
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
		next = min
		if next == x {
			idx[a]++
		}
		if next == y {
			idx[b]++
		}
		if next == z {
			idx[c]++
		}
	}
	return next
}

// no 1201
func NthUglyNumber2(n int, a int, b int, c int) int {
	lo, hi := int64(1), int64(1e9)*2

	A, B, C := int64(a), int64(b), int64(c)
	AB := LCM(A, B)
	BC := LCM(C, B)
	AC := LCM(A, C)
	ABC := LCM(A, BC)

	for lo < hi {
		mid := int64(lo + (hi-lo)/2)
		cnt := mid/A + mid/C + mid/B - mid/AB - mid/AC - mid/BC + mid/ABC
		if cnt < int64(n) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return int(lo)
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

// NthSuperUglyNumber
// 313
func NthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	if len(primes) == 0 {
		return 1
	}

	heapInt := make([]Prime, len(primes))
	idx := make([]int, len(primes))
	for i, v := range primes {
		heapInt[i] = Prime{Index: i, Val: v}
		idx[i] = 1
	}

	h := IntHeap(heapInt)
	heap.Init(&h)
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {

		dp[i] = heapInt[0].Val
		for heapInt[0].Val == dp[i] {
			index := heapInt[0].Index
			idx[index]++
			heapInt[0].Val = primes[index] * dp[idx[index]]
			heap.Fix(&h, 0)
		}

	}
	return dp[n]
}

type Prime struct {
	Index int
	Val   int
}
type IntHeap []Prime

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Prime))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NumSquares 由数的平方相加等于n最少需要几个数字？
// 279
func NumSquares(n int) int {
	sqrts := make([]int, 0)
	var i int = 1
	for ; i*i <= n; i++ {
		sqrts = append(sqrts, i*i)
	}
	if i*i == n {
		return 1
	}
	var count int = int(1e4)
	nSum(n, 0, int(1e4), sqrts, &count)
	return count
}

func nSum(n, tmp int, cur int, set []int, cnt *int) {
	if cur > *cnt {
		return
	}

	if tmp > n {
		return
	}

	if tmp == n {
		*cnt = cur
		return
	}

	for i := len(set) - 1; i >= 0; i-- {

		cur++
		if cur < (*cnt) {
			tmp += set[i]
			nSum(n, tmp, cur, set, cnt)
			tmp -= set[i]
		}
		cur--
	}
}

// 279
func NumSquaresDP(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 1000000
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for k := 1; k*k < i; k++ {
			tmp := dp[i-k*k] + 1
			if dp[i] > tmp {
				dp[i] = tmp
			}
		}
	}

	return dp[n]
}

// AddDigits
// 258
func AddDigits(num int) int {
	if num < 10 {
		return num
	}
	tmp := 0
	for num/10 != 0 {
		tmp += num % 10
		num = num / 10
	}
	tmp += num
	fmt.Println(tmp)
	return AddDigits(tmp)
}

func AddDigits1(num int) int {
	if num == 0 {
		return 0
	}

	return 1 + (num-1)%9
}

// MinimumPerimeter 求目标苹果数量苹果的周长
// 1954
func MinimumPerimeter(neededApples int64) int64 {
	if neededApples == 0 {
		return 0
	}
	n := math.Sqrt(float64(neededApples/12)) + 1
	dp := make([]int64, int(n)+1)
	dp[0] = 0

	for i := 1; i <= int(n); i++ {
		dp[i] = dp[i-1] + int64(12*i*i)
		if dp[i] >= neededApples {
			return int64(8 * i)
		}
	}

	return 0
}

// 372
// 1131
// 1185
