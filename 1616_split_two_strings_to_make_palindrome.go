package main

func checkPalindromeFormation(a string, b string) bool {
	return checkPalindromeFormationSub(a, b) || checkPalindromeFormationSub(b, a)
}

func checkPalindromeFormationSub(a string, b string) bool {
	n := len(a)
	start := 0
	end := n - 1

	for start < end && a[start] == b[end] {
		start++
		end--
	}

	return validPalindrome(a[start:end+1]) || validPalindrome(b[start:end+1])
}
