package leetcode

type point struct {
	x int
	y int
}

var dirX = []int{0, 1, 0, -1}
var dirY = []int{1, 0, -1, 0}

func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[point]bool)
	for _, obstacle := range obstacles {
		m[point{obstacle[0], obstacle[1]}] = true
	}
	x, y := 0, 0
	d := 0
	res := 0
	for _, command := range commands {
		if command == -1 {
			d = (d + 1 + 4) % 4
		} else if command == -2 {
			d = (d - 1 + 4) % 4
		} else {
			for i := 0; i < command; i++ {
				newX, newY := x+dirX[d], y+dirY[d]
				if _, ok := m[point{newX, newY}]; ok {
					break
				} else {
					x, y = newX, newY
					res = max(res, x*x+y*y)
				}
			}
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
