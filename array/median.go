package array

import (
	"container/heap"
	"gopls-workspace/utils"
)

// MedianFinder 数据流中找到中位数：（排序之后位于中间位置的元素）
// 295
type MedianFinder struct {
	min utils.MinHeap
	max utils.Heap

	even bool
}

func NewMedianFinder() MedianFinder {
	ret := MedianFinder{
		min:  utils.MinHeap{},
		max:  make(utils.Heap, 0),
		even: true,
	}
	ret.min.Heap = make(utils.Heap, 0)
	return ret
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.min, num)
	heap.Push(&this.max, heap.Pop(&this.min))
	this.even = !this.even
	if this.even {
		heap.Push(&this.min, heap.Pop(&this.max))
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.max.Len() == 0 && this.min.Len() == 0 {
		return 0
	}

	if this.even {
		return float64(this.max.Top()+this.min.Top()) / 2
	}

	return float64(this.max.Top())
}
