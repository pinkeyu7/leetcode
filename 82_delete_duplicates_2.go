package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	current := head
	next := head.Next

	if current.Val != next.Val {
		head.Next = deleteDuplicates2(head.Next)
		return head
	}

	for next != nil {
		if current.Val != next.Val {
			return deleteDuplicates2(next)
		}

		current = current.Next
		next = next.Next
	}

	return nil
}
