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

func Sqrt(n int64) int64 {
	g := float64(n)

	for math.Abs(g*g-float64(n)) > 0.000001 {
		g = (g + float64(n)/g) / 2
	}

	if int64(g)*int64(g) == n {
		return int64(g)
	}

	return -1
}
