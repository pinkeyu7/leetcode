package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var cursor *ListNode

	current := head
	for current != nil {
		pointer := &ListNode{
			Val:  current.Val,
			Next: cursor,
		}
		cursor = pointer
		current = current.Next
	}

	return cursor
}
