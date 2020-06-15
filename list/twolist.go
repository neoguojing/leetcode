package list

import(
	"fmt"
)

type ListNode struct {
	Next *ListNode
	Val int
}

func GeneListByArray(in []int) *ListNode{
	head := &ListNode{}
	cur := head
	for _,v := range in {
		cur.Next = &ListNode{
			Val:v,
		}
		cur = cur.Next
	}

	return head.Next
}

func Print(head *ListNode) {
	cur := head
	out := ""
	for cur != nil {
		out += fmt.Sprintf("%d->",cur.Val)
		cur = cur.Next
	}

	fmt.Println(out)
}


//no 2
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	pos := head
	carry := 0
	sum := 0
	cur := 0
    for;l1 != nil || l2 != nil; {
		if l1 == nil {
			sum = l2.Val + carry
			l2 = l2.Next
		}

		if l2 == nil {
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


func AddTwoNumbersRecur(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func addTwoNumbersRecur(l1 *ListNode, l2 *ListNode,carry int,output *ListNode){
	if l1 == nil {

	}
}