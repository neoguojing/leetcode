package list

/*SwapPairs ...
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

/*RemoveNthFromEnd ...
19 给定一个链表，将倒数第 n 个结点删除
快慢指针
n = 2
 0->1->2->3->4->nil
nth       i
q         p
	  nth        i
	   q         p

**/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	nth := 0
	p := head
	q := head

	for ; p != nil; p = p.Next {
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

/*ReverseKGroup ...
25 将一个链表，每 k 个倒置，最后一组不足 k 个就不倒置
递归：
1.子问题：g(n-1).Next = g(n)
2.结束条件：组元素个数小于k，则不反转直接返回
**/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 0 {
		return head
	}

	if k == 1 {
		head, _ = ReverseList(head)
		return head
	}

	head = reverseKGroup(head, k)

	return head
}

/**

 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var count = 0
	//需要指针做偏移
	p := head
	//找到第n组的尾指针
	for p.Next != nil && count+1 < k {
		p = p.Next
		count++
	}
	//确定本组的尾指针和下一组的头指针
	tail := p
	//编码时要注意空指针
	p = p.Next
	//长度不足直接返回
	if count+1 < k {
		return head
	}
	//断开胃部进行本组反转
	tail.Next = nil
	head, tail = ReverseList(head)
	//进行下一组反转
	if p != nil {
		tail.Next = reverseKGroup(p, k)
	}

	return head
}

/*ReverseList ...
a ->     b -> c -> other
dummy         head
1.需要一个额外指针，保存即将插入的元素
3.head保存接下来将插入的元素
2.提前将末尾指针设置为nil
*/
func ReverseList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	tail := head
	dummyHead := &ListNode{}
	dummyHead.Next = head
	q := head
	head = head.Next
	dummyHead.Next.Next = nil
	for head != nil {
		q = head
		head = head.Next
		q.Next = dummyHead.Next
		dummyHead.Next = q
	}

	return dummyHead.Next, tail
}

/*RotateRight ...
no 61
将最后一个链表节点移到最前边，然后重复这个过程 k 次。
快慢指针：如何找到目标指针?
1.倒转倒数k个，需要找出倒数第k+1个
2.
不满足k的，不处理
k=3
3    2    1    0
0 -> 1 -> 2 -> 3 -> 4
			   k
head
					k
    head
*/
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head
	p := head

	for p.Next != nil {
		if k <= 0 {
			head = head.Next
		}
		k--
		p = p.Next
	}

	//长度不足，则返回
	if k-1 > 0 {
		return head
	}

	q := head.Next
	head.Next = nil

	h, t := ReverseList(q)
	t.Next = dummyHead.Next
	dummyHead.Next = h

	return dummyHead.Next
}

/*DeleteDuplicates ...
no 83
列表有序
给定一个链表，去重，每个数字只保留一个
*/
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	p := head.Next

	for p != nil {
		if p.Val != head.Val {
			p = p.Next
			head = head.Next
		} else {
			head.Next = p.Next
			p = head.Next
		}

	}
	return dummy.Next
}
