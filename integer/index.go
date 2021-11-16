package integer

import "fmt"

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
		fmt.Println(idx)
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

		if dp[i] == dp[idx[2]]*2 {
			idx[2]++
		}
		if dp[i] == dp[idx[3]]*3 {
			idx[3]++
		}
		if dp[i] == dp[idx[5]]*5 {
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

// no 202
// no 204
// 1201
// 313
// 279
