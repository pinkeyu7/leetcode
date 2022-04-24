package main

func uniquePaths(m int, n int) int {
	// initial
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		a[0][i] = 1
	}
	for i := 0; i < m; i++ {
		a[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			a[i][j] = a[i-1][j] + a[i][j-1]
		}
	}

	return a[m-1][n-1]
}
