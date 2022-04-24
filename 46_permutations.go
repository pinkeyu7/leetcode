package main

func permute(nums []int) [][]int {
	numsLen := len(nums)
	res := make([][]int, 0)

	// initial case
	if numsLen == 1 {
		return append(res, nums)
	}

	// after recursive callback
	callback := permute(nums[1:])
	callbackLen := len(callback)
	for i := 0; i < callbackLen; i++ {
		for j := 0; j < len(callback[i]); j++ {
			temp := make([]int, 0)
			temp = append(temp, callback[i][0:j]...)
			temp = append(temp, nums[0])
			temp = append(temp, callback[i][j:len(callback[i])]...)
			res = append(res, temp)
		}
		res = append(res, append(callback[i], nums[0]))
	}
	return res
}
