package main

import "sort"

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	mapTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if ind, ok := mapTable[temp]; ok {
			res = append(res, ind)
			res = append(res, i)
			break
		}
		mapTable[nums[i]] = i
	}

	sort.Ints(res)

	return res
}
