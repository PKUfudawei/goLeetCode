package goLeetCode

import "goLeetCode/structure"

type TreeNode = structure.TreeNode

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := InorderTraversal(root.Left)
	left_middle := append(left, root.Val)
	right := InorderTraversal(root.Right)
	left_middle_right := append(left_middle, right...)
	return left_middle_right
}
