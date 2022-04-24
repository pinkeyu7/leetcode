package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var start *ListNode
	var current *ListNode

	addFlag := false
	for ok := true; ok; ok = (l1 != nil || l2 != nil) {
		tempSum := 0
		if addFlag {
			tempSum += 1
			addFlag = false
		}

		if l1 != nil {
			tempSum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tempSum += l2.Val
			l2 = l2.Next
		}

		if tempSum >= 10 {
			addFlag = true
			tempSum = tempSum % 10
		}

		temp := ListNode{
			Val:  tempSum,
			Next: nil,
		}

		if start == nil {
			start = &temp
			current = start
		} else {
			current.Next = &temp
			current = &temp
		}

	}

	if addFlag {
		temp := ListNode{
			Val:  1,
			Next: nil,
		}

		if start == nil {
			start = &temp
			current = start
		} else {
			current.Next = &temp
			current = &temp
		}

	}

	return start
}
