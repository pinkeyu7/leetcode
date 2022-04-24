package main

func mergeKLists(lists []*ListNode) *ListNode {
	var start *ListNode

	listLength := len(lists)
	if listLength == 0 {
		return start
	}

	for i := 0; i < listLength; i++ {
		start = mergeTwoLists(start, lists[i])
	}

	return start
}
