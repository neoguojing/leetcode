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
