package goLeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	prev, next := head, head.Next
	for next != nil {
		if next.Val == val {
			prev.Next = next.Next
		} else {
			prev = next
		}
		next = next.Next
	}
	return head
}
