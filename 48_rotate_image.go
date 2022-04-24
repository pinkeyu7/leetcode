package main

func rotateImage(matrix [][]int) [][]int {
	borderSize := len(matrix)
	lastIndex := borderSize - 1

	res := make([][]int, borderSize)
	for i := 0; i < borderSize; i++ {
		res[i] = make([]int, borderSize)
		for j := 0; j < borderSize; j++ {
			res[i][j] = matrix[i][j]
		}
	}

	for i := 0; i < borderSize; i++ {
		for j := 0; j < borderSize; j++ {
			matrix[j][lastIndex-i] = res[i][j]
		}
	}

	return matrix
}
