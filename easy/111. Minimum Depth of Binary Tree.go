package goLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMinDepth, rightMinDepth := MinDepth(root.Left), MinDepth(root.Right)
	if root.Left == nil {
		return rightMinDepth + 1
	} else if root.Right == nil {
		return leftMinDepth + 1
	} else {
		if leftMinDepth < rightMinDepth {
			return leftMinDepth + 1
		} else {
			return rightMinDepth + 1
		}
	}
}
