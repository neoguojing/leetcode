package tree

import (
	"leetcode/utils"
)

//IsSameTree ...
//no 100
//判断两个二叉树是否相同
//中序遍历
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return inorderTraversal(p, q)
}

func inorderTraversal(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if !inorderTraversal(p.Left, q.Left) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !inorderTraversal(p.Right, q.Right) {
		return false
	}

	return true
}

// PreOrder ...
func PreOrder(root *TreeNode, op func(*TreeNode)) {
	if root == nil {
		return
	}

	if op != nil {
		op(root)
	}

	PreOrder(root.Left, op)
	PreOrder(root.Right, op)
}

// BFS ...
func BFS(root *TreeNode, op func(*TreeNode)) {
	if root == nil {
		return
	}
	queue := utils.NewQueue()
	queue.Push(root)
	for !queue.Empty() {
		node := queue.Pop().(*TreeNode)
		if op != nil {
			op(node)
		}

		if node.Left != nil {
			queue.Push(node.Left)
		}

		if node.Right != nil {
			queue.Push(node.Right)
		}
	}
}

// BFSByRow ...
func BFSByRow(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := utils.NewQueue()
	queue.Push(root)
	left := 0
	right := 1
	row := make([]int, 0)
	for !queue.Empty() {
		node := queue.Pop().(*TreeNode)

		row = append(row, node.Val)
		left++

		if node.Left != nil {
			queue.Push(node.Left)
		}

		if node.Right != nil {
			queue.Push(node.Right)
		}

		if left == right {
			result = append(result, row)
			left = 0
			right = queue.Len()
			row = make([]int, 0)
		}

	}

	return result
}
