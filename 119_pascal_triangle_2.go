package main

func getRow(rowIndex int) []int {
	rowIndex++
	res := make([][]int, rowIndex)

	for i := 0; i < rowIndex; i++ {
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

	return res[rowIndex-1]
}
