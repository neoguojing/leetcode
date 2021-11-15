package list

func mergeKLists(lists []*ListNode) *ListNode {
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
	if i < len(lists) {
		ret = append(ret, lists[i])
	}
	return mergeKLists(ret)
}

func mergeTwo(a, b *ListNode) *ListNode {
	var ret *ListNode
	if a.Val > b.Val {
		ret = b
	} else {
		ret = a
	}
	for a != nil && b != nil {
		for a != nil && a.Val < b.Val {
			a = a.Next
		}

		a.Next = b

		for b != nil && b.Val < a.Val {
			b = b.Next
		}

		b.Next = a
	}

	if a != nil {
		b.Next = a
	}

	if b != nil {
		a.Next = b
	}
	return ret
}
