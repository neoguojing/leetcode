package string

import()

//no 3给定一个字符串，找到没有重复字符的最长子串，返回它的长度。
//滑动窗口
func LengthOfLongestSubstring(in string) int{
	max := 0
	set := make(map[byte]int)
	for start,end:=0,0;end<len(in);end++ {
		if _,ok := set[in[end]];ok {
			if max < (end-start) {
				max = end - start
			}
			delete(set,in[start])
			start++
		}
		set[in[end]] = 1
	}

	return max
}