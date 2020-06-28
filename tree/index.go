package tree

import "fmt"

//TreeNode ...
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

//Print ...
func (l *TreeNode) Print(root *TreeNode) {
	out := l.ToString(root)

	fmt.Println(out)
}

//ToString ...
func (l *TreeNode) ToString(root *TreeNode) []int {
	out := make([]int, 0)
	op := func(node *TreeNode) {
		if node != nil {
			out = append(out, node.Val)

		} else {
			out = append(out, -1)
		}
	}

	PreOrder(root, op)
	return out
}

//Copy ...
func (l *TreeNode) Copy(root *TreeNode) *TreeNode {

	return l.copy(root)
}

func (l *TreeNode) copy(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	copyRoot := &TreeNode{}
	copyRoot.Val = root.Val

	copyRoot.Left = l.copy(root.Left)
	copyRoot.Right = l.copy(root.Right)

	return copyRoot
}

//GeneBinaryTree ...
//以数组表示二叉树，左子节点为2*index+1，父节点为(index-1)/2
// tag 为nil节点的标示
func GeneBinaryTree(in []int, tag int) *TreeNode {
	return geneBinaryTree(in, 0, tag)
}

func geneBinaryTree(in []int, index int, tag int) *TreeNode {

	if index >= len(in) {
		return nil
	}

	if in[index] == tag {
		return nil
	}

	node := &TreeNode{}
	node.Val = in[index]
	node.Left = geneBinaryTree(in, 2*index+1, tag)
	node.Right = geneBinaryTree(in, 2*index+2, tag)

	return node
}
