package main

import (
	"sort"
)

func findDuplicates(nums []int) []int {
	sort.Ints(nums)

	mapTable := make(map[int]bool)
	result := make([]int, 0)
	length := len(nums)

	for i := 0; i < length-1; i++ {
		if nums[i] == nums[i+1] {
			mapTable[nums[i]] = true
		}
	}

	for key, _ := range mapTable {
		result = append(result, key)
	}

	sort.Ints(result)

	return result
}
