package list

/**
24 给定一个链表，然后两两交换链表的位置
递归法
**/

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	next := head.Next
	head.Next = SwapPairs(head.Next.Next)
	next.Next = head
	return next
}

func SwapPairsMy(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	if head == nil {
		return nil
	}
	swapPairs(dummyHead, head, head.Next)

	return dummyHead.Next
}

func swapPairs(p, n1, n2 *ListNode) {
	if n2 == nil {
		p.Next = n1
		return
	}

	n1.Next = n2.Next
	p.Next = n2
	n2.Next = n1

	if n1.Next == nil {
		return
	}

	swapPairs(n1, n1.Next, n1.Next.Next)
}
