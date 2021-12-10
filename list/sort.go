package list

// no 128 列表排序 o(nlogn) o(1)
// 列表排序，一般用归并排序
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	a := SortList(head)
	b := SortList(mid)
	return merge(a, b)
}

func merge(a, b *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for a != nil && b != nil {
		if a.Val < b.Val {
			p.Next = a
			a = a.Next
			p = p.Next
			p.Next = nil
		} else {
			p.Next = b
			b = b.Next
			p = p.Next
			p.Next = nil
		}
	}

	if a != nil {
		p.Next = a
	}

	if b != nil {
		p.Next = b
	}

	return head.Next
}

func getMid(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy
	for p != nil && p.Next != nil {
		p = p.Next.Next
		dummy = dummy.Next
	}
	p = dummy.Next
	dummy.Next = nil
	return p
}

// InsertionSortList 列表的插入排序
// no 147
func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	p := head
	head = dummy
	for p != nil {
		for head.Next != nil && head.Next.Val < p.Val {
			head = head.Next
		}
		q := head.Next
		head.Next = p
		p = p.Next
		head.Next.Next = q
		head = dummy

	}

	return dummy.Next
}
