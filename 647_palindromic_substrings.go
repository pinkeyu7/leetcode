package main

import "fmt"

func countSubstrings(s string) int {
	length := len(s)
	total := 0
	for i := 0; i < length; i++ {
		total += countSubstringsRecursive(s, i, i)
		total += countSubstringsRecursive(s, i, i+1)
	}

	return total
}

func countSubstringsRecursive(text string, left int, right int) int {
	total := 0
	for left >= 0 && right < len(text) {
		fmt.Println("left]:", left, "right:", right, "text[left]:", string(text[left]), "text[right]:", string(text[right]))
		if text[left] == text[right] {
			total++
		} else {
			break
		}
		left--
		right++
	}
	return total
}
