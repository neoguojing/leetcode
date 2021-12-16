package tree

import "leetcode/utils"

type Node struct {
	Val      int
	Children []*Node
}

// n叉树的前序遍历
// 589
// [1,null,3,2,4,null,5,6]
func NryPreorderWithStack(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack := utils.NewStack()
	ret := []int{}
	stack.Push(root)
	for !stack.Empty() {
		tmp := stack.Pop().(*Node)
		ret = append(ret, tmp.Val)

		for i := len(tmp.Children) - 1; i >= 0; i-- {
			stack.Push(tmp.Children[i])
		}
	}

	return ret
}
func NryPreorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{root.Val}

	for _, v := range root.Children {
		ret = append(ret, NryPreorder(v)...)
	}

	return ret
}

// n叉树后序遍历
// 590
func NryPostorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}

	for _, v := range root.Children {
		ret = append(ret, NryPostorder(v)...)
	}

	ret = append(ret, root.Val)

	return ret
}

// LevelOrder 层序遍历
//429
func LevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := utils.NewQueue()
	q.Push(root)
	ret := [][]int{}
	for !q.Empty() {
		row := []int{}
		level := []*Node{}
		for !q.Empty() {
			tmp := q.Pop().(*Node)
			row = append(row, tmp.Val)
			level = append(level, tmp.Children...)
		}
		for _, v := range level {
			q.Push(v)
		}
		ret = append(ret, row)
	}

	return ret
}
