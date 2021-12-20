package tree

import (
	"container/list"
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

// ConstructFromPrePost
// 889
// preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1] Output: [1,2,3,4,5,6,7]
var preIdx, postIdx = 0, 0

func ConstructFromPrePost(preorder []int, postorder []int) *TreeNode {
	root := &TreeNode{}
	root.Val = preorder[preIdx]
	preIdx++
	if root.Val != postorder[postIdx] {
		root.Left = ConstructFromPrePost(preorder, postorder)
	}

	if root.Val != postorder[postIdx] {
		root.Right = ConstructFromPrePost(preorder, postorder)
	}

	postIdx++
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

// InorderTraversal 二叉树 中序遍历
// 94
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := InorderTraversal(root.Left)
	ret = append(ret, root.Val.(int))
	ret = append(ret, InorderTraversal(root.Right)...)

	return ret
}

// PreorderTraversal
// 144
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{root.Val.(int)}
	ret = append(ret, PreorderTraversal(root.Left)...)
	ret = append(ret, PreorderTraversal(root.Right)...)

	return ret
}

// PostorderTraversal
// 145
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := PostorderTraversal(root.Right)
	ret = append(ret, PostorderTraversal(root.Left)...)
	ret = append(ret, root.Val.(int))

	return ret
}

// BSTIterator 实现树的迭代器，中序
// 173
type BSTIterator struct {
	stack *list.List
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		stack: list.New(),
		cur:   root,
	}
	(&iter).pushLeft(root)

	return iter
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		this.stack.PushBack(node)
		node = node.Left
	}

	this.cur = this.stack.Back().Value.(*TreeNode)
}

func (this *BSTIterator) Next() int {
	ret := this.cur.Val
	this.cur = this.stack.Remove(this.stack.Back()).(*TreeNode)
	this.cur = this.cur.Right
	this.pushLeft(this.cur)
	return ret.(int)
}

func (this *BSTIterator) HasNext() bool {
	if this.cur != nil {
		return true
	}

	return false
}

// BinaryLevelOrder 层序遍历
// 102
func BinaryLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := utils.NewQueue()
	q.Push(root)
	ret := [][]int{}
	for !q.Empty() {
		row := []int{}
		level := []*TreeNode{}
		for !q.Empty() {
			tmp := q.Pop().(*TreeNode)
			row = append(row, tmp.Val.(int))
			if tmp.Left != nil {
				level = append(level, tmp.Left)
			}

			if tmp.Right != nil {
				level = append(level, tmp.Right)
			}

		}
		for _, v := range level {
			q.Push(v)
		}
		ret = append(ret, row)
	}

	return ret
}

// 993
// 637
// ZigzagLevelOrder z型遍历
// 103
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := utils.NewQueue()
	q.Push(root)
	ret := [][]int{}
	zigzag := false
	qSize := q.Len()
	for !q.Empty() {
		row := make([]int, qSize)
		for i := 0; i < qSize; i++ {
			tmp := q.Pop().(*TreeNode)
			if !zigzag {
				row[i] = tmp.Val.(int)
			} else {
				row[qSize-i-1] = tmp.Val.(int)
			}

			if tmp.Left != nil {
				q.Push(tmp.Left)
			}
			if tmp.Right != nil {
				q.Push(tmp.Right)
			}
		}

		zigzag = !zigzag

		ret = append(ret, row)

	}

	return ret
}

// LowestCommonAncestor
// 236
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	type Elem struct {
		Parent *TreeNode
		Level  int
	}
	uSet := map[*TreeNode]*Elem{}
	queue := utils.NewQueue()
	queue.Push(root)

	qSize := queue.Len()
	level := 0
	uSet[root] = &Elem{
		Parent: root,
		Level:  level,
	}
	for !queue.Empty() {
		for i := 0; i < qSize; i++ {
			tmp := queue.Pop().(*TreeNode)

			if tmp.Left != nil {
				queue.Push(tmp.Left)
				uSet[tmp.Left] = &Elem{
					Parent: tmp,
					Level:  level,
				}
			}

			if tmp.Right != nil {
				queue.Push(tmp.Right)
				uSet[tmp.Right] = &Elem{
					Parent: tmp,
					Level:  level,
				}
			}
		}
		level++
		if uSet[p] != nil && uSet[q] != nil {
			break
		}
	}

	for uSet[p].Parent != q && uSet[q].Parent != p && uSet[p].Parent != uSet[q].Parent {
		if uSet[p].Level > uSet[q].Level {
			uSet[p] = uSet[uSet[p].Parent]
		} else if uSet[p].Level < uSet[q].Level {
			uSet[q] = uSet[uSet[q].Parent]
		} else {
			uSet[p] = uSet[uSet[p].Parent]
			uSet[q] = uSet[uSet[q].Parent]
		}
	}

	if uSet[p].Parent == q {
		return q
	}

	if uSet[q].Parent == p {
		return p
	}

	return uSet[p].Parent
}

func LowestCommonAncestorSimple(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestorSimple(root.Left, p, q)
	right := LowestCommonAncestorSimple(root.Right, p, q)

	if left == nil && right == nil {
		return nil
	}

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

// MaxPathSum二叉树中的权重最长的路径 路径不一定经过根
// 124
func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

}

// 2096
// 107

//1022

// 675
// 778
