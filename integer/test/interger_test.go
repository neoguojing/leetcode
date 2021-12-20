package test

import (
	"leetcode/integer"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	t.Log(integer.NthUglyNumber(10))

	t.Log(integer.NthUglyNumber2(5, 2, 11, 13))

	t.Log(integer.NthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func TestBase(t *testing.T) {
	t.Log(integer.GCD(2, 3))
	t.Log(integer.Sqrt(16))
	t.Log(integer.AddDigits(38))
	t.Log(2 / 10)
}

func TestMinimumPerimeter(t *testing.T) {
	t.Log(integer.MinimumPerimeter(1000000000))
	t.Log(integer.DayOfTheWeek(1, 1, 1971))
}

func TestGenerate(t *testing.T) {
	t.Log(integer.Generate(5))
}
