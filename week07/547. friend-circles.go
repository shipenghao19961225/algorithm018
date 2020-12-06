package leetcode

// dfs解法
var visited map[int]bool

func findCircleNum(M [][]int) int {
	m := len(M)
	visited = make(map[int]bool)
	res := 0
	for i := 0; i < m; i++ {
		if !visited[i] {
			dfs(i, M)
			res++
		}
	}
	return res
}

func dfs(i int, M [][]int) {
	for j := 0; j < len(M); j++ {
		if !visited[j] && M[i][j] == 1 {
			visited[j] = true
			dfs(j, M)
		}
	}
}

// bfs解法
// bfs 注意添加visited数组
// 注意添加visited数组的位置，一般入队的时候就标记为已经访问
func findCircleNum(M [][]int) int {
	m := len(M)
	visited := make(map[int]bool)
	res := 0
	for i := 0; i < m; i++ {
		if !visited[i] {
			res++
			q := make([]int, 0)
			q = append(q, i)
			visited[i] = true
			for len(q) > 0 {
				size := len(q)
				for j := 0; j < size; j++ {
					topIndex := q[0]
					q = q[1:]
					for k := 0; k < m; k++ {
						if !visited[k] && M[topIndex][k] == 1 {
							q = append(q, k)
							visited[k] = true
						}
					}
				}
			}
		}
	}
	return res
}

// 并查集解法
func findCircleNum(M [][]int) int {
	UF := new(UnionFind)

	n := len(M)
	UF.Init(n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				UF.Union(i, j)
			}
		}
	}
	return UF.count
}

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

func (u *UnionFind) Find(node int) int {
	for u.parent[node] != node {
		u.parent[node] = u.parent[u.parent[node]]
		node = u.parent[node]
	}

	return node
}

// func(u *UnionFind) Find(node int) int {
//     if u.parent[node] == node {
//         return node
//     }
//     u.parent[node] = u.Find(u.parent[node])
//     return u.parent[node]
// }

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
