package main

func lengthOfLastWord(s string) int {
	right := len(s) - 1
	result := 0

	for right >= 0 {
		if string(s[right]) != " " {
			result++
		} else if result > 0 {
			break
		}

		right--
	}

	return result
}
