package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	cursor := 0
	input1Cursor := 0
	input2Cursor := 0
	firstIndex := totalLength / 2
	secondIndex := totalLength / 2
	firstNum := 0
	secondNum := 0

	if totalLength%2 == 0 {
		firstIndex--
	}

	for cursor <= secondIndex {
		currentNum := 0
		if input1Cursor < len(nums1) && input2Cursor < len(nums2) {
			if nums1[input1Cursor] > nums2[input2Cursor] {
				currentNum = nums2[input2Cursor]
				input2Cursor++
			} else {
				currentNum = nums1[input1Cursor]
				input1Cursor++
			}
		} else if input1Cursor < len(nums1) {
			currentNum = nums1[input1Cursor]
			input1Cursor++
		} else {
			currentNum = nums2[input2Cursor]
			input2Cursor++
		}

		if cursor == firstIndex {
			firstNum = currentNum
		}
		if cursor == secondIndex {
			secondNum = currentNum
		}
		cursor++
	}

	return float64(firstNum+secondNum) / 2
}
