package main

func removeElement(nums []int, val int) int {
	length := len(nums)
	first := 0
	last := length - 1

	for first <= last {
		if nums[first] == val {
			nums[first] = nums[last]
			last--
		} else {
			first++
		}
	}
	return first
}
