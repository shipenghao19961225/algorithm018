package leetcode

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfs(i, 0, board)
		dfs(i, n-1, board)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, board)
		dfs(m-1, j, board)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
func dfs(i, j int, board [][]byte) {
	// termination
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if board[i][j] == '#' || board[i][j] == 'X' {
		return
	}
	// process
	board[i][j] = '#'
	// drill down
	dfs(i+1, j, board)
	dfs(i-1, j, board)
	dfs(i, j+1, board)
	dfs(i, j-1, board)
	// reverse the state if needed

}

// 可以并查集寫
// 1. 將所有邊緣的零點鏈接到哑结点
// 2. 将矩阵中的零链接到边缘结点
// 3. 遍历矩阵，看矩阵中的0是否链接到了哑结点，若没有，则置为'X'
//https://leetcode-cn.com/problems/surrounded-regions/solution/dfs-bing-cha-ji-java-by-liweiwei1419/

type UnionFind struct {
	count  int
	parent []int
	rank   []int
}

func (u *UnionFind) Init(count int) {
	u.count = count
	u.parent = make([]int, count)
	for i := 0; i < count; i++ {
		u.parent[i] = i
	}
	u.rank = make([]int, count)
}

// func(u *UnionFind) Find(node int) int {
//     for u.parent[node] != node {
//         u.parent[node] = u.parent[u.parent[node]]
//         node = u.parent[node]
//     }

//     return node
// }
func (u *UnionFind) Find(node int) int {
	if u.parent[node] == node {
		return node
	}
	u.parent[node] = u.Find(u.parent[node])
	return u.parent[node]
}

func (u *UnionFind) Union(node1, node2 int) {
	root1, root2 := u.Find(node1), u.Find(node2)
	if root1 == root2 {
		return
	}
	if u.rank[root1] > u.rank[root2] {
		u.parent[root2] = root1
	} else if u.rank[root1] < u.rank[root2] {
		u.parent[root1] = root2
	} else {
		u.parent[root1] = root2
		u.rank[root2]++
	}

	u.count--
}

func (u *UnionFind) Connected(node1, node2 int) bool {
	return u.Find(node1) == u.Find(node2)
}
