package main

func partition(head *ListNode, x int) *ListNode {
	less := make([]int, 0)
	more := make([]int, 0)

	current := head
	for current != nil {
		if current.Val < x {
			less = append(less, current.Val)
		} else {
			more = append(more, current.Val)
		}

		current = current.Next
	}

	var result *ListNode
	var start *ListNode
	for _, num := range less {
		node := &ListNode{
			Val:  num,
			Next: nil,
		}

		if result == nil {
			result = node
			start = node
		} else {
			result.Next = node
			result = node
		}
	}
	for _, num := range more {
		node := &ListNode{
			Val:  num,
			Next: nil,
		}

		if result == nil {
			result = node
			start = node
		} else {
			result.Next = node
			result = node
		}
	}

	return start
}
