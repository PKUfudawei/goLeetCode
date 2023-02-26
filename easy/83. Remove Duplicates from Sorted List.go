package goLeetCode

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next_head := DeleteDuplicates(head.Next)
	if head.Val == next_head.Val {
		head.Next = next_head.Next
	}
	return head
}
