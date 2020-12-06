package leetcode

func isValidSudoku(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				value := board[i][j] - '1'
				if row[i][value] == 1 {
					return false
				} else {
					row[i][value] = 1
				}
				if col[j][value] == 1 {
					return false
				} else {
					col[j][value] = 1
				}
				if sbox[i/3*3+j/3][value] == 1 {
					return false
				} else {
					sbox[i/3*3+j/3][value] = 1
				}
			}
		}
	}
	return true
}
