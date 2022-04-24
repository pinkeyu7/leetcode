package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	nums1 := []int{2, 4, 3}
	nums2 := []int{5, 6, 4}

	l1 := generateListNode(nums1)
	l2 := generateListNode(nums2)

	output := addTwoNumbers(l1, l2)

	printListNode(output)
}
