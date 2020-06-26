package test

import (
	strings "leetcode/string"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	ret := strings.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	t.Log(ret)

	ret = strings.LongestCommonPrefix([]string{"dog", "cat", "bee"})
	t.Log(ret)
}
