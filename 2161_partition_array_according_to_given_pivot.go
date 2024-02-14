package main

func pivotArray(nums []int, pivot int) []int {
	more := make([]int, 0)
	result := make([]int, 0)
	pivotCount := 0

	for _, num := range nums {
		if num > pivot {
			more = append(more, num)
		} else if num < pivot {
			result = append(result, num)
		} else {
			pivotCount++
		}
	}

	for i := 0; i < pivotCount; i++ {
		result = append(result, pivot)
	}
	result = append(result, more...)

	return result
}
