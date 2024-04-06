package main

func wordSearchExist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] && wordSearchExistRecursive(board, word, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func wordSearchExistRecursive(board [][]byte, word string, cursor int, i int, j int) bool {
	if cursor == len(word) {
		return true
	}

	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] != word[cursor] {
		return false
	}

	origin := board[i][j]
	board[i][j] = '_'

	if wordSearchExistRecursive(board, word, cursor+1, i+1, j) ||
		wordSearchExistRecursive(board, word, cursor+1, i-1, j) ||
		wordSearchExistRecursive(board, word, cursor+1, i, j+1) ||
		wordSearchExistRecursive(board, word, cursor+1, i, j-1) {
		return true
	}
	board[i][j] = origin

	return false
}
