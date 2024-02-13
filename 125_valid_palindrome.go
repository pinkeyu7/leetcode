package main

func isPalindromeString(s string) bool {
	word := ""
	for _, char := range s {
		if char > 47 && char < 58 {
			word += string(char)
		}
		if char > 96 && char < 123 {
			word += string(char)
		}
		if char > 64 && char < 91 {
			word += string(char + 32)
		}
	}

	if len(word) <= 1 {
		return true
	}

	return validPalindrome(word)
}
