package main

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1

		for j := 1; j <= i; j++ {
			if j < i {
				row[j] = res[i-1][j-1] + res[i-1][j]
			} else {
				row[j] = 1
			}
		}
		res[i] = row
	}

	return res
}
