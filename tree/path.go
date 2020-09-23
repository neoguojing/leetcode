package tree

//HasPathSum ...
// no 112
// 给定一个sum，判断是否有一条从根节点到叶子节点的路径，该路径上所有数字的和等于sum。
func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val.(int) == sum
	}

	return HasPathSum(root.Left, sum-root.Val.(int)) || HasPathSum(root.Right, sum-root.Val.(int))
}

//PathSum ...
// no 113
// 给定一个sum，输出一条从根节点到叶子节点的路径，该路径上所有数字的和等于sum。
func PathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	tmp := make([]int, 0)

	pathSum(root, sum, tmp, &ret)
	return ret
}

func pathSum(root *TreeNode, sum int, tmp []int, ret *[][]int) {
	if root == nil {
		return
	}

	sum = sum - root.Val.(int)
	tmp = append(tmp, root.Val.(int))
	//处理合法的叶节点和非合法的叶节点
	if root.Left == nil && root.Right == nil && sum == 0 {
		dst := make([]int, len(tmp))
		copy(dst, tmp)
		*ret = append(*ret, tmp)
		return
	} else if root.Left == nil && root.Right == nil && sum != 0 {
		return
	}

	pathSum(root.Left, sum, tmp, ret)
	pathSum(root.Right, sum, tmp, ret)
}
