package list

import "fmt"

// PriQueue ...
type PriQueue struct {
	Next *PriQueue
	Pre  *PriQueue
	Val  interface{}
	Key  interface{}
}

// LRU ...
// no 146
type LRU struct {
	queue    *PriQueue
	tail     *PriQueue
	hashMap  map[interface{}]*PriQueue
	capacity int
}

// NewLRU ...
func NewLRU(cap int) *LRU {
	q := &PriQueue{}
	ret := &LRU{
		queue:    q,
		hashMap:  make(map[interface{}]*PriQueue),
		capacity: cap,
		tail:     q,
	}
	return ret
}

// Get ...
// 获取时需要判断key是否存在，存在的话需要将元素移动到列表头部
func (lru *LRU) Get(key interface{}) interface{} {
	var ok bool
	var ret *PriQueue
	if ret, ok = lru.hashMap[key]; ok {
		ret.Pre.Next = ret.Next
		if ret.Next != nil {
			ret.Next.Pre = ret.Pre
		}

		ret.Next = lru.queue.Next
		ret.Pre = lru.queue

		lru.queue.Next = ret

	}
	if ret == nil {
		return -1
	}
	return ret.Val
}

// Put ...
// 放置元素时,如key已经存在，在需要判断容量是否已经满了,满了移除末尾的元素和map中元素，否则添加到末尾
func (lru *LRU) Put(key interface{}, val interface{}) {
	if v, ok := lru.hashMap[key]; ok {
		v.Val = val
	} else {
		node := &PriQueue{}
		node.Val = val
		node.Key = key

		if lru.capacity == len(lru.hashMap) {
			tmp := lru.tail
			lru.tail = lru.tail.Pre
			lru.tail.Next = nil
			tmp.Next = nil
			tmp.Pre = nil
			delete(lru.hashMap, tmp.Key)
		}

		//将node插入末尾
		lru.tail.Next = node
		node.Pre = lru.tail
		lru.tail = lru.tail.Next

		lru.hashMap[key] = node
	}

}

// Size ...
func (lru *LRU) Size() int {
	return len(lru.hashMap)
}

// Print ...
func (lru *LRU) Print() {
	cur := lru.queue
	out := ""
	for cur != nil {
		if cur.Next != nil && cur.Next.Pre != nil {
			out += fmt.Sprintf("(%v,%v)<->", cur.Key, cur.Val)
		} else {
			out += fmt.Sprintf("(%v,%v)->", cur.Key, cur.Val)
		}

		cur = cur.Next
	}
	fmt.Println(out)
}
