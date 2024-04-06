package main

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)

	for _, num := range nums {
		_, ok := numMap[num]
		if ok {
			return true
		} else {
			numMap[num] = true
		}
	}

	return false
}
