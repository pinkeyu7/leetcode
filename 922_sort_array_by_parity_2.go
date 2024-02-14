package main

func sortArrayByParityII(nums []int) []int {
	odd := 1
	even := 0

	for odd <= len(nums) && even <= len(nums) {
		if nums[odd]%2 == 1 {
			odd += 2
		} else if nums[even]%2 == 0 {
			even += 2
		} else {
			temp := nums[odd]
			nums[odd] = nums[even]
			nums[even] = temp
		}
	}

	return nums
}
