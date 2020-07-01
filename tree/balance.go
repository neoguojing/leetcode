package tree

/*
平衡树：
它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。
*/

// IsBalanced ...
// 判断一棵树是否是平衡二叉树，平衡二叉树定义如下：
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return false
}
