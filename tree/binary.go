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

// PreOrderWithStack ...
// 注意压栈顺讯
func PreOrderWithStack(root *TreeNode, op func(*TreeNode)) {
	if root == nil {
		return
	}
	stack := utils.NewStack()
	stack.Push(root)

	for !stack.Empty() {
		v := stack.Pop()
		if v == nil {
			continue
		}
		cur := v.(*TreeNode)

		if op != nil {
			op(cur)
		}
		// 先压右子树
		stack.Push(cur.Right)
		stack.Push(cur.Left)
	}
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

// InOrderWithStack ...
//只有一处压栈处理
func InOrderWithStack(root *TreeNode, op func(*TreeNode)) {
	if root == nil {
		return
	}
	stack := utils.NewStack()
	cur := root
	for !stack.Empty() || cur != nil {
		//左子树全部入栈
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		v := stack.Pop()

		//访问左节点
		cur = v.(*TreeNode)
		if op != nil {
			op(cur)
		}

		cur = cur.Right
	}

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

// PostOrderWithStack ...
// 两次入栈解决问题
// 注意入栈顺序
func PostOrderWithStack(root *TreeNode, op func(*TreeNode)) {
	if root == nil {
		return
	}
	stack := utils.NewStack()
	cur := root
	stack.Push(cur)
	stack.Push(cur)

	for !stack.Empty() {
		v := stack.Pop()
		if v == nil {
			continue
		}

		cur = v.(*TreeNode)

		if !stack.Empty() && cur.Val == stack.Peak().(*TreeNode) {
			stack.Push(cur.Right)
			stack.Push(cur.Right)

			stack.Push(cur.Left)
			stack.Push(cur.Left)
		} else {
			if op != nil {
				op(cur)
			}
		}
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
		row = append(row, node.Val.(int))

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

// MaxDepth ...
// 求二叉树的深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDep := MaxDepth(root.Left) + 1
	rightDep := MaxDepth(root.Right) + 1

	if leftDep > rightDep {
		return leftDep
	}

	return rightDep
}

// MinDepth ...
// no 111
// 求二叉树的深度
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDep, rightDep := 65535, 65535
	if root.Left != nil {
		leftDep = MinDepth(root.Left) + 1
	}

	if root.Right != nil {
		rightDep = MinDepth(root.Right) + 1
	}

	if leftDep > rightDep {
		return rightDep
	}

	return leftDep
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
	root := &TreeNode{}
	root.Val = preorder[0]
	mid := 0
	for i, v := range inorder {
		if v == preorder[0] {
			mid = i
			break
		}
	}

	root.Left = BuildTree(preorder[1:mid+1], inorder[0:mid])

	root.Right = BuildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

//BuildTreeByPost ...
// no 106
// 根据二叉树的中序遍历和后序遍历还原二叉树。
// 3,9,20,15,7
// 9 3 15 20 7
func BuildTreeByPost(inorder []int, postorder []int) *TreeNode {
	if postorder == nil || inorder == nil {
		return nil
	}

	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{}
	root.Val = postorder[len(postorder)-1]
	mid := len(postorder) - 1
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			mid = i
			break
		}
	}

	root.Left = BuildTreeByPost(inorder[:mid], postorder[:mid])
	root.Right = BuildTreeByPost(inorder[mid+1:], postorder[mid:len(postorder)-1])
	return root
}

//Flatten ...
// no 114
// 题目其实就是将二叉树通过右指针，组成一个链表。
//将左子树插入到右子树的地方
// 将原来的右子树接到左子树的最右边节点
// 考虑新的右子树的根节点，一直重复上边的过程，直到新的右子树为 null
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	tmp := root.Right

	root.Right = root.Left
	root.Left = nil

	var rChild *TreeNode = root
	for rChild.Right != nil {
		rChild = rChild.Right
	}

	rChild.Right = tmp

	Flatten(root.Right)
}

//Connect ...
//no 116 117 ,区别1个是满秩二叉树，一个不是
//二叉树，每个节点多了一个next指针，然后将所有的next指针指向它的右边的节点。并且要求空间复杂度是O(1)。
/*
			1
		2       3
	4      5 6      7
*/
func Connect(root *TreeNode) *TreeNode {

	// 指向本层
	cur := root
	for cur != nil {
		// 下层的虚拟头
		dummyHead := &TreeNode{}
		tail := dummyHead

		for cur != nil {
			if cur.Left != nil {
				tail.Next = cur.Left
				tail = tail.Next
			}

			if cur.Right != nil {
				tail.Next = cur.Right
				tail = tail.Next
			}

			cur = cur.Next
		}
		// 虚拟头的下一个节点，为下层的第一个节点
		cur = dummyHead.Next
	}
	return root
}

//ConnectForFullBT ...
// 仅适用于满二叉树
func ConnectForFullBT(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 保存每层的第一个节点
	pre := root
	// cur保存要处理层的上层节点
	var cur *TreeNode
	for pre.Left != nil {
		// cur 站在上一级
		cur = pre
		for cur != nil {
			// 子树的左孩子指向右孩子
			cur.Left.Next = cur.Right
			// 子树的右孩子指向兄弟节点的左孩子
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			// 处理兄弟节点
			cur = cur.Next
		}
		//指向下一层
		pre = pre.Left
	}

	return root
}

// RightSideView ...
// no 199
//给一个二叉树，然后想象自己站在二叉树右边向左边看过去，返回从上到下看到的数字序列。
// 解法1:层序遍历，输出每行最右边的
// 解法2:递归，如下
func RightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	rightSideView(root, 0, &ret)
	return ret
}

func rightSideView(root *TreeNode, level int, ret *[]int) {
	if root == nil {
		return
	}

	//res.size() 的值理解成当前在等待的层级数
	//res.size() == 0, 在等待 level = 0 的第一个数
	//res.size() == 1, 在等待 level = 1 的第一个数
	//res.size() == 2, 在等待 level = 2 的第一个数
	if level == len(*ret) {
		*ret = append(*ret, root.Val.(int))
	}

	// 先遍历右子树
	rightSideView(root.Right, level+1, ret)
	rightSideView(root.Left, level+1, ret)
}
