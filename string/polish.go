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
