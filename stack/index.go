package stack

import (
	"container/list"
	"math"
)

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
	min  int
	list *list.List
}

func Constructor() MinStack {
	return MinStack{
		list: list.New(),
		min:  math.MaxInt64,
	}
}

func (this *MinStack) Push(val int) {
	if val < this.min {
		this.min = val
	}
	this.list.PushBack(val)
}

func (this *MinStack) Pop() {
	e := this.list.Back()
	if e != nil {
		this.list.Remove(e)
	}
}

func (this *MinStack) Top() int {
	e := this.list.Back()
	if e != nil {
		return -1
	}

	return e.Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.min
}
