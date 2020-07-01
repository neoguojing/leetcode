package tree

import (
	"leetcode/list"
	"leetcode/utils"
)

/*
所谓二分查找树，定义如下：

若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
任意节点的左、右子树也分别为二叉查找树；
没有键值相等的节点。

特性1:中序遍历的二叉树输出一定是一个有序数组
*/

// NumTrees ...
// no 96
// 1-n 能够成多少二分查找树
func NumTrees(n int) int {
	return numTrees(1, n)
}

func numTrees(start, end int) int {
	ret := 0
	if start >= end {
		return 1
	}

	for k := start; k <= end; k++ {
		leftTrees := numTrees(start, k-1)
		rightTrees := numTrees(k+1, end)

		ret += leftTrees * rightTrees
	}

	return ret
}

// GenerateBSTrees ...
// no 95
// 1-n 输出所有不同的bst
// 递归 ，
func GenerateBSTrees(n int) []*TreeNode {
	ret := generateBSTrees(1, n)
	return ret
}

func generateBSTrees(start int, end int) []*TreeNode {
	ret := make([]*TreeNode, 0)
	if start > end {
		ret = append(ret, nil)
		return ret
	}

	if start == end {
		root := &TreeNode{}
		root.Val = start
		ret = append(ret, root)
		return ret
	}

	for k := start; k <= end; k++ {
		leftTrees := generateBSTrees(start, k-1)
		rightTrees := generateBSTrees(k+1, end)

		for i := range leftTrees {
			for j := range rightTrees {
				root := &TreeNode{}
				root.Val = k
				root.Left = leftTrees[i]
				root.Right = rightTrees[j]
				ret = append(ret, root)
			}
		}

	}

	return ret
}

//IsValidBST ...
// no 98
// 输入一个树，判断该树是否是合法二分查找树
//误区：分别判断根节点和左子树右子树的大小不对，需要确认左子树的最大值小于根同理又子树的最小值大于根
// 解法一：中序遍历
// 解法二：递归：根节点取值任意，左节点取值范围为（-inf，根的值），同理
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := utils.NewStack()
	var pre *TreeNode = nil

	for root != nil || !stack.Empty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		root = stack.Pop().(*TreeNode)
		if pre != nil && pre.Val.(int) >= root.Val.(int) {
			return false
		}
		pre = root
		root = root.Right
	}
	return true
}

//RecoverTree ...
//no 99
// 一个合法的二分查找树随机交换了两个数的位置，然后让我们恢复二分查找树。不能改变原来的结构，只是改变两个数的位置
// 解法：利用中序遍历的时，顺序的原理，找到数组中逆序的序列：1.一组逆序 2.两组逆序，用第一个逆序的首字符，和第二个逆序的末字符交换
func RecoverTree(root *TreeNode) {
	var first *TreeNode = nil
	var second *TreeNode = nil
	var pre *TreeNode = nil
	op := func(root *TreeNode) {
		if root == nil {
			return
		}

		if pre != nil && pre.Val.(int) > root.Val.(int) {
			if second == nil {
				first = pre
			}
			second = root
		} else {
			pre = root
		}
	}
	InOrder(root, op)

	tmp := first.Val
	first.Val = second.Val
	second.Val = tmp
}

// SortedArrayToBST ...
// no 108
// 给一个升序数组，生成一个平衡二叉搜索树。平衡二叉树定义如下：
// 该序列即中续遍历
// -10,-3,0,5,9
func SortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	start := 0
	end := len(nums) - 1
	mid := start + (end-start)/2

	root := &TreeNode{}
	root.Val = nums[mid]

	left := nums[0:mid]
	root.Left = SortedArrayToBST(left)

	right := nums[mid+1:]
	root.Right = SortedArrayToBST(right)

	return root
}

//SortedListToBST ...
// no 109
// 是一样的，都是给定一个升序序列，然后生成二分平衡查找树。区别在于 108 题给定的是数组，这里给的是链表
//有序列表本质上是数的中序遍历
//0 1 2
func SortedListToBST(head *list.ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	cur := head
	end := 0
	for cur != nil {
		end++
		cur = cur.Next
	}

	return sortedListToBST(head, 0, end-1)
}

// 0 1  0
// 1 1  1
func sortedListToBST(head *list.ListNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	cur := head
	//start 和end分别代表实际的起始和结束位置，而不是位置的下一位
	mid := start + (end-start)/2

	//mid 是全量数据的全局位置，用<号确保走到mid的位置
	for i := 0; i < mid; i++ {
		cur = cur.Next
	}

	left := sortedListToBST(head, start, mid-1)

	root := &TreeNode{}
	root.Val = cur.Val
	root.Left = left

	right := sortedListToBST(head, mid+1, end)
	root.Right = right

	return root
}
