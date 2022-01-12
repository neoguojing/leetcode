package string

import (
	"leetcode/utils"
)

/*IsParenthesesValid ...
20 括号匹配问题
*/
func IsParenthesesValid(in string) bool {
	s := utils.NewStack()

	for _, v := range in {
		switch v {
		case '(', '[', '{':
			s.Push(v)
		case ')':
			if s.Empty() {
				return false
			}
			if s.Peak().(byte) == '(' {
				s.Pop()
			} else {
				return false
			}
		case '}':
			if s.Empty() {
				return false
			}
			if s.Peak().(byte) == '{' {
				s.Pop()
			} else {
				return false
			}

		case ']':
			if s.Empty() {
				return false
			}
			if s.Peak().(byte) == '[' {
				s.Pop()
			} else {
				return false
			}
		default:
			continue
		}
	}

	if s.Empty() {
		return true
	}

	return false
}

//LongestValidParentheses 最长合法括号组
// 32
func LongestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	stack := utils.NewStack()
	stack.Push(-1)
	max := 0
	for i := range s {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.Empty() {
				stack.Push(i)
			} else {
				if i-stack.Peak().(int) > max {
					max = i - stack.Peak().(int)
				}
			}
		}
	}

	return max
}
