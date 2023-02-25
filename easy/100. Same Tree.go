package goLeetCode

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p.Val == q.Val {
		return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
