package integer

import (
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

}
