package main

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)

	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] != i {
			return i
		}
	}

	return length
}
