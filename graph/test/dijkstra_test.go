package test

import (
	"leetcode/graph"
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {

	flight := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}

	t.Log(graph.FindCheapestPrice(4, flight, 0, 3, 1))
}
