package stack

import(
  "container/heap"
  "fmt"
)
func getSkyline(buildings [][]int) [][]int {
    heights := make([][]int, 0)
	for _, building := range buildings {
		heights = append(heights, []int{building[0], -building[2]})
		heights = append(heights, []int{building[1], building[2]})
	}

	sort.Sort(Dim2Array(heights))
    // fmt.Println(heights)
    pq := make(PriorityQueue,0)
	heap.Push(&pq, 0)
	prev := 0
	ret := [][]int{}
	for i := range heights {
		if heights[i][1] < 0 {
			heap.Push(&pq, -heights[i][1])
            // fmt.Println("add:",-heights[i][1])
             // pq.Print()
		} else {
			pq.Remove(heights[i][1])
            // fmt.Println("remove:",heights[i][1])
            // pq.Print()
		}
        
        
        
		cur := pq[0]
        
		if prev != cur {
			ret = append(ret, []int{heights[i][0], cur})
            // fmt.Println("one:",[]int{heights[i][0], cur})
			prev = cur
		}
	}
	return ret
}


type Dim2Array [][]int

func (d2 Dim2Array) Len() int {
	return len(d2)
}

func (d2 Dim2Array) Less(i, j int) bool {
	if d2[i][0] != d2[j][0] {//x不相等，则比较x
		return d2[i][0] < d2[j][0]
	}
	return d2[i][1] < d2[j][1] //相等则比较h
}

func (d2 Dim2Array) Swap(i, j int) {
	d2[i], d2[j] = d2[j], d2[i]
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

// 删除第一个遇到的相同元素
func (pq *PriorityQueue) Remove(val int) {
	idx := -1
	if (*pq)[len(*pq)-1] == val {
		idx = len(*pq) - 1
	} else {
		for i := range *pq {
			if (*pq)[i] == val {
				idx = i
				break
			}
		}
	}

	if idx >= 0 {
		heap.Remove(pq, idx)
	}
}

func (pq PriorityQueue) Print() {
    fmt.Println(pq)
}

