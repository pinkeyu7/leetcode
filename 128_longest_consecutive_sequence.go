package main

func longestConsecutive(nums []int) int {
	numMap := make(map[int]struct{})
	longest := 0

	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	for num, _ := range numMap {
		if _, ok := numMap[num-1]; !ok {
			length := 1
			for _, ok := numMap[num+length]; ok; _, ok = numMap[num+length] {
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
