package test

import (
	"leetcode/combination"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	in := "234"
	ret := combination.LetterCombinations(in)
	t.Log(ret)

	// ret = combination.LetterCombinationsWithQueue(in)
	// t.Log(ret)
}

func TestStringVectorMul(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	ret := combination.StringVectorMul1(a, b)
	t.Log(ret)
	ret = combination.StringVectorMul1(a, ret)
	t.Log(ret)

}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(combination.GenerateParenthesis(3))
}
