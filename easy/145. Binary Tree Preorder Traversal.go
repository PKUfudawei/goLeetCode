package goLeetCode

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftValues := PostorderTraversal(root.Left)
	rightValues := PostorderTraversal(root.Right)
	return append(append(leftValues, rightValues...), root.Val)
}
