package leetcode

func solveSudoku(board [][]byte) {
	var row, col, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				value := board[i][j] - '1'
				row[i][value] = 1
				col[j][value] = 1
				sbox[i/3*3+j/3][value] = 1
			}
		}
	}
	dfs(board, row, col, sbox)
}

func dfs(board [][]byte, row, col, sbox [9][9]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				for k := '1'; k <= '9'; k++ {
					v := k - '1'
					if row[i][v] == 0 && col[j][v] == 0 && sbox[i/3*3+j/3][v] == 0 {
						row[i][v], col[j][v], sbox[i/3*3+j/3][v] = 1, 1, 1
						board[i][j] = byte(k)
						if dfs(board, row, col, sbox) {
							return true
						}
						board[i][j] = '.'
						row[i][v], col[j][v], sbox[i/3*3+j/3][v] = 0, 0, 0
					}
				}
				return false
			}
		}
	}
	return true
}
