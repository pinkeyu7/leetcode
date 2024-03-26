package main

import (
	"sort"
)

func findDuplicate(nums []int) int {
	sort.Ints(nums)

	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return length
}
