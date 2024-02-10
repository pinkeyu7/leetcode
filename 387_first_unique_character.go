package main

func firstUniqChar(s string) int {
	charMap := make(map[string]int)

	for _, c := range s {
		charMap[string(c)]++
	}

	for i, c := range s {
		if charMap[string(c)] == 1 {
			return i
		}
	}

	return -1
}
