package main

func numSubarrayProductLessThanK(nums []int, target int) int {
	if target <= 1 {
		return 0
	}

	left := 0
	product := 1
	res := 0
	for right, num := range nums {
		product *= num
		for product >= target {
			product /= nums[left]
			left++
		}

		res += right - left + 1
	}

	return res
}

func numSubarrayProductLessThanKTLE(nums []int, target int) int {
	if target <= 1 {
		return 0
	}

	length := len(nums)
	count := 0

	boolList := make([]bool, length)

	for i := 1; i <= length; i++ {
		for j := 0; j <= length-i; j++ {
			if boolList[j] {
				continue
			}

			if getProduct(nums[j:j+i], target) {
				count++
			} else {
				boolList[j] = true
			}
		}
	}
	return count
}

func getProduct(nums []int, target int) bool {
	length := len(nums)
	product := 1
	for i := 0; i < length; i++ {
		product *= nums[i]
		if product >= target {
			return false
		}
	}
	return true
}
