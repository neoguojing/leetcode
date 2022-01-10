package list

import "fmt"

/*RemoveNthFromEnd ...
no 19 给定一个链表，将倒数第 n 个结点删除
快慢指针
n = 2
 0->1->2->3->4->nil
nth       i
q         p
	  nth        i
	   q         p

**/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var p, q *ListNode
	p, q = head, head
	cnt := 1
	for p != nil {
		if cnt > n+1 {
			q = q.Next
		}

		p = p.Next
		cnt++
	}

	// 正好n个元素，移除开头
	if cnt == n+1 {
		head = head.Next
		return head
	}

	// 小于n，返回nil
	if cnt < n {
		return nil
	}
	// q指向第n个元素的前一个元素
	q.Next = q.Next.Next
	return head
}

/*ReverseKGroup ...
no 25 将一个链表，每 k 个倒置，最后一组不足 k 个就不倒置
递归：
1.子问题：g(n-1).Next = g(n)
2.结束条件：组元素个数小于k，则不反转直接返回
**/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 0 || k == 1 {
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
no 206
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
双指针
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

/*DeleteAllDuplicates ...
no 82
列表有序
给一个链表，如果一个数属于重复数字，就把这个数删除，一个都不留
双指针
*/
func DeleteAllDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	//指想前一个节点
	p := dummy
	head = head.Next
	for head != nil {
		if p.Next.Val != head.Val {
			p = p.Next
			head = head.Next
		} else {
			//遍历所有相同的
			for p.Next.Val == head.Val && head != nil {
				head = head.Next
			}
			p.Next = head
			if head != nil {
				head = head.Next
			}
		}

	}
	return dummy.Next
}

/*Partition ...
no 86
将链表分成了两部分，一部分的数字全部小于分区点 x，另一部分全部大于等于分区点 x。最后就是 1 2 2 和 4 3 5 两部分。
双指针
*/
func Partition(head *ListNode, x int) *ListNode {
	//两个头指针，两个活动指针
	dummyLeft, dummyRight := &ListNode{}, &ListNode{}
	p, q := dummyLeft, dummyRight
	for head != nil {
		if head.Val.(int) < x {
			p.Next = head
			p = p.Next
		} else {
			q.Next = head
			q = q.Next
		}
		head = head.Next
	}
	//切断尾防止环路
	q.Next = nil
	//连接两部分列表
	p.Next = dummyRight.Next
	return dummyLeft.Next
}

// ReorderList ...
// no 143
// 给一个链表，然后依次头尾头尾头尾取元素，组成新的链表。
// 1.平分列表
// 2.第二个列表倒序
// 3.依次取值
func ReorderList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head
	fast := head

	//  slow 指向中间位置的前一个指针
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second, _ := ReverseList(slow.Next)
	slow.Next = nil
	fmt.Println(second.ToString(second))
	newHaed := head
	// 交替指向
	for second != nil {
		fmt.Println(head.Val)

		tmp := second.Next
		second.Next = head.Next
		head.Next = second
		head = second.Next
		second = tmp
	}

	return newHaed
}

// OddEvenList o(1)space o(n)time
// 328
func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// NestedIterator 实现迭代器，将[[1,1],2,[1,1]] 平铺
// 341
type NestedInteger struct {
}

type NestedIterator struct {
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return nil
}

func (this *NestedIterator) Next() int {
	return 0
}

func (this *NestedIterator) HasNext() bool {
	return false
}
