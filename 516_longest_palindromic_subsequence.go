package main

func longestPalindromeSubseq(s string) int {
	sLength := len(s)
	rev := reverseString(s)

	scoreMap := make([][]int, sLength+1)
	for i := 0; i <= sLength; i++ {
		scoreMap[i] = make([]int, sLength+1)
	}

	for i := 1; i <= sLength; i++ {
		for j := 1; j <= sLength; j++ {
			if s[i-1] == rev[j-1] {
				scoreMap[i][j] = scoreMap[i-1][j-1] + 1
			} else {
				if scoreMap[i-1][j] > scoreMap[i][j-1] {
					scoreMap[i][j] = scoreMap[i-1][j]
				} else {
					scoreMap[i][j] = scoreMap[i][j-1]
				}
			}
		}
	}

	return scoreMap[sLength][sLength]
}
