package list

// MergeKLists ...
// no 23 合并排序几个列表
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	ret := make([]*ListNode, 0)
	i := 1
	for ; i < len(lists); i += 2 {
		ret = append(ret, mergeTwo(lists[i-1], lists[i]))
	}
	if i == len(lists) {
		ret = append(ret, lists[i])
	}
	return MergeKLists(ret)
}

// mergeTwo ...
func mergeTwo(a, b *ListNode) *ListNode {

	if a == nil {
		return b
	}

	if b == nil {
		return a
	}
	var ret *ListNode
	if a.Val.(int) > b.Val.(int) {
		ret = b
	} else {
		ret = a
	}
	var prevA, prevB *ListNode
	for a != nil && b != nil {
		for a != nil && b != nil && a.Val.(int) <= b.Val.(int) {
			prevA = a
			a = a.Next
		}
		if prevA != nil {
			prevA.Next = b
			prevA = nil
		}

		for b != nil && a != nil && b.Val.(int) <= a.Val.(int) {
			prevB = b
			b = b.Next
		}

		if prevB != nil {
			prevB.Next = a
			prevB = nil
		}
	}

	return ret
}
