package utils

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	list := list.New()
	return &Queue{list}
}

func (queue *Queue) Push(value interface{}) {
	queue.list.PushBack(value)
}

func (queue *Queue) Pop() interface{} {
	e := queue.list.Front()
	if e != nil {
		queue.list.Remove(e)
		return e.Value
	}
	return nil
}

func (queue *Queue) PopBack() interface{} {
	e := queue.list.Back()
	if e != nil {
		queue.list.Remove(e)
		return e.Value
	}
	return nil
}

func (queue *Queue) Peak() interface{} {
	e := queue.list.Front()
	if e != nil {
		return e.Value
	}

	return nil
}

func (queue *Queue) Len() int {
	return queue.list.Len()
}

func (queue *Queue) Empty() bool {
	return queue.list.Len() == 0
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i] > pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
