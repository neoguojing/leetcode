package string

import (
	"leetcode/utils"
	"strconv"
)

// EvalRPN 计算逆波兰式的值
// no 150
// ["2","1","+","3","*"] = （2+1）*3
func EvalRPN(tokens []string) int {
	stack := utils.NewStack()

	for i := range tokens {
		switch tokens[i] {
		case "+":
			val := stack.Pop().(int) + stack.Pop().(int)
			stack.Push(val)
		case "-":
			val := -(stack.Pop().(int)) + stack.Pop().(int)
			stack.Push(val)
		case "*":
			val := stack.Pop().(int) * stack.Pop().(int)
			stack.Push(val)
		case "/":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			val := b / a
			stack.Push(val)
		default:
			val, _ := strconv.Atoi(tokens[i])
			stack.Push(val)
		}
	}

	return stack.Pop().(int)
}

// Calculate 计算字符串所表示的值 "3+2*2"
// 227
func Calculate(s string) int {

	slice := make([]string, 0)
	word := ""
	for i := range s {
		if s[i] == ' ' {
			continue
		}

		if s[i] != '+' && s[i] != '-' && s[i] != '*' && s[i] != '/' {
			word += string(s[i])
		} else {
			slice = append(slice, word)
			word = ""
			slice = append(slice, string(s[i]))
		}
	}

	if word != "" {
		slice = append(slice, word)
		word = ""
	}

	stack := utils.NewStack()
	curOp := "+"
	curChar := -1
	for _, v := range slice {
		if v != "+" && v != "-" && v != "*" && v != "/" {
			curChar, _ = strconv.Atoi(v)
		}

		switch curOp {
		case "+":
			if curChar != -1 {
				stack.Push(curChar)
				curChar = -1
			}

		case "-":
			if curChar != -1 {
				stack.Push(-curChar)
				curChar = -1
			}

		case "*":
			if curChar != -1 {
				stack.Push(curChar * stack.Pop().(int))
				curChar = -1
			}
		case "/":
			if curChar != -1 {
				stack.Push(stack.Pop().(int) / curChar)
				curChar = -1
			}

		}
	}
	ret := 0
	for !stack.Empty() {
		ret += stack.Pop().(int)
	}
	return ret
}
