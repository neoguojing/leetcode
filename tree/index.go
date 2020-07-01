package tree

import (
	"fmt"
	"leetcode/utils"
)

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
		headSpace := fmt.Sprintf("% [1]*s", 2*headSpaceNum, " ")
		// for i := 0; i < headSpaceNum; i++ {
		// headSpace += " "
		// }
		neibSpaceNum := (1 << (dep - i)) - 1
		neibSpace := fmt.Sprintf("% [1]*s", 2*neibSpaceNum, " ")
		/*for i := 0; i < neibSpaceNum; i++ {
			neibSpace += " "
		}*/
		out += headSpace
		for _, elem := range row {

			if elem == nil {
				out += fmt.Sprintf("  ")
			} else {
				out += fmt.Sprintf("%2v", elem)
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

// BFSByRow ...
func BFSByRow(root *TreeNode) [][]interface{} {
	if root == nil {
		return nil
	}

	result := make([][]interface{}, 0)
	queue := utils.NewQueue()
	queue.Push(root)
	left := 0
	right := 1
	level := 1
	row := make([]interface{}, 0)
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
				if row[i] != nil {
					isLeaf = false
				}
			}
			if isLeaf {
				break
			}

			result = append(result, row)
			left = 0
			right = queue.Len()
			row = make([]interface{}, 0)
		}

	}

	return result
}
