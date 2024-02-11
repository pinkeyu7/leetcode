package main

func longestPalindrome(s string) string {
	sLength := len(s)

	if sLength <= 1 {
		return s
	}

	maxString := s[0:1]

	for i := 0; i < sLength-1; i++ {
		odd := expandFromCenter(s, i, i)
		even := expandFromCenter(s, i, i+1)

		if len(odd) > len(maxString) {
			maxString = odd
		}
		if len(even) > len(maxString) {
			maxString = even
		}
	}
	return maxString
}

func expandFromCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
