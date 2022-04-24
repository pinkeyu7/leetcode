package main

func rotate(nums []int, k int) []int {
	numsLen := len(nums)
	k = k % numsLen

	if numsLen == 0 || k == 0 {
		return nums
	}
	temp := make([]int, numsLen)

	copy(temp, nums)
	for i, num := range temp {
		nums[(k+i)%numsLen] = num
	}

	return nums
}
