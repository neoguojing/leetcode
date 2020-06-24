package test

import (
	"leetcode/string"
	"testing"
)

func TestMinDistance(t *testing.T) {
	a := "horse"
	b := "ros"
	n := string.MinDistance(a, b)
	if n != 3 {
		t.Error(n)
	}

	a = "intention"
	b = "execution"

	n = string.MinDistance(a, b)
	if n != 5 {
		t.Error(n)
	}
}
