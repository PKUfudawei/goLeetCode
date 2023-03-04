package goLeetCode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return &TreeNode{Val: nums[len(nums)/2], Left: SortedArrayToBST(nums[:len(nums)/2]), Right: SortedArrayToBST(nums[len(nums)/2+1:])}
}
