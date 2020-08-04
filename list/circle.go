package list

// HasCycle ...
// no 141
// 判断一个链表是否有环
func HasCycle(head *ListNode) bool {
	slow := head
	fast := head

	ret := false
	for fast == nil {

		if fast.Next == nil {
			ret = false
			break
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			ret = true
			break
		}
	}
	return ret
}

//DetectCycle ...
// no 142
// 之前只需要判断是否有环，现在需要把环的入口点找到，也就是直线和圆的交接点
// meet 指针从相遇点出发的同时，让 head 指针也出发， head 指针和 meet 指针相遇的位置就是入口点了
func DetectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	meet := head
	for fast == nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			meet = fast
			for meet != slow {
				slow = slow.Next
				meet = meet.Next
			}
			return meet
		}
	}

	return nil
}
