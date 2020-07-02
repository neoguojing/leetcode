package tree

//HasPathSum ...
// no 112
// 给定一个sum，判断是否有一条从根节点到叶子节点的路径，该路径上所有数字的和等于sum。
func HasPathSum(root *TreeNode, sum int) bool {
	// 该条件会导致叶节点重复判断两次
	if sum == 0 && root == nil {
		return true
	}

	if sum != 0 && root == nil {
		return false
	}

	sum = sum - root.Val.(int)

	return HasPathSum(root.Left, sum) || HasPathSum(root.Right, sum)
}

//PathSum ...
// no 113
// 给定一个sum，输出一条从根节点到叶子节点的路径，该路径上所有数字的和等于sum。
func PathSum(root *TreeNode, sum int) [][]int {
	if sum == 0 && root == nil {
		return nil
	}

	if root == nil && sum != 0 {
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
		*ret = append(*ret, tmp)
		return
	} else if root.Left == nil && root.Right == nil && sum != 0 {
		return
	}

	pathSum(root.Left, sum, tmp, ret)
	pathSum(root.Right, sum, tmp, ret)
}