package main

func minDistance(word1 string, word2 string) int {
	text1Length := len(word1)
	text2Length := len(word2)

	scoreMap := make([][]int, text1Length+1)
	for i := 0; i <= text1Length; i++ {
		scoreMap[i] = make([]int, text2Length+1)
	}

	for i := 1; i <= text1Length; i++ {
		for j := 1; j <= text2Length; j++ {
			if word1[i-1] == word2[j-1] {
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

	return text1Length + text2Length - scoreMap[text1Length][text2Length]*2
}
