package main

func minPathSum(grid [][]int) int {
	totalX := len(grid)
	if totalX == 0 {
		return 0
	}
	totalY := len(grid[0])

	for i := 1; i < totalY; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < totalX; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < totalX; i++ {
		for j := 1; j < totalY; j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] += grid[i-1][j]
			}
		}
	}

	return grid[totalX-1][totalY-1]
}
