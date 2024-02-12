package main

import "sort"

func majorityElement2(nums []int) []int {
	scoreMap := make(map[int]int)
	resultMap := make(map[int]bool)
	result := make([]int, 0)

	for _, num := range nums {
		scoreMap[num]++
		if scoreMap[num]*3 > len(nums) && !resultMap[num] {
			resultMap[num] = true
			result = append(result, num)
		}
	}

	sort.Ints(result)
	return result
}
