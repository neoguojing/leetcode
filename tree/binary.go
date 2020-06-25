package tree

//IsSameTree ...
//no 100
//判断两个二叉树是否相同
//中序遍历
func IsSameTree( p *TreeNode,  q *TreeNode) bool{
	return inorderTraversal(p,q)
}

func inorderTraversal(p *TreeNode, q *TreeNode) bool{
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if !inorderTraversal(p.Left,q.Left) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !inorderTraversal(p.Right,q.Right) {
		return false
	}

	return true
}