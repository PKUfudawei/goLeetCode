package goLeetCode

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
	}
}
