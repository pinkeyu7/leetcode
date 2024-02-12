package main

import (
	"math"
)

func mostFrequentEven(nums []int) int {
	scoreMap := make(map[int]int)
	maxTimes := 0
	for _, num := range nums {
		if num%2 == 0 {
			scoreMap[num]++
			if scoreMap[num] > maxTimes {
				maxTimes = scoreMap[num]
			}
		}
	}

	minNum := math.MaxInt64
	for num, times := range scoreMap {
		if times == maxTimes && minNum > num {
			minNum = num
		}
	}

	if minNum == math.MaxInt64 {
		return -1
	}

	return minNum
}
