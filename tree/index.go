package tree

import "fmt"

//TreeNode ...
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   interface{}
}

//Print ...
func (l *TreeNode) Print(root *TreeNode) {
	out := l.ToString(root)

	fmt.Println(out)
}

//ToString ...
/*
0          1          7
1	   2       3      3 7
2	 4   5   6   7    1 3
3	1 1 1 1 1 1 1 1   0 1     2^3
*/
func (l *TreeNode) ToString(root *TreeNode) string {
	out := "\n"
	ret := BFSByRow(root)
	if ret == nil {
		return out
	}

	dep := len(ret)
	for i, row := range ret {
		headSpaceNum := (1 << (dep - 1 - i)) - 1
		headSpace := ""
		for i := 0; i < headSpaceNum; i++ {
			headSpace += " "
		}
		neibSpaceNum := (1 << (dep - i)) - 1
		neibSpace := ""
		for i := 0; i < neibSpaceNum; i++ {
			neibSpace += " "
		}
		out += headSpace
		for _, elem := range row {

			if elem == 0 {
				out += fmt.Sprintf(" ")
			} else {
				out += fmt.Sprintf("%d", elem)
			}
			out += neibSpace
		}

		out += "\n"
	}
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
