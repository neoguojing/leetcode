package list

//AddTwoNumbers ...
//no 2
//carry变量
//空头节点的应用
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	pos := head
	carry := 0
	sum := 0
	cur := 0
	for l1 != nil || l2 != nil {

		if l1 == nil && l2 != nil {
			sum = l2.Val + carry
			l2 = l2.Next
		}

		if l2 == nil && l1 != nil {
			sum = l1.Val + carry
			l1 = l1.Next
		}

		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		}

		cur = sum % 10
		carry = sum / 10

		pos.Next = &ListNode{}
		pos.Next.Val = cur
		pos = pos.Next

	}

	if carry != 0 {
		pos.Next = &ListNode{}
		pos.Next.Val = carry
	}

	return head.Next
}

//递归解法
func AddTwoNumbersRecur(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func addTwoNumbersRecur(l1 *ListNode, l2 *ListNode, carry int, output *ListNode) {
	if l1 == nil {

	}
}

/*MergeTwoLists ...
no 21
合并两个有序链表。

1.首尾比较，一次合并
2.递归结束条件：
**/
func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	mergeTwoLists(l1, l2, cur)
	return dummyHead.Next
}
func mergeTwoLists(l1, l2, cur *ListNode) {
	if l1 == nil && l2 == nil {
		return
	}
	if l1 == nil && l2 != nil {
		cur.Next = l2
		return
	}

	if l1 != nil && l2 == nil {
		cur.Next = l1
		return
	}

	if l1.Val < l2.Val {
		cur.Next = l1
		mergeTwoLists(l1.Next, l2, cur.Next)
	} else {
		cur.Next = l2
		mergeTwoLists(l1, l2.Next, cur.Next)
	}

}
