package main

import(
	"leetcode/array"
	"leetcode/list"
	"leetcode/string"
	"leetcode/integer"
	"fmt"
)

func main() {
	input := []int{2,7,11,15}
	result := array.TwoSum(input,9)
	fmt.Println(result)

	input = []int{-1,0,1,2,-1,-4}

	result1 := array.ThreeSum(input,0)
	fmt.Println(result1)

	l1 := list.GeneListByArray([]int{2,4,3})
	l2 := list.GeneListByArray([]int{5,6,4})

	result2 := list.AddTwoNumbers(l1,l2)
	list.Print(result2)

	in:= "abcabcbb"
	out := string.LengthOfLongestSubstring(in);
	fmt.Println(out)

	in = "babad"
	outStr := string.LongestPalindrome(in)
	fmt.Println(outStr)

	fmt.Println(integer.MAX_INT32)
	fmt.Println(integer.MIN_INT32)
}