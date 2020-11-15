package leetcode

import "fmt"

// dfs
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				dfs(row, col, grid)
				count++
			}
		}
	}
	return count
}

func dfs(row, col int, grid [][]byte) {
	// termination
	if !inBound(row, col, grid) || grid[row][col] == '0' {
		return
	}
	// process
	grid[row][col] = '0'
	for i := 0; i < 4; i++ {
		newRow := row + dx[i]
		newCol := col + dy[i]
		// drill down
		dfs(newRow, newCol, grid)
	}
	// reverse
}

func inBound(row, col int, grid [][]byte) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

// bfs
var directions [][]int = [][]int{[]int{-1, 0}, []int{0, -1}, []int{1, 0}, []int{0, 1}}
var visited [][]bool

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	visited = make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	fmt.Println(visited)
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[i][j] {

				bfs(i, j, grid)
				count++
			}
		}
	}
	return count
}

func bfs(m, n int, grid [][]byte) {

	visited[m][n] = true
	cols := len(grid[0])
	queue := []int{m*cols + n}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		curX := front / cols
		curY := front % cols
		fmt.Println(curX, curY)
		for i := 0; i < 4; i++ {
			nRow := curX + directions[i][0]
			nCol := curY + directions[i][1]
			if nRow >= 0 && nRow < len(grid) && nCol >= 0 && nCol < len(grid[0]) && grid[nRow][nCol] == '1' && visited[nRow][nCol] == false {
				visited[nRow][nCol] = true
				queue = append(queue, nRow*cols+nCol)
			}
		}
	}
}
