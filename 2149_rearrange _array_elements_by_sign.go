package main

func rearrangeArray(nums []int) []int {
	positive := make([]int, 0)
	negative := make([]int, 0)
	result := make([]int, 0)

	for _, num := range nums {
		if num > 0 {
			positive = append(positive, num)
		} else {
			negative = append(negative, num)
		}
	}

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			result = append(result, positive[i/2])
		} else {
			result = append(result, negative[(i+1)/2-1])
		}
	}

	return result
}
