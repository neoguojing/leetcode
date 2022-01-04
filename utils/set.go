package utils

import "math/rand"

// 380 插入删除和获取随机，O（1）
type RandomizedSet struct {
	set map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
		arr: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.set[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	n := len(this.arr)
	idx, ok := this.set[val]
	if !ok {
		return false
	}

	this.set[this.arr[n-1]] = idx
	this.arr[idx], this.arr[n-1] = this.arr[n-1], this.arr[idx]
	this.arr = this.arr[:n-1]
	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
