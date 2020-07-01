package tree

import (
	"leetcode/utils"
)

/*
	1.树的递归，可以转换为栈的方式
*/

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
		op(root)
		return
	}

	if op != nil {
		op(root)
	}

	PreOrder(root.Left, op)
	PreOrder(root.Right, op)
}

// InOrder ...
func InOrder(root *TreeNode, op func(*TreeNode)) {
	if root == nil {
		return
	}

	InOrder(root.Left, op)
	if op != nil {
		op(root)
	}
	InOrder(root.Right, op)
}

// PostOrder ...
func PostOrder(root *TreeNode, op func(*TreeNode)) {
	if root == nil {
		return
	}

	PostOrder(root.Left, op)
	PostOrder(root.Right, op)
	if op != nil {
		op(root)
	}
}

// BFS ...
func BFS(root *TreeNode, op func(*TreeNode)) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	queue := utils.NewQueue()
	queue.Push(root)
	left := 0
	right := 1
	row := make([]int, 0)
	for !queue.Empty() {
		node := queue.Pop().(*TreeNode)
		if op != nil {
			op(node)
		}
		row = append(row, node.Val)

		left++

		if node.Left != nil {
			queue.Push(node.Left)
		}

		if node.Right != nil {
			queue.Push(node.Right)
		}

		if left == right {
			ret = append(ret, row)
			left = 0
			right = queue.Len()
			row = make([]int, 0)
		}
	}

	return ret
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
	level := 1
	row := make([]int, 0)
	for !queue.Empty() {
		node := queue.Pop().(*TreeNode)

		row = append(row, node.Val)
		left++

		if node.Left != nil {
			queue.Push(node.Left)
		} else {
			nilNode := &TreeNode{}
			queue.Push(nilNode)
		}

		if node.Right != nil {
			queue.Push(node.Right)
		} else {
			nilNode := &TreeNode{}
			queue.Push(nilNode)
		}

		if left == right {
			level++
			//判断是否叶节点
			isLeaf := true
			for i := range row {
				if row[i] != 0 {
					isLeaf = false
				}
			}
			if isLeaf {
				break
			}

			result = append(result, row)
			left = 0
			right = queue.Len()
			row = make([]int, 0)
		}

	}

	return result
}

// Depth ...
// 求二叉树的深度
func Depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDep := Depth(root.Left) + 1
	rightDep := Depth(root.Right) + 1

	if leftDep > rightDep {
		return leftDep
	}

	return rightDep
}

// IsSymmetric ...
// no 101
// 判断一个二叉树是否关于中心轴对称。
// 解法：递归，同时递归两个树
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil && right != nil {
		return false
	}

	if left != nil && right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetric(left.Left, right.Right) && isSymmetric(left.Right, right.Left)

}

//BuildTree ...
// no 105
// 根据二叉树的先序遍历和中序遍历还原二叉树
// 3,9,20,15,7
// 9 3 15 20 7
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	return buildTree(preorder, 0, inorder)
}

func buildTree(preorder []int, idx int, inorder []int) *TreeNode {
	if idx >= len(preorder) {
		return nil
	}

	root := &TreeNode{}
	root.Val = preorder[idx]
	mid := 0
	for i, v := range inorder {
		if v == preorder[idx] {
			mid = i
		}
	}

	left := inorder[0:mid]
	if len(left) != 0 {
		root.Left = buildTree(preorder, idx+1, left)
	}

	right := inorder[mid+1:]
	if len(right) != 0 {
		root.Right = buildTree(preorder, idx+2, right)
	}

	return root
}

//BuildTreeByPost ...
// no 106
// 根据二叉树的中序遍历和后序遍历还原二叉树。
// 3,9,20,15,7
// 9 3 15 20 7
func BuildTreeByPost(postorder []int, inorder []int) *TreeNode {
	if postorder == nil || inorder == nil {
		return nil
	}

	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	return buildTreeByPost(postorder, len(postorder)-1, inorder)
}

func buildTreeByPost(postorder []int, idx int, inorder []int) *TreeNode {
	if idx < 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = postorder[idx]
	mid := 0
	for i, v := range inorder {
		if v == postorder[idx] {
			mid = i
		}
	}

	right := inorder[mid+1:]
	if len(right) != 0 {
		root.Right = buildTreeByPost(postorder, idx-1, right)
	}

	left := inorder[0:mid]
	if len(left) != 0 {
		root.Left = buildTreeByPost(postorder, idx-2, left)
	}

	return root
}
