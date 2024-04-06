package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		bMap := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			word := board[i][j]
			if word == 46 {
				continue
			}
			_, ok := bMap[word]
			if ok {
				return false
			} else {
				bMap[word] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		bMap := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			word := board[j][i]
			if word == 46 {
				continue
			}
			_, ok := bMap[word]
			if ok {
				return false
			} else {
				bMap[word] = true
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			bMap := make(map[byte]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					word := board[i*3+k][j*3+l]
					if word == 46 {
						continue
					}
					_, ok := bMap[word]
					if ok {
						return false
					} else {
						bMap[word] = true
					}
				}
			}
		}
	}

	return true
}
