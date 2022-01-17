package tree

import "math"

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
		return
	}

	if root.Left == nil && root.Right == nil && sum == 0 {
		tmp = append(tmp, root.Val.(int))
		tmp1 := make([]int, len(tmp))
		copy(tmp1, tmp)
		*ret = append(*ret, tmp1)
		return
	}

	tmp = append(tmp, root.Val.(int))
	pathSumHelper(root.Left, sum-root.Val.(int), tmp, ret)
	pathSumHelper(root.Right, sum-root.Val.(int), tmp, ret)
}

// PathSumIII 路径和等于targetSum，路径可以不经过根，且路径方向朝下,返回路径个数
// 437
func PathSumIII(root *TreeNode, targetSum int) int {
	ret := 0
	pathSumIIIDFS(root, targetSum, &ret)
	return ret
}

func pathSumIIIDFS(root *TreeNode, targetSum int, cnt *int) {
	if root == nil {
		return
	}

	pathSumIIIHelper(root, targetSum, cnt)
	pathSumIIIDFS(root.Left, targetSum, cnt)
	pathSumIIIDFS(root.Right, targetSum, cnt)
}

func pathSumIIIHelper(root *TreeNode, targetSum int, cnt *int) {
	if root == nil {
		return
	}

	if root.Val.(int) == targetSum {
		*cnt++
	}

	pathSumIIIHelper(root.Left, targetSum-root.Val.(int), cnt)
	pathSumIIIHelper(root.Right, targetSum-root.Val.(int), cnt)

}

// MaxPathSum二叉树中的权重最长的路径 路径不一定经过根
// 124
func MaxPathSum(root *TreeNode) int {
	maxPath := math.MinInt32
	getMax(root, &maxPath)
	return maxPath
}

func getMax(node *TreeNode, maxPath *int) int {
	if node == nil {
		return 0
	}

	leftMax := getMax(node.Left, maxPath)
	if leftMax < 0 {
		leftMax = 0
	}

	rightMax := getMax(node.Right, maxPath)
	if rightMax < 0 {
		rightMax = 0
	}

	curMax := node.Val.(int) + leftMax + rightMax
	if curMax > *maxPath {
		*maxPath = curMax
	}

	if rightMax > leftMax {
		leftMax = rightMax
	}
	return node.Val.(int) + leftMax //返回上层时只能保留左右子树中的最大的一个，否则不是路径
}

// DiameterOfBinaryTree 求二叉树的最长路径（直径）
// 543
func DiameterOfBinaryTree(root *TreeNode) int {

	max := 0
	diameterOfBinaryTreeHelper(root, &max)
	if max > 0 {
		return max - 1
	}
	return 0
}

func diameterOfBinaryTreeHelper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := diameterOfBinaryTreeHelper(root.Left, max)

	right := diameterOfBinaryTreeHelper(root.Right, max)

	curLen := left + right + 1
	if curLen > *max {
		*max = curLen
	}

	if right > left {
		left = right
	}

	return left + 1
}
