package main

func swapPairs(head *ListNode) *ListNode {
	var start *ListNode
	var cursor *ListNode

	current := head
	buffer := 0
	flag := false

	for current != nil {
		if !flag {
			flag = true
			buffer = current.Val
		} else {
			flag = false
			second := &ListNode{
				Val:  buffer,
				Next: nil,
			}
			first := &ListNode{
				Val:  current.Val,
				Next: second,
			}

			if start == nil {
				start = first
				cursor = second
			} else {
				cursor.Next = first
				cursor = second
			}
		}
		current = current.Next
	}

	if flag {
		if start == nil {
			start = &ListNode{
				Val:  buffer,
				Next: nil,
			}
		} else {
			cursor.Next = &ListNode{
				Val:  buffer,
				Next: nil,
			}
		}
	}

	return start
}
