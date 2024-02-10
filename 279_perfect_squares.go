package main

func numSquares(n int) int {
	scoreMap := make([]int, n+1)

	for i := 1; i <= n; i++ {
		minValue := n
		for j := 1; j*j <= i; j++ {
			if minValue > scoreMap[i-j*j] {
				minValue = scoreMap[i-j*j] + 1
			}
		}
		scoreMap[i] = minValue
	}

	return scoreMap[n]
}
