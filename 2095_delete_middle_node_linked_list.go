package main

func deleteMiddle(head *ListNode) *ListNode {
	stackCursor := 0

	current := head
	for current != nil {
		current = current.Next
		stackCursor++
	}

	if stackCursor == 1 {
		return nil
	}

	middle := stackCursor / 2

	stackCursor = 0
	current = head
	var prev *ListNode
	for {
		if stackCursor == middle {
			prev.Next = current.Next
			break
		}

		prev = current
		current = current.Next
		stackCursor++
	}

	return head
}
