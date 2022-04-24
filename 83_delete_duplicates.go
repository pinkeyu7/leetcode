package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	flag := head
	start := head
	current := head

	for current != nil {
		if flag.Val != current.Val {
			flag.Next = current
			flag = current
		}
		current = current.Next
	}

	if flag.Next != nil {
		if flag.Val == flag.Next.Val {
			flag.Next = nil
		}
	}

	return start
}
