package main

func maxArea(height []int) int {
	max := 0
	cursor := 0
	left := 0
	right := len(height) - 1

	for left < right {
		tempHeight := 0
		if height[left] > height[right] {
			tempHeight = height[right]
			cursor = left
		} else {
			tempHeight = height[left]
			cursor = right
		}

		area := tempHeight * (right - left)
		if area > max {
			max = area
		}

		if left == cursor {
			right--
		} else {
			left++
		}
	}

	return max
}
