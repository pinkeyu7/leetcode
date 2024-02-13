package main

func reorderList(head *ListNode) *ListNode {
	stack := make([]int, 0)
	stackCursor := 0

	current := head
	for current != nil {
		stack = append(stack, current.Val)
		stackCursor++

		current = current.Next
	}

	listLength := stackCursor

	current = head
	for stackCursor > 0 {
		node := &ListNode{
			Val:  stack[stackCursor-1],
			Next: current.Next,
		}
		current.Next = node
		current = node.Next
		stackCursor--
	}

	current = head
	for {
		stackCursor++
		if stackCursor == listLength {
			current.Next = nil
			break
		}
		current = current.Next
	}

	return head
}
