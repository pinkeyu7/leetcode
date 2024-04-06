package main

func numberOfSubArrays(nums []int, k int) int {
	totalCount := 0
	oddCount := 0  // keep track of odds to see if odds == k
	niceCount := 0 // keep track of current nice subarrays count (reset if a new odd number is found)

	start := 0
	for end := 0; end < len(nums); end++ { // outer-loop
		if nums[end]%2 != 0 {
			oddCount++
			niceCount = 0 // new odd number encountered, reset the nice subarray count
		}

		for oddCount == k { // inner-loop
			niceCount++

			if nums[start]%2 != 0 {
				oddCount--
			}
			start++
		}

		totalCount += niceCount
	}

	return totalCount
}
