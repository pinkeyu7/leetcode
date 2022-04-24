package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	numsLen := len(nums)
	res := make([][]int, 0)

	// initial case
	if numsLen == 1 {
		return append(res, nums)
	}

	// after recursive callback
	callback := permuteUnique(nums[1:])
	callbackLen := len(callback)

	mapCheck := make(map[string]bool)
	for i := 0; i < callbackLen; i++ {
		for j := 0; j < len(callback[i]); j++ {
			temp := make([]int, 0)
			temp = append(temp, callback[i][0:j]...)
			temp = append(temp, nums[0])
			temp = append(temp, callback[i][j:len(callback[i])]...)

			tempStr := fmt.Sprint(temp)
			if mapCheck[tempStr] {
				continue
			} else {
				mapCheck[tempStr] = true
				res = append(res, temp)
			}
		}

		temp := append(callback[i], nums[0])
		tempStr := fmt.Sprint(temp)
		if mapCheck[tempStr] {
			continue
		} else {
			mapCheck[tempStr] = true
			res = append(res, temp)
		}
	}

	return res
}
