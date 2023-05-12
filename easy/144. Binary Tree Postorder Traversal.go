package goLeetCode

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftValues := PreorderTraversal(root.Left)
	rightValues := PreorderTraversal(root.Right)
	return append(append([]int{root.Val}, leftValues...), rightValues...)
}
