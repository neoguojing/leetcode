package list

import "fmt"

//ListNode ...
type ListNode struct {
	Next *ListNode
	Pre  *ListNode
	Val  interface{}
}

// GeneListByArray ...
func GeneListByArray(in []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, v := range in {
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}

	return head.Next
}

func (l *ListNode) Print(head *ListNode) {
	out := l.ToString(head)

	fmt.Println(out)
}

func (l *ListNode) ToString(head *ListNode) string {
	cur := head
	out := ""
	for cur != nil {
		out += fmt.Sprintf("%v->", cur.Val)
		cur = cur.Next
	}
	return out
}
