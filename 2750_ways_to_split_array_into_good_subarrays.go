package main

func numberOfGoodSubarraySplits(nums []int) int {
	totalCount := 1
	lastOneIndex := -1

	for i, num := range nums {
		if num == 1 {
			if lastOneIndex != -1 {
				totalCount *= i - lastOneIndex
				totalCount %= 1000000007 // Why should mod7
			}
			lastOneIndex = i
		}
	}

	if lastOneIndex == -1 {
		return 0
	}

	return totalCount
}
