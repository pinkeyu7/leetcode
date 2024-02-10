package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	length := len(nums)
	maxIndex := 0
	maxScore := 0
	dfsMap := make([]int, length)

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if dfsMap[j] >= dfsMap[i] {
					dfsMap[i] = dfsMap[j] + 1
				}
				if dfsMap[i] > maxScore {
					maxIndex = i
					maxScore = dfsMap[i]
				}
			}
		}
	}

	maxValue := nums[maxIndex]
	result := make([]int, 0)
	for i := maxIndex; i >= 0; i-- {
		if maxValue%nums[i] == 0 && dfsMap[i] == maxScore {
			result = append(result, nums[i])
			maxValue = nums[i]
			maxScore--
		}
	}

	sort.Sort(sort.IntSlice(result))
	return result
}
