package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(l *ListNode) {
	current := l
	count := 1
	for current != nil {
		fmt.Println("index: ", count, "value: ", current.Val)
		current = current.Next
		count++
	}
}

func generateListNode(nums []int) *ListNode {
	var start *ListNode
	var current *ListNode

	length := len(nums)

	for i := 0; i < length; i++ {
		temp := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		if i == 0 {
			start = temp
		} else {
			current.Next = temp
		}
		current = temp
	}

	return start
}
