package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	total := 0
	var start *ListNode
	var current *ListNode
	var cursor *ListNode
	current = head

	// Return nil if length = 1
	if current.Next == nil {
		return nil
	}

	// Get total amount
	for current != nil {
		total++
		current = current.Next
	}

	// Remove nth
	target := total - n + 1
	index := 0
	current = head
	for current != nil {
		index++
		if index != target {
			if start == nil {
				start = current
				cursor = current
			} else {
				cursor.Next = current
				cursor = cursor.Next
			}
		}

		current = current.Next
	}

	if index == target {
		cursor.Next = nil
	}

	return start
}
