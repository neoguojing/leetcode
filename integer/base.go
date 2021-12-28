package integer

import (
	"fmt"
	"math"
)

func GCD(a, b int64) int64 {

	if a < b {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(a, b int64) int64 {
	gcd := GCD(a, b)

	return a * b / gcd
}

// Sqrt
// no 69
func Sqrt(x int64) int64 {
	g := float64(x)

	for math.Abs(g*g-float64(x)) > 0.000001 {
		g = (g + float64(x)/g) / 2
	}

	if int64(g)*int64(g) == x {
		return int64(g)
	}

	return -1
}

// FractionToDecimal numerator/denominator 返回字符串表示的小数，循环小数用括号包裹
// 小数计算
func FractionToDecimal(numerator int, denominator int) string {
	ret := ""

	for numerator >= denominator {
		ret += fmt.Sprintf("%d", numerator/denominator)
		numerator %= denominator
	}

	if numerator == 0 {
		return ret
	}

	if len(ret) == 0 {
		ret = "0."
	} else {
		ret += "."
	}

	set := map[int]bool{}
	little := ""
	repeat := ""
	for {
		for numerator < denominator {
			numerator = numerator * 10
			if _, ok := set[numerator]; !ok {
				set[numerator] = true
				if numerator < denominator {
					little += "0"
				}
			} else {
				if numerator < denominator {
					repeat += "0"
				}
			}
		}

		tmp := numerator / denominator
		if _, ok := set[numerator]; !ok {
			set[numerator] = true
			little += fmt.Sprintf("%d", tmp)
		} else {
			repeat += fmt.Sprintf("%d", tmp)
		}

		numerator = numerator % denominator
		if numerator == 0 {
			break
		}

		if repeat == little {
			little = "(" + little + ")"
			break
		}
	}

	return ret + little
}

// TrailingZeroes n的阶乘的结果的结尾0的个数
//  no 172
// n = 5  5! = 120 答案是1
// 会越界
func TrailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	count := 0
	for n >= 5 {
		if n%5 == 0 {
			count += n / 5
			n = n / 5
		} else {
			n = n - (n % 5)
		}
	}
	return count
}
