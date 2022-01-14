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

	pathSumHelper(root, sum, tmp, &ret)
	return ret
}

func pathSumHelper(root *TreeNode, sum int, tmp []int, ret *[][]int) {
	if root == nil {
		if sum == 0 {
			tmp1 := make([]int, len(tmp))
			copy(tmp1, tmp)
			*ret = append(*ret, tmp1)
		}
		return
	}
	tmp = append(tmp, root.Val.(int))
	pathSumHelper(root.Left, sum-root.Val.(int), tmp, ret)
	pathSumHelper(root.Right, sum-root.Val.(int), tmp, ret)
	tmp = tmp[:len(tmp)-1]
}

// PathSumIII 路径和等于targetSum，路径可以不经过根，且路径方向朝下
// 437
func PathSumIII(root *TreeNode, targetSum int) int {
	return 0
}
