package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var start *ListNode
	var resultCursor *ListNode
	cursor1 := l1
	cursor2 := l2
	moreThanTen := false

	for cursor1 != nil || cursor2 != nil || moreThanTen {
		tempSum := 0
		if moreThanTen {
			tempSum += 1
		}

		if cursor1 != nil {
			tempSum += cursor1.Val
		}
		if cursor2 != nil {
			tempSum += cursor2.Val
		}

		// handle bigger than ten
		moreThanTen = tempSum/10 == 1

		newNode := &ListNode{
			Val:  tempSum % 10,
			Next: nil,
		}

		// Handle current cursor
		if resultCursor == nil {
			resultCursor = newNode
		} else {
			resultCursor.Next = newNode
			resultCursor = resultCursor.Next
		}

		// Setting start node
		if start == nil {
			start = newNode
		}

		if cursor1 != nil {
			cursor1 = cursor1.Next
		}
		if cursor2 != nil {
			cursor2 = cursor2.Next
		}
	}

	return start
}
