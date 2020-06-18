package list

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

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

func Print(head *ListNode) {
	out := ToString(head)

	fmt.Println(out)
}

func ToString(head *ListNode) string {
	cur := head
	out := ""
	for cur != nil {
		out += fmt.Sprintf("%d->", cur.Val)
		cur = cur.Next
	}
	return out
}
