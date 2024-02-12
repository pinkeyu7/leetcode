package main

func topKFrequent(nums []int, k int) []int {
	scoreMap := make(map[int]int)
	resultMap := make(map[int][]int)
	result := make([]int, 0)

	maxTimes := 0
	for _, num := range nums {
		scoreMap[num]++
		if scoreMap[num] > maxTimes {
			maxTimes = scoreMap[num]
		}
	}

	for num, times := range scoreMap {
		resultMap[times] = append(resultMap[times], num)
	}

	for i := maxTimes; k > 0; i-- {
		k -= len(resultMap[i])
		result = append(result, resultMap[i]...)
	}

	return result
}
