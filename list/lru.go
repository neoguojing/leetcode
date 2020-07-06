package list

// LRU ...
// no 146
type LRU struct {
	queue    *ListNode
	hashMap  map[interface{}]*ListNode
	size     int
	capacity int
}

// NewLRU ...
func NewLRU(cap int) *LRU {
	ret := &LRU{
		queue:    &ListNode{},
		hashMap:  make(map[interface{}]*ListNode),
		capacity: cap,
	}
	return ret
}

// Get ...
// 获取时需要判断key是否存在，存在的话需要将元素移动到列表头部
func (lru *LRU) Get(key interface{}) interface{} {
	var ok bool
	var ret *ListNode
	if ret, ok = lru.hashMap[key]; ok {
		ret.Pre.Next = ret.Next
		ret.Next = lru.queue.Next
		lru.queue.Next = ret
	}

	return ret.Val
}

// Put ...
// 放置元素时需要判断容量是否已经满了,满了移除末尾的元素和map中元素，否则添加到末尾
func (lru *LRU) Put(key interface{}, val interface{}) {
	if lru.capacity == lru.size {
		tmp := lru.queue.Pre
		lru.queue.Pre = lru.queue.Pre.Pre
		tmp.Next = nil
		lru.size--
	}

	node := &ListNode{}
	node.Val = val
	lru.hashMap[key] = node
	node.Next = lru.queue
	node.Pre = lru.queue.Pre
	lru.queue.Pre = node
	lru.size++
}
