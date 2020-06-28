package tree

/*
所谓二分查找树，定义如下：

若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
任意节点的左、右子树也分别为二叉查找树；
没有键值相等的节点。
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
