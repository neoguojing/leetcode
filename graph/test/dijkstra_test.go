package test

import (
	"leetcode/graph"
	"testing"
)

func TestFindCheapestPrice(t *testing.T) {

	flight := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}

	t.Log(graph.FindCheapestPrice(4, flight, 0, 3, 1))

	flight = [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}}

	t.Log(graph.FindCheapestPrice(5, flight, 0, 2, 2))
}

func TestFindCheapestPriceBell(t *testing.T) {

	flight := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}

	t.Log(graph.FindCheapestPriceBell(4, flight, 0, 3, 1))

	flight = [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}}

	t.Log(graph.FindCheapestPriceBell(5, flight, 0, 2, 2))
}

func TestFindCheapestPriceFloyd(t *testing.T) {

	flight := [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}

	t.Log(graph.FindCheapestPriceFloyd(4, flight, 0, 3, 1))

	flight = [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}}

	t.Log(graph.FindCheapestPriceFloyd(5, flight, 0, 2, 2))
}
