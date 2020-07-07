package test

import (
	"leetcode/list"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := list.NewLRU(2)
	lru.Put(1, 1)
	lru.Print()
	lru.Put(2, 2)
	lru.Print()
	t.Log(lru.Get(1))
	lru.Print()
	lru.Put(3, 3)
	lru.Print()
	t.Log(lru.Get(2))
	lru.Print()
	lru.Put(4, 4)
	lru.Print()
	t.Log(lru.Get(1))
	lru.Print()
	t.Log(lru.Get(3))
	lru.Print()
	t.Log(lru.Get(4))
	//lru.Print()

}
