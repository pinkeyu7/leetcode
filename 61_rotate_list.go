package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	listLength := 1
	for current.Next != nil {
		listLength++
		current = current.Next
	}

	k = k % listLength
	if k == 0 {
		return head
	}

	shift := listLength - k
	cursor := head
	for shift > 0 {
		shift--
		if shift == 0 {
			temp := cursor.Next
			cursor.Next = nil
			cursor = temp
		} else {
			cursor = cursor.Next
		}
	}
	current.Next = head

	return cursor
}
