package utils

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h Heap) Less(a, b int) bool {
	return h[a] > h[b]
}

func (h *Heap) Push(a interface{}) {
	*h = append(*h, a.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	x := (old)[len(old)-1]
	*h = (old)[:len(old)-1]
	return x
}

func (h Heap) Top() int {
	return h[0]
}

type MinHeap struct {
	Heap
}

func (m MinHeap) Less(a, b int) bool {

	return m.Heap[a] < m.Heap[b]
}
