package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	// sort input
	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			start := j + 1
			end := length - 1

			for start < end {
				sum := nums[i] + nums[j] + nums[start] + nums[end]

				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					end--
					start++

					for start < end && nums[start] == nums[start-1] {
						start++
					}

					for start < end && nums[end] == nums[end+1] {
						end--
					}

				} else if sum > target {
					end--
				} else {
					start++
				}
			}
		}
	}

	return res
}
