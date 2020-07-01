package test

import (
	"leetcode/list"
	"leetcode/tree"
	"testing"
)

func TestNumTrees(t *testing.T) {
	ret := tree.NumTrees(3)
	t.Log(ret)
}

func TestGenerateBSTrees(t *testing.T) {
	ret := tree.GenerateBSTrees(3)
	for _, v := range ret {
		t.Log(v.ToString(v))
	}
}

func TestIsValidBST(t *testing.T) {
	root := tree.GeneBinaryTree([]int{2, 1, 3}, -1)
	ret := tree.IsValidBST(root)
	t.Log(ret)

	root = tree.GeneBinaryTree([]int{5, 1, 4, -1, -1, 3, 6}, -1)
	ret = tree.IsValidBST(root)
	t.Log(ret)
}

func TestRecoverTree(t *testing.T) {
	root := tree.GeneBinaryTree([]int{1, 3, -1, -1, 2}, -1)
	t.Log(root.ToString(root))
	tree.RecoverTree(root)
	t.Log(root.ToString(root))

	root = tree.GeneBinaryTree([]int{3, 1, 4, -1, -1, 2}, -1)
	t.Log(root.ToString(root))
	tree.RecoverTree(root)
	t.Log(root.ToString(root))

}

func TestSortedArrayToBST(t *testing.T) {
	in := []int{-10, -3, 0, 5, 9}
	ret := tree.SortedArrayToBST(in)
	t.Log(ret.ToString(ret))
}

func TestSortedListToBST(t *testing.T) {
	in := []int{1, 3, 0, 5, 9}
	head := list.GeneListByArray(in)
	ret := tree.SortedListToBST(head)
	t.Log(ret.ToString(ret))
}

// 默认下标从0开始
func TestListIndex(t *testing.T) {
	in := []int{0, 1, 2}
	head := list.GeneListByArray(in)
	cur := head
	end := 0
	//当前指针等于nil的判定，返回的计数：
	// 1是为尾节点的下一个的索引
	// 2.是整个列表元素的个数
	for cur != nil {
		end++
		cur = cur.Next
	}
	t.Log(1, cur)
	t.Log(1, end)
	cur = head
	end = 0
	// 用next作为判定条件，则计数为尾节点的计数
	for cur.Next != nil {
		end++
		cur = cur.Next
	}

	t.Log(2, cur)
	t.Log(2, end)
	cur = head
	end = 0
	//以i=0遍历做计数的，条件用<,总是返回上限对饮的位置的指针和计数
	for i := 0; i < 3; i++ {
		end++
		cur = cur.Next
	}
	t.Log(3, cur)
	t.Log(3, end)
	cur = head
	end = 0

	for i := 0; i < 2; i++ {
		end++
		cur = cur.Next
	}
	t.Log(5, cur)
	t.Log(5, end)
	cur = head
	end = 0
}
