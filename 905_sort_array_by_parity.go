package main

func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]%2 == 1 && nums[right]%2 == 0 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
			right--
		} else if nums[left]%2 == 0 {
			left++
		} else if nums[right]%2 == 1 {
			right--
		}
	}

	return nums
}
