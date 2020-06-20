package test

import(
	"testing"
	"leetcode/combination"
)


func TestLetterCombinations(t *testing.T){
	in := "234"
	ret := combination.LetterCombinations(in)
	t.Log(ret)

	ret = combination.LetterCombinationsWithQueue(in)
	t.Log(ret)
}