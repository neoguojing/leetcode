package stack

import "math"

type Stack []int

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}

func (s *Stack) Pop() int {
	if len(*s) > 0 {
		top := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return top
	}

	return -1
}

func (s *Stack) Top() int {
	if len(*s) > 0 {
		return (*s)[len(*s)-1]
	}
	return -1
}

type MinStack struct {
	head *Element
}

type Element struct {
	Val  int
	Min  int
	Next *Element
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head != nil {
		this.head = &Element{Val: val, Min: int(math.Min(float64(val), float64(this.head.Min))), Next: this.head}
	} else {
		this.head = &Element{Val: val, Min: val, Next: nil}
	}
}
func (this *MinStack) Pop() {
	this.head = this.head.Next
}

func (this *MinStack) Top() int {
	return this.head.Val
}

func (this *MinStack) GetMin() int {
	return this.head.Min
}
