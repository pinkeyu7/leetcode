package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// sort input
	sort.Ints(nums)
	length := len(nums)

	smallest := nums[0] + nums[1] + nums[2]
	biggest := nums[length-1] + nums[length-2] + nums[length-3]

	for i := 0; i < length; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		start := i + 1
		end := length - 1

		for start < end {
			sum := nums[i] + nums[start] + nums[end]

			if sum == target {
				return target

			} else if sum > target {
				end--
				if sum < biggest {
					biggest = sum
				}
			} else {
				start++
				if sum > smallest {
					smallest = sum
				}
			}
		}
	}

	if biggest-target > target-smallest {
		return smallest
	} else {
		return biggest
	}
}
