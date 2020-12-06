package leetcode

var res []string

// 这个题其实还是有剪枝的方法的，后面有时间再看！！！
type TrirNode struct {
	next [26]*TrirNode
	str  string
}

func findWords(board [][]byte, words []string) []string {
	res = make([]string, 0)
	// build trie
	root := new(TrirNode)
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'
			if node.next[index] == nil {
				node.next[index] = new(TrirNode)
			}
			node = node.next[index]
		}
		node.str = word
	}
	// start search
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, root)
		}
	}
	return res
}
func dfs(i, j int, board [][]byte, node *TrirNode) {
	// termination c1
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	// termination c2
	index := board[i][j] - 'a'
	if board[i][j] == '#' || node.next[index] == nil {
		return
	}
	// process
	// this step may be ignored
	node = node.next[index]
	if node.str != "" {
		res = append(res, node.str)
		node.str = ""
	}
	// drill down
	c := board[i][j]
	board[i][j] = '#'
	dfs(i+1, j, board, node)
	dfs(i, j+1, board, node)
	dfs(i-1, j, board, node)
	dfs(i, j-1, board, node)
	// reverse
	board[i][j] = c
}
