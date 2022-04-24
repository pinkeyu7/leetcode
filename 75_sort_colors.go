package main

func sortColors(nums []int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return nums
	}

	var red int
	var white int
	var blue int

	for _, num := range nums {
		switch num {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}

	for i := range nums {
		if red > 0 {
			nums[i] = 0
			red--
			continue
		}
		if white > 0 {
			nums[i] = 1
			white--
			continue
		}
		if blue > 0 {
			nums[i] = 2
			blue--
			continue
		}
	}
	return nums
}
