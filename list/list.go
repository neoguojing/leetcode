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

/*
19 给定一个链表，将倒数第 n 个结点删除
快慢指针
n = 2
 0->1->2->3->4->nil
nth       i
q         p 
	  nth        i
	   q         p         
	   
**/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode{
	i:=0
	nth := 0
	p := head
	q := head
	
	for ;p != nil;p = p.Next {
		if i-nth == n+1 {
			nth++
			q = q.Next
		}
		i++
	}


	if q.Next != nil {
		q.Next = q.Next.Next
	}
	
	return head
}


/*
25 将一个链表，每 k 个倒置，最后一组不足 k 个就不倒置
**/
func ReverseKGroup(head *ListNode, k int) *ListNode{
	return nil
}

/*
a ->     b -> c -> other

dummy    head
dummy         head
*/
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{}
	for head!= nil {
		dummyHead.Next = head
		head = head.Next
	}

	return dummyHead.Next
}

