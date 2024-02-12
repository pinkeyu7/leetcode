package main

func majorityElement(nums []int) int {
	scoreMap := make(map[int]int)

	for _, num := range nums {
		scoreMap[num]++
	}

	maxNum := 0
	maxTimes := 0
	for num, times := range scoreMap {
		if times > maxTimes {
			maxNum = num
			maxTimes = times
		}
	}

	return maxNum
}
