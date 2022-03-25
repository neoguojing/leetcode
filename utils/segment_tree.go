package utils

import "fmt"

type SegmentTree []int

type Merge func(int, int) int

func NewSegmentTree(in []int, fn Merge) SegmentTree {
	seg := make(SegmentTree, 4*len(in))
	seg.build(in, 0, 0, len(in)-1, fn)
	return seg
}

func (seg SegmentTree) build(arr []int, v, l, r int, fn Merge) {

	if l == r {
		seg[v] = arr[l]
	} else {
		md := l + (r-l)/2
		seg.build(arr, v*2+1, l, md, fn)
		seg.build(arr, v*2+2, md+1, r, fn)
		seg[v] = fn(seg[v*2+1], seg[v*2+2])
	}
}

func (seg SegmentTree) Query(i, j int, fn Merge) int {
	return seg.query(0, 0, len(seg)/4-1, i, j, fn)
}

func (seg SegmentTree) query(v, l, r int, i, j int, fn Merge) int {
	if i > j || l > j || r < i { //不在任何区间内或者入参非法
		return 0
	}

	if i <= l && j >= r { //正好落在树上区间的情况,或者略大
		return seg[v]
	}

	md := l + (r-l)/2

	if j <= md {
		return seg.query(v*2+1, l, md, i, j, fn)
	} else if i > md {
		return seg.query(v*2+2, md+1, r, i, j, fn)
	}

	left := seg.query(v*2+1, l, md, i, md, fn)
	right := seg.query(v*2+2, md+1, r, md+1, j, fn)
	return fn(left, right)

}

func (seg SegmentTree) Update(pos, val int, fn Merge) {
	fmt.Println(0, len(seg)/4-1)
	seg.update(0, 0, len(seg)/4-1, pos, val, fn)
	fmt.Println("&&&&&&&", seg)
}

func (seg SegmentTree) update(v, l, r int, pos, val int, fn Merge) {
	if l == r {
		seg[l] = val
		fmt.Println(seg[l], l)
		return
	}

	md := l + (r-l)/2
	if pos <= md {
		seg.update(v*2+1, l, md, pos, val, fn)
	} else if pos > md {
		seg.update(v*2+2, md+1, r, pos, val, fn)
	}

	seg[v] = fn(seg[v*2+1], seg[v*2+2])
	fmt.Println(seg[v], v)
}
