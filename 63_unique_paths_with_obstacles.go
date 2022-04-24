package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	totalX := len(obstacleGrid)
	if totalX == 0 {
		return 0
	}
	totalY := len(obstacleGrid[0])

	// initial
	a := make([][]int, totalX)
	for i := range a {
		a[i] = make([]int, totalY)
	}

	for i := 0; i < totalX; i++ {
		for j := 0; j < totalY; j++ {
			if obstacleGrid[i][j] == 1 {
				a[i][j] = -1
			}
		}
	}

	for i := 0; i < totalY; i++ {
		if a[0][i] == -1 {
			break
		}
		a[0][i] = 1
	}
	for i := 0; i < totalX; i++ {
		if a[i][0] == -1 {
			break
		}
		a[i][0] = 1
	}

	for i := 1; i < totalX; i++ {
		for j := 1; j < totalY; j++ {
			if a[i][j] == -1 {
				continue
			}
			if a[i-1][j] != -1 {
				a[i][j] += a[i-1][j]
			}
			if a[i][j-1] != -1 {
				a[i][j] += a[i][j-1]
			}
		}
	}

	if a[totalX-1][totalY-1] == -1 {
		return 0
	}

	return a[totalX-1][totalY-1]
}
