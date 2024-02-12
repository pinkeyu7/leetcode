package main

func lengthOfLIS(nums []int) int {
	numLength := len(nums)
	maxLevel := 1

	scoreMap := make([]int, numLength)
	for i := 0; i < numLength; i++ {
		scoreMap[i] = 1
	}

	for i := 1; i < numLength; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if scoreMap[j] >= scoreMap[i] {
					scoreMap[i] = scoreMap[j] + 1
					if scoreMap[i] > maxLevel {
						maxLevel = scoreMap[i]
					}
				}
			}
		}
	}

	return maxLevel
}
