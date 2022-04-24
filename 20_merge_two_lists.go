package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var start *ListNode
	var current *ListNode

	cursor1 := list1
	cursor2 := list2

	for cursor1 != nil || cursor2 != nil {
		if cursor1 == nil {
			start, current = addNode(start, current, cursor2.Val)
			cursor2 = cursor2.Next
			continue
		}
		if cursor2 == nil {
			start, current = addNode(start, current, cursor1.Val)
			cursor1 = cursor1.Next
			continue
		}
		if cursor1.Val > cursor2.Val {
			start, current = addNode(start, current, cursor2.Val)
			cursor2 = cursor2.Next
		} else {
			start, current = addNode(start, current, cursor1.Val)
			cursor1 = cursor1.Next
		}
	}

	return start
}

func addNode(start, current *ListNode, number int) (*ListNode, *ListNode) {
	if start == nil {
		start = &ListNode{
			Val:  number,
			Next: nil,
		}
		current = start
	} else {
		current.Next = &ListNode{
			Val:  number,
			Next: nil,
		}
		current = current.Next
	}
	return start, current
}
