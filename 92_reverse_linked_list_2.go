package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left != 1 {
		head.Next = reverseBetween(head.Next, left-1, right-1)
		return head
	} else {
		current := head
		previous := &ListNode{}

		count := 0

		for count < right {
			next := current.Next
			current.Next = previous
			previous = current
			current = next
			count++
		}
		head.Next = current

		return previous
	}
}
