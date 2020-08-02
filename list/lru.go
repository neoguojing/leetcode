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
	head *PriQueue
	tail *PriQueue

	hashMap  map[interface{}]*PriQueue
	capacity int
}

// NewLRU ...
func NewLRU(cap int) *LRU {
	q := &PriQueue{}
	t := &PriQueue{}
	// 关键首尾相连
	q.Next = t
	t.Pre = q

	ret := &LRU{
		head:     q,
		hashMap:  make(map[interface{}]*PriQueue),
		capacity: cap,
		tail:     t,
	}
	return ret
}

// Get ...
// 获取时需要判断key是否存在，存在的话需要将元素移动到列表头部
func (lru *LRU) Get(key interface{}) interface{} {
	if v, ok := lru.hashMap[key]; ok {

		// 摘除v
		v.Pre.Next = v.Next
		v.Next.Pre = v.Pre

		// 插入头
		v.Next = lru.head.Next
		v.Pre = lru.head

		lru.head.Next.Pre = v
		lru.head.Next = v
		return v.Val
	}

	return nil
}

// Put ...
// 若已存在则，重置新值，并添加的列表头
// 若不存在，且容量已满，删除列表最后元素，在表头插入元素
// 若不存在，且容量不满，则添加元素到列表头
func (lru *LRU) Put(key interface{}, val interface{}) {
	if v, ok := lru.hashMap[key]; ok {
		v.Val = val
		// 从原位置删除
		v.Next.Pre = v.Pre
		v.Pre.Next = v.Next
		// 插入开头
		v.Pre = lru.head
		v.Next = lru.head.Next

		lru.head.Next.Pre = v
		lru.head.Next = v
	} else {
		if len(lru.hashMap) == lru.capacity {
			if lru.tail.Pre != nil {
				end := lru.tail.Pre
				delete(lru.hashMap, end.Key)

				if end.Pre != nil {
					end.Pre.Next = lru.tail
					lru.tail.Pre = end.Pre
				}
			}
		}

		// 插入开头
		node := &PriQueue{
			Key: key,
			Val: val,
		}

		node.Pre = lru.head
		node.Next = lru.head.Next

		lru.head.Next.Pre = node
		lru.head.Next = node

		lru.hashMap[key] = node
	}
}

// Size ...
func (lru *LRU) Size() int {
	return len(lru.hashMap)
}

// Print ...
func (lru *LRU) Print() {
	cur := lru.head
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
