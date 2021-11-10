package test

import (
	"fmt"
	"leetcode/array"
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	ret := array.FindCheapestPrice(3, flights, 0, 2, 1)
	fmt.Println(ret)
	t.Log(ret)
}
