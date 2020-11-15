package leetcode

// dfs
var directionsX = []int{-1, 0, 1, 1, 1, 0, -1, -1}
var directionsY = []int{1, 1, 1, 0, -1, -1, -1, 0}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(x, y, board)
	}
	return board
}
func dfs(x, y int, board [][]byte) {
	// termination
	if !inbound(x, y, board) || board[x][y] != 'E' {
		return
	}

	// process current level
	count := 0
	for i := 0; i < 8; i++ {
		newX := x + directionsX[i]
		newY := y + directionsY[i]
		if inbound(newX, newY, board) && board[newX][newY] == 'M' {
			count++
		}
	}
	if count != 0 {
		board[x][y] = byte('0' + count)
		return
	}
	// drill down
	board[x][y] = 'B'
	for i := 0; i < 8; i++ {

		dfs(x+directionsX[i], y+directionsY[i], board)
	}
	// reverse
}
func inbound(x, y int, board [][]byte) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

// bfs
var directionsX = []int{-1, 0, 1, 1, 1, 0, -1, -1}
var directionsY = []int{1, 1, 1, 0, -1, -1, -1, 0}
var visited [][]bool

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	visited = make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		bfs(x, y, board)
	}
	return board
}
func bfs(x, y int, board [][]byte) {
	queue := [][]int{[]int{x, y}}
	visited[x][y] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			count := 0
			for i := 0; i < 8; i++ {
				newX := top[0] + directionsX[i]
				newY := top[1] + directionsY[i]
				if inbound(newX, newY, board) && board[newX][newY] == 'M' {
					count++
				}
			}
			if count > 0 {
				board[top[0]][top[1]] = byte(count + '0')
				continue
			}
			for i := 0; i < 8; i++ {
				newX := top[0] + directionsX[i]
				newY := top[1] + directionsY[i]
				if inbound(newX, newY, board) && !visited[newX][newY] && board[newX][newY] == 'E' {
					visited[newX][newY] = true
					queue = append(queue, []int{newX, newY})

				}
			}
			board[top[0]][top[1]] = 'B'

		}
	}
}

func inbound(x, y int, board [][]byte) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
